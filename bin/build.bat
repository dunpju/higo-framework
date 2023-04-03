@echo off

:a
if "%1" == "" (
    set /p compile=Please enter the compile file [eg: main.go]:
) else (
    set compile=%1%
)
if "%compile%" == "" (
    if "%1" == "" (
        goto a
    )
)

:b
if "%2%" == "" (
    if "%compile%" == "" (
        set /p executable=Please enter the executable file name [eg: main]:
    ) else (
        set executable=%compile:.go=%
    )
) else (
    set executable=%2%
)
if "%executable%" == "" (
    if "%2" == "" (
        goto b
    )
)

echo %compile%
echo %executable%
set GOOS=linux
set GOARCH=amd64
go build -o build/%executable% %compile%
echo Compile the complete!