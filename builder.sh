#!/bin/bash

# Almacenar el directorio actual
current_dir=$(pwd)

# Cambiar al directorio src
cd ./src || exit

# Iterar sobre cada carpeta dentro de src
for carpeta in */; do
    # Entrar a la carpeta
    cd "$carpeta" || exit
    
    # Ejecutar el comando
    GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X main.version=1.0.0 -X main.buildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ)" ./main.go
    
    # Volver al directorio anterior
    cd "$current_dir" || exit
done
