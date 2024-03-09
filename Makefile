.PHONY: web

build: web server

web: 
	cd web && pnpm i && vite build --outDir ../build/static

server:
	cd judge && go build -o ../build/nano-oj
