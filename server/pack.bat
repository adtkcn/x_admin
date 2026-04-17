go install github.com/swaggo/swag/cmd/swag@latest
go mod tidy
swag fmt
swag init

xcopy  .\public /S /Y .\dist\bin\linux\public\
@REM xcopy  .\public /S /Y .\dist\bin\windows\public\
@REM xcopy  .\public /S /Y .\dist/bin\darwin\public\

@REM # 1 目标平台的体系架构（386、amd64、arm） 
set GOARCH=amd64
@REM #2 目标平台的操作系统（darwin、freebsd、linux、windows）
set GOOS=linux
go build -ldflags "-s -w"  -o ./dist/bin/linux/x_admin
.\upx.exe ./dist/bin/linux/x_admin

@REM 打包window
@REM set GOOS=windows
@REM go build -ldflags "-s -w" -o ./dist/bin/windows/x_admin.exe
@REM .\upx.exe ./dist/bin/windows/x_admin.exe

@REM 打包苹果darwin
@REM set GOOS=darwin
@REM go build -ldflags "-s -w" -o ./dist/bin/darwin/x_admin
@REM .\upx.exe ./dist/bin/darwin/x_admin
