include Makefile.defs

.PHONY: help
help:
	@echo 'Management commands for CompanyService:'
	@echo
	@echo 'Usage:'
	@echo '    make build           	Compile the project.'
	@echo '    make get-deps			Runs dep ensure'
	@echo '    make gen-openapi 		Generate openapi code'
	@echo '    make gen-openapi-server  Generate openapi code for server'
	@echo '    make run-openapi  		Run openapi server'
	@echo '    make run  				Run service'
	@echo '    make docker_build  		Docker build'
	@echo '    make docker_run  		Docker run'
	@echo '    make test            	Run tests on a compiled project.'
	@echo '    make clean           	Clean the directory tree'
	@echo '    make full_clean      	Clean the directory tree (and vendor directory)'
	@echo

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo $(GOBUILD) ./cmd/$(PROJECT_NAME)/main.go

.PHONY: get-deps
get-deps:
	dep ensure

.PHONY: gen-openapi
gen-openapi:
	@make gen-openapi-server
	@make gen-openapi-productservice

.PHONY: gen-openapi-server
gen-openapi-server:
	mkdir -p $(OPENAPI_GEN_DIR)
	swagger generate server -t $(OPENAPI_GEN_DIR) -f $(OPENAPI_SPEC_FILE) -s server --exclude-main -A $(PROJECT_CAPITAL_NAME)

.PHONY: gen-openapi-productservice
gen-openapi-productservice:
	mkdir -p $(OPENAPI_PRODUCTSERVICE_GEN_DIR)
	swagger generate client -t $(OPENAPI_PRODUCTSERVICE_GEN_DIR) -f $(OPENAPI_PRODUCTSERVICE_SPEC_FILE)

.PHONY: run-openapi
run-openapi:
	swagger serve ./openapi/spec/$(PROJECT_NAME)-openapi.yaml

.PHONY: run
run:
	go run cmd/$(PROJECT_NAME)/main.go

.PHONY: docker_build
docker_build:
	rm -Rf $(OPENAPI_GEN_DIR)
	docker build -t $(PROJECT_NAME) .

.PHONY: docker_v1
docker_v1:
	docker build -t zhangzj/company_service:v1 .
	docker push zhangzj/company_service:v1 

.PHONY: docker_v2
docker_v2:
	docker build -t zhangzj/company_service:v2 .
	docker push zhangzj/company_service:v2 

.PHONY: docker_run
docker_run:
	docker run -e SERVICEPORT=3000  -p 3000:3000 $(PROJECT_NAME)

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm ./main
	rm -Rf ./openapi/gen

.PHONY: full_clean
full_clean: clean
	rm -Rf ./vendor