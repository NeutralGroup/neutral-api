BINDIR=$(CURDIR)/bin
SRCDIR=$(CURDIR)/src
GOPATH=$(CURDIR)

PROTO_DIR=$(CURDIR)/../../protobuf
PROTO_SOURCES := $(wildcard $(PROTO_DIR)/*.proto)
PROTO_SOURCE_FILENAMES := $(foreach path, $(PROTO_SOURCES), $(shell basename $(path)))
PROTO_SOURCE_BASENAMES := $(basename $(PROTO_SOURCE_FILENAMES))

GO_OUT_DIR=$(SRCDIR)/neutralproject.com/lib/protocol
GO_PROTO_OUTPUT_FILES := $(foreach base, $(PROTO_SOURCE_BASENAMES), $(GO_OUT_DIR)/$(base).pb.go)

all:	$(BINDIR)/demo

.EXPORT_ALL_VARIABLES:

.ONESHELL:

$(BINDIR):
	mkdir -p $(BINDIR)

$(GO_PROTO_OUTPUT_FILES): grpc-go
	
grpc-go: $(PROTO_SOURCES)
	@echo $(PROTO_SOURCE_FILENAMES)
	@echo $(GO_PROTO_OUTPUT_FILES)
	@mkdir -p $(GO_OUT_DIR)
	@rm -r $(GO_OUT_DIR)/* || true
	@echo compiling $(PROTO_SOURCES) to $(GO_OUT_DIR)
	export PATH=$(PATH):$(GOPATH)/bin;
	go get -u github.com/golang/protobuf/protoc-gen-go
	protoc \
		-I$(PROTO_DIR) \
		--go_out=plugins=grpc:$(GO_OUT_DIR) \
		$(PROTO_SOURCES)
	git add $(GO_OUT_DIR)

$(BINDIR)/demo: $(BINDIR)
	@echo "building $@"
	@echo "GOPATH = $(GOPATH)"
	cd $(SRCDIR)/neutralproject.com/clients/demo
	@echo "go get ......"
	go get
	go build -o $@

clean:
	rm -rf $(BINDIR)
