GO ?= go

OUTDIR ?= $(CURDIR)/out
APPNAME ?= $(shell basename $(CURDIR))
SOURCES ?= 	$(wildcard *.go) \
			$(wildcard logutil/*.go) \
			Makefile go.mod go.sum

.PHONY: all
all: bin

.PHONY: clean
clean:
	@rm -rf $(OUTDIR)

$(OUTDIR):
	@mkdir -p $(OUTDIR)

.PHONY: bin
bin: $(OUTDIR)/$(APPNAME)

$(OUTDIR)/$(APPNAME): $(OUTDIR) $(SOURCES)
	@echo "building with $(GO)..."
	@$(GO) build -a -ldflags="-s -w" -o $@ main.go
	@ls -ahl $@

.PHONY: run
run:
	@$(GO) run main.go

.PHONY: info
info:
	@echo SOURCES=$(SOURCES)
	@echo GO=$(GO)
	@echo STRIP=$(STRIP)
	@echo OUTDIR=$(OUTDIR)
	@echo OUTDIR=$(OUTDIR)
	@echo APPNAME=$(APPNAME)

.PHONY: test
test:
	@go test github.com/leizongmin/dev-clean/...
