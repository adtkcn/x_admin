go mod tidy

xcopy  .\public /S /Y .\bin\linux\public\
xcopy  .\public /S /Y .\bin\windows\public\
@REM xcopy  .\public /S /Y .\bin\darwin\public\

@REM # 1 目标平台的体系架构（386、amd64、arm） 
set GOARCH=amd64
@REM #2 目标平台的操作系统（darwin、freebsd、linux、windows）
@REM set GOOS=linux
@REM go build -ldflags "-s -w"  -o ./bin/linux/x-admin
@REM .\upx.exe ./bin/linux/x-admin

@REM 打包window
set GOOS=windows
go build -ldflags "-s -w" -o ./bin/windows/x-admin.exe
.\upx.exe ./bin/windows/x-admin.exe

@REM 打包苹果darwin
@REM set GOOS=darwin
@REM go build -ldflags "-s -w" -o ./bin/darwin/x-admin
@REM .\upx.exe ./bin/darwin/x-admin
