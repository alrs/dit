SHA := $(shell git rev-parse --short HEAD)
TARNAME := ditcli-$(SHA)
DATE := $(shell date -u)
LDFLAGS := -ldflags="-X 'main.sha=$(SHA)' -X 'main.buildDate=$(DATE)'"

help:
	@echo make all: build all binaries
	@echo make clean: delete all built binaries
	@echo make tar: build release tarball
	@echo make install: build and install dit on local GOPATH
 
.PHONY: clean
clean:
	rm -rf ditcli-*
	
.PHONY: all
all: \
	$(TARNAME)/linux/amd64/dit	\
	$(TARNAME)/linux/arm/dit	\
	$(TARNAME)/openbsd/amd64/dit	\
	$(TARNAME)/darwin/amd64/dit	\
	$(TARNAME)/darwin/arm64/dit	\
	$(TARNAME)/windows/amd64/dit
 
ditcli-$(SHA)/%/dit:
	mkdir -p $*
	GOOS=$(shell echo $* | awk -F/ '{print $$1}')	\
	GOARCH=$(shell echo $* | awk -F/ '{print $$2}')	\
	go build $(LDFLAGS) -o $@

.PHONY: install
install:
	go install $(LDFLAGS) ./...

.PHONY: tar
tar: all $(TARNAME)/README $(TARNAME)/LICENSE
	tar czv $(TARNAME) > $(TARNAME).tar.gz

$(TARNAME)/README:
	cp README $(TARNAME)

$(TARNAME)/LICENSE:
	cp ../../LICENSE $(TARNAME)
