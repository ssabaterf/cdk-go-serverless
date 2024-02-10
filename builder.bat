@echo off
setlocal

cd .\src

rem Almacenar el directorio actual
set "current_dir=%cd%"
echo current_dir: %current_dir%
rem Cambiar al directorio src
rem Iterar sobre cada carpeta dentro de src
for /D %%F in (*) do (
    rem Entrar a la carpeta
    echo %%F
    cd "%%F"
    
    rem Ejecutar el comando
    set "GOOS=linux"
    set "GOARCH=amd64"
    go build -ldflags "-s -w -X main.version=1.0.0 -X main.buildTime=%date:~10,4%-%date:~4,2%-%date:~7,2%T%time:~0,2%:%time:~3,2%:%time:~6,2%Z" .\main.go
    
    rem Volver al directorio anterior
    cd "%current_dir%"
)

endlocal
