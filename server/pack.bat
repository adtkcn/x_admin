go install github.com/swaggo/swag/cmd/swag@latest
go mod tidy
swag fmt
swag init

xcopy  .\public /S /Y .\dist\bin\linux\public\
xcopy  .\public /S /Y .\dist\bin\windows\public\

xcopy ".env" ".\dist\bin\windows\" /Y /I

@REM xcopy  .\public /S /Y .\dist/bin\darwin\public\

@REM # 1 目标平台的体系架构（386、amd64、arm） 
set GOARCH=amd64
@REM #2 目标平台的操作系统（darwin、freebsd、linux、windows）
set GOOS=linux
go build -ldflags "-s -w"  -o ./dist/bin/linux/x-admin
.\upx.exe ./dist/bin/linux/x-admin

@REM 打包window
set GOOS=windows
go build -ldflags "-s -w" -o ./dist/bin/windows/x-admin.exe
.\upx.exe ./dist/bin/windows/x-admin.exe

@REM 打包苹果darwin
@REM set GOOS=darwin
@REM go build -ldflags "-s -w" -o ./dist/bin/darwin/x-admin
@REM .\upx.exe ./dist/bin/darwin/x-admin
