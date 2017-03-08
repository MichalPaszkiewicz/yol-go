@echo off
cls
set GOPATH=C:/work
echo building packages...
go build gogame
echo installing program...
go install hello
"bin/hello.exe"
pause