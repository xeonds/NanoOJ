NAME=nano-oj
BINDIR=build
VERSION=1.1.0
BUILDTIME=$(shell date -u)
GOBUILD=cd judge && go mod tidy && go build -ldflags '-s -w -X "main.version=$(VERSION)" -X "main.buildTime=$(BUILDTIME)"'
FRONTBUILD=cd web && pnpm i && vite build --outDir=../$(BINDIR)/dist --emptyOutDir

.PHONY: web init

all: linux-amd64 windows-amd64 web

web:
	$(FRONTBUILD)

linux-amd64: 
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o ../$(BINDIR)/$(NAME)-$@

windows-amd64: 
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o ../$(BINDIR)/$(NAME)-$@.exe

dev:
	(cd $(BINDIR) && ./$(NAME)-linux-amd64) & \
	(cd web && pnpm i && vite dev --host --port 8080)

run:
	cd $(BINDIR) && ./$(NAME)-linux-amd64

deploy: init web linux-amd64
	docker-compose up -d

init:
	(cd judge && go mod tidy) && \
	(cd web && pnpm i)

push:
	git push zero --all

clean:
	rm -rf $(BINDIR)/$(NAME)-* $(BINDIR)/dist
