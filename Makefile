NAME=nano-oj
BINDIR=build
VERSION=1.0.0
BUILDTIME=$(shell date -u)
GOBUILD=cd judge && go mod tidy && go build -ldflags '-s -w -X "main.version=$(VERSION)" -X "main.buildTime=$(BUILDTIME)"'
FRONTBUILD=cd web && pnpm i && vite build --outDir=../$(BINDIR)/dist --emptyOutDir

.PHONY: web init

all: linux-amd64 windows-amd64 web

init:
	(cd judge && go mod tidy)
	(cd web && pnpm i)

push:
	git push zero --all

web:
	$(FRONTBUILD)

web-dev:
	cd web && pnpm i && vite dev --host

linux-amd64: 
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o ../$(BINDIR)/$(NAME)-$@-$(VERSION)

linux-amd64-dev: linux-amd64 web
	cd $(BINDIR) && ./$(NAME)-linux-amd64-$(VERSION)

windows-amd64: 
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o ../$(BINDIR)/$(NAME)-$@-$(VERSION).exe

clean:
	rm -rf $(BINDIR)/*
