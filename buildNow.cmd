@echo off

cd %cd%

set GOPATH=%cd%


go build  main

echo build finished
echo start service
main.exe
pause /s