
go install github.com/weseliu/alphanet/cmd/auth
go install github.com/weseliu/alphanet/cmd/connect

xcopy F:\GoProjects\src\github.com\weseliu\alphanet\cmd\conf\*.* F:\GoProjects\bin\conf\  /s /e

start auth
start connect

pause
