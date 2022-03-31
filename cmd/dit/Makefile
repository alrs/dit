SHA := $(shell git rev-parse --short HEAD)
TARNAME := ditcli-$(SHA)
DATE := $(shell date -u)

help:
	@echo make all: build all binaries
	@echo make clean: delete all built binaries
	@echo make tar: build release tarball

.PHONY: clean
clean:
	rm -rf ditcli-*
	
.PHONY: all
all: \
	$(TARNAME)/linux/amd64/dit	\
	$(TARNAME)/linux/arm/dit 	\
	$(TARNAME)/openbsd/amd64/dit	\
	$(TARNAME)/darwin/amd64/dit	\
	$(TARNAME)/darwin/arm64/dit	\
	$(TARNAME)/windows/amd64/dit	
 
ditcli-$(SHA)/%/dit:
	mkdir -p `echo $@ | awk -F/ '{print $$1 FS $$2}'`
	go build -ldflags="-X 'main.sha=$(SHA)' -X 'main.buildDate=$(DATE)'" -o $@

tar: all $(TARNAME)/README $(TARNAME)/LICENSE
	tar czv $(TARNAME) > $(TARNAME).tar.gz

$(TARNAME)/README:
	cp README $(TARNAME)

$(TARNAME)/LICENSE:
	cp ../../LICENSE $(TARNAME)
