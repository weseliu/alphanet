@echo off

go install github.com/weseliu/alphanet/cmd/auth
if %errorlevel% == 0 (
    echo install auth success!
) else (
    goto error
)
go install github.com/weseliu/alphanet/cmd/connect
if %errorlevel% == 0 (
    echo install connect success!
) else (
    goto error
)
go install github.com/weseliu/alphanet/cmd/game
if %errorlevel% == 0 (
    echo install game success!
) else (
    goto error
)

xcopy F:\GoProjects\src\github.com\weseliu\alphanet\cmd\auth\conf\*.* F:\GoProjects\bin\conf\  /s /e /y
xcopy F:\GoProjects\src\github.com\weseliu\alphanet\cmd\connect\conf\*.* F:\GoProjects\bin\conf\  /s /e /y
xcopy F:\GoProjects\src\github.com\weseliu\alphanet\cmd\game\conf\*.* F:\GoProjects\bin\conf\  /s /e /y

cd F:\GoProjects\bin

start auth
if %errorlevel% == 0 (
    echo run auth success!
) else (
    goto error
)
start connect
if %errorlevel% == 0 (
    echo run connect success!
) else (
    goto error
)
start game
if %errorlevel% == 0 (
    echo run game success!
) else (
    goto error
)

goto success

:error
pause

:success
pause
