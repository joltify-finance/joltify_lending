#!/usr/bin/make -f

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
TM_PKG_VERSION := $(shell go list -m  github.com/cometbft/cometbft  | sed 's:.* ::')
COSMOS_PKG_VERSION := $(shell go list -m github.com/cosmos/cosmos-sdk | sed 's:.* ::')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= true
PROJECT_NAME = joltify
DOCKER:=docker
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf
HTTPS_GIT := https://github.com/joltify-finance/joltify_lending.git

export GO111MODULE = on

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  build_tags += gcc
endif

ifeq (secp,$(findstring secp,$(COSMOS_BUILD_OPTIONS)))
  build_tags += libsecp256k1_sdk
endif

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=joltify\
		  -X github.com/cosmos/cosmos-sdk/version.AppName=joltify \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
		  -X github.com/cometbft/cometbft/version.TMCoreSemVer=$(TM_PKG_VERSION)


# DB backend selection
ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq (badgerdb,$(findstring badgerdb,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb
  BUILD_TAGS += badgerdb
endif
# handle rocksdb
ifeq (rocksdb,$(findstring rocksdb,$(COSMOS_BUILD_OPTIONS)))
  CGO_ENABLED=1
  BUILD_TAGS += rocksdb
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb
endif
# handle boltdb
ifeq (boltdb,$(findstring boltdb,$(COSMOS_BUILD_OPTIONS)))
  BUILD_TAGS += boltdb
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=boltdb
endif


ldflagsdev = $(ldflags) -X github.com/joltify-finance/joltify_lending/client/client.MAINNETFLAG=false
ldflagsmainnet = $(ldflags) -X github.com/joltify-finance/joltify_lending/client/client.MAINNETFLAG=true



ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  ldflagsdev += -w -s
  ldflagsmainnet += -w -s
endif
#ldflags += $(LDFLAGS)
ldflagsmainnet := $(strip $(ldflagsmainnet))
ldflagsdev := $(strip $(ldflagsdev))



build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

BUILD_FLAGS_MAINNET := -tags "$(build_tags)" -ldflags '$(ldflagsmainnet)'
BUILD_FLAGS_DEV := -tags "$(build_tags)" -ldflags '$(ldflagsdev)'

# check for nostrip option
ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
  BUILD_FLAGS_DEV += -trimpath
endif

all: install

build: go.sum
ifeq ($(OS), Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS_MAINNET) -o build/$(shell go env GOOS)/joltify.exe ./cmd/joltify
else
	go build -mod=readonly $(BUILD_FLAGS_MAINNET) -o build/$(shell go env GOOS)/joltify ./cmd/joltify
endif

dev: go.sum
ifeq ($(OS), Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS_DEV) -o build/$(shell go env GOOS)/joltify.exe ./cmd/joltify
else
	go build -mod=readonly $(BUILD_FLAGS_DEV) -o build/$(shell go env GOOS)/joltify ./cmd/joltify
endif




build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS_MAINNET) ./cmd/joltify

#dev: go.sum
#	go install -mod=readonly $(BUILD_FLAGS_DEV) ./cmd/joltify

########################################
### Tools & dependencies

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download
PHONY: go-mod-cache

#go.sum: go.mod
go.sum:
	@echo "--> Ensuring dependencies have not been modified"
	@go mod verify

clean:
	rm -rf build/


lint:
	@golangci-lint run --out-format=tab  -v --timeout 3600s -c ./.golangci.yml
	go mod verify
.PHONY: lint

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs goimports -w -local github.com/tendermint
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs goimports -w -local github.com/cosmos/cosmos-sdk
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -name '*.pb.go' | xargs goimports -w -local github.com/joltify-finance/joltify_lending
.PHONY: format

###############################################################################
###                                Localnet                                 ###
###############################################################################
# Launch a new single validator chain
start:
	rm -rf ~/.joltify
	./contrib/devnet/init-new-chain.sh
	joltify start
test:
	gotestsum  --junitfile report.xml --format testname  -- -coverprofile=coverage.out -timeout 15m ./...
	cat coverage.out |grep -v "erc20.go"|grep -v "oppy_transfer.go" > cover.out
	go tool cover -func=cover.out


###############################################################################
###                                Protobuf                                 ###
###############################################################################


#
#protoVer=v0.13.0
#protoImageName=joltify/joltify-proto-gen:$(protoVer)
#
#
#containerProtoGen=cosmos-sdk-proto-gen-$(protoVer)
#containerProtoGenSwagger=cosmos-sdk-proto-swagger-$(protoVer)
#containerProtoFmt=cosmos-sdk-proto-fmt-$(protoVer)
#protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)


protoVer=0.13.0
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)



proto-all: proto-gen proto-format proto-lint proto-swagger-gen

proto-gen:
	@echo "Generating Protobuf files"
	@$(protoImage) sh ./scripts/protocgen.sh
#	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGen}$$"; then docker start -a $(containerProtoGen); else docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) \
		sh ./scripts/protocgen.sh; fi

proto-swagger-gen:
	@echo "Generating Protobuf Swagger"
	@$(protoImage) sh ./scripts/protoc-swagger-gen.sh
	#@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoGenSwagger}$$"; then docker start -a $(containerProtoGenSwagger); else docker run --name $(containerProtoGenSwagger) -v $(CURDIR):/workspace --workdir /workspace $(protoImageName) \
	#	sh ./scripts/protoc-swagger-gen.sh; fi
	@echo "now statik all the files"
	@statik -src=client/docs/swagger-ui -dest=client/docs -f -m
.PHONY: proto-swagger-gen



proto-format:
	@echo "Formatting Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoFmt}$$"; then docker start -a $(containerProtoFmt); else docker run --name $(containerProtoFmt) -v $(CURDIR):/workspace --workdir /workspace tendermintdev/docker-build-proto \
		find ./ -not -path "./third_party/*" -name *.proto -exec clang-format -style=file -i {} \; ; fi
proto-image-build:
	@DOCKER_BUILDKIT=1 docker build -t $(protoImageName) -f ./proto/Dockerfile ./proto

proto-lint:
	@$(DOCKER_BUF) lint --error-format=json

proto-check-breaking:
	@$(DOCKER_BUF) breaking --against $(HTTPS_GIT)#branch=master

GOOGLE_PROTO_URL = https://raw.githubusercontent.com/googleapis/googleapis/master/google/api
PROTOBUF_GOOGLE_URL = https://raw.githubusercontent.com/protocolbuffers/protobuf/master/src/google/protobuf
COSMOS_PROTO_URL = https://raw.githubusercontent.com/cosmos/cosmos-proto/master

GOOGLE_PROTO_TYPES = third_party/proto/google/api
PROTOBUF_GOOGLE_TYPES = third_party/proto/google/protobuf
COSMOS_PROTO_TYPES = third_party/proto/cosmos_proto

#GOGO_PATH := $(shell go list -m -f '{{.Dir}}' github.com/gogo/protobuf)
#TENDERMINT_PATH := $(shell go list -m -f '{{.Dir}}' github.com/cometbft/cometbft/)
#COSMOS_PROTO_PATH := $(shell go list -m -f '{{.Dir}}' github.com/cosmos/cosmos-proto)
#COSMOS_SDK_PATH := $(shell go list -m -f '{{.Dir}}' github.com/cosmos/cosmos-sdk)
#IBC_GO_PATH := $(shell go list -m -f '{{.Dir}}' github.com/cosmos/ibc-go/v6)
#
#proto-update-deps:
#	mkdir -p $(GOOGLE_PROTO_TYPES)
#	curl -sSL $(GOOGLE_PROTO_URL)/annotations.proto > $(GOOGLE_PROTO_TYPES)/annotations.proto
#	curl -sSL $(GOOGLE_PROTO_URL)/http.proto > $(GOOGLE_PROTO_TYPES)/http.proto
#	curl -sSL $(GOOGLE_PROTO_URL)/httpbody.proto > $(GOOGLE_PROTO_TYPES)/httpbody.proto
#
#	mkdir -p $(PROTOBUF_GOOGLE_TYPES)
#	curl -sSL $(PROTOBUF_GOOGLE_URL)/any.proto > $(PROTOBUF_GOOGLE_TYPES)/any.proto
#
#	mkdir -p client/docs
#	cp $(COSMOS_SDK_PATH)/client/docs/swagger-ui/swagger.yaml client/docs/cosmos-swagger.yml
#	cp $(IBC_GO_PATH)/docs/client/swagger-ui/swagger.yaml client/docs/ibc-go-swagger.yml
#
#	mkdir -p $(COSMOS_PROTO_TYPES)
#	cp $(COSMOS_PROTO_PATH)/cosmos.proto $(COSMOS_PROTO_TYPES)/cosmos.proto
#
#	rsync -r --chmod 644 --include "*.proto" --include='*/' --exclude='*' $(GOGO_PATH)/gogoproto third_party/proto
#	rsync -r --chmod 644 --include "*.proto" --include='*/' --exclude='*' $(TENDERMINT_PATH)/proto third_party
#	rsync -r --chmod 644 --include "*.proto" --include='*/' --exclude='*' $(COSMOS_SDK_PATH)/proto third_party
#	rsync -r --chmod 644 --include "*.proto" --include='*/' --exclude='*' $(IBC_GO_PATH)/proto third_party
#	rsync -r --chmod 644 --include "*.proto" --include='*/' --exclude='*' $(ETHERMINT_PATH)/proto third_party
#	cp -f $(IBC_GO_PATH)/third_party/proto/proofs.proto third_party/proto/proofs.proto

integration-test:
	make -C contrib/devnet/integrationtest build


.PHONY: proto-all proto-gen proto-gen-any proto-swagger-gen proto-format proto-lint proto-check-breaking proto-update-deps
