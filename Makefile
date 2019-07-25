all:	update-python-abi update-python-protobuf update-go-protobuf

# pants cannot handle symlinks that point to outside its root

update-python-abi:	clients/python/resources/compiled-contracts/NeutralValidator.json

clients/python/resources/compiled-contracts/NeutralValidator.json: compiled-contracts/NeutralValidator.json
	@cp -f compiled-contracts/NeutralValidator.json $@
	git add $@


DIFF_PYTHON_PROTOBUF := $(shell diff -r ./protobuf clients/python/src/protobuf -x BUILD > /dev/null 2>&1; echo $$?)

clients/python/src/protobuf:
	@mkdir -p $@

update-python-protobuf: clients/python/src/protobuf
ifeq ($(DIFF_PYTHON_PROTOBUF), 1)
	@echo "updating clients/python/src/protobuf with ./protobuf ......"
	@rm -f clients/python/src/protobuf/*.proto
	@cp -f ./protobuf/*.proto  clients/python/src/protobuf/
	git add clients/python/src/protobuf/
else
	@echo "./protobuf and clients/python/src/protobuf are the same."
endif

update-go-protobuf:
	cd clients/go && make grpc-go

clean:
	cd clients/java && ./gradlew clean
	cd clients/python && ./pants clean-all
	cd clients/go && make clean
