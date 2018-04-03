cd /d %~dp0

protoc --go_out=..\..\protocal\connect  connect.proto

pause