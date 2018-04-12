
go install github.com/weseliu/alphanet/cmd/auth
go install github.com/weseliu/alphanet/cmd/connect
go install github.com/weseliu/alphanet/cmd/game

xcopy F:\GoProjects\src\github.com\weseliu\alphanet\cmd\auth\conf\*.* F:\GoProjects\bin\conf\  /s /e /y
xcopy F:\GoProjects\src\github.com\weseliu\alphanet\cmd\connect\conf\*.* F:\GoProjects\bin\conf\  /s /e /y

cd F:\GoProjects\bin

start auth
start connect
start game

pause
