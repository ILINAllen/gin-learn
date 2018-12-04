@echo off
set curdir=%~dp0
echo curdir=%curdir%
set gopath=%gopath%;%curdir%;C:\Go\
echo gopath=%gopath%

set gobin=%curdir%bin
echo gobin=%gobin%
