@echo off

@REM Windows
SET GOOS=windows
go build -o ../build/nano-oj.exe
@REM Linux
SET CGO_ENABLED=0
SET GOARCH=amd64
SET GOOS=linux
go build -o ../build/nano-oj

echo Done.
