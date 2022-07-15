BIN_DIR = $(PWD)/bin

OSFLAG 				:=
ifeq ($(OS),Windows_NT)
	OSFLAG += -D WIN32
	ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
		OSFLAG += -D AMD64
	endif
	ifeq ($(PROCESSOR_ARCHITECTURE),x86)
		OSFLAG += -D IA32
	endif
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		OSFLAG += -D LINUX
	endif
	ifeq ($(UNAME_S),Darwin)
		OSFLAG += -D OSX
	endif
		UNAME_P := $(shell uname -p)
	ifeq ($(UNAME_P),x86_64)
		OSFLAG += -D AMD64
	endif
		ifneq ($(filter %86,$(UNAME_P)),)
	OSFLAG += -D IA32
		endif
	ifneq ($(filter arm%,$(UNAME_P)),)
		OSFLAG += -D ARM
	endif
endif

echoos:
	@echo $(OSFLAG)

# Will setup linting tools
setup-linting:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin v1.41.0
	chmod +x ./bin/golangci-lint

# Download dependencies
install:
	go mod download

# remove unused dependencies
tidy:
	go mod tidy -v

# Runs project
run:
	go run app/cmd/main.go

test:
	go test ./...

test-coverage:
	go test -tags testing -v -cover -covermode=atomic ./...

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

clean:
	if [ -f ${BIN_DIR} ] ; then rm ${BIN_DIR} ; fi

lint:
	./bin/golangci-lint run ./...

build:
	@echo "Building application"
	go build -o $(BIN_DIR) cmd/main.go

all: install lint test