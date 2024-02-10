include .env

.PHONY: all
all: build compress update clean

light:env
	@go build -o ./$(OUT_BIN) ./main.go

init: policy compress
	@aws lambda create-function --function-name ${FUNCTION_NAME} --runtime go1.x --role ${LAMBDA_ROLE} --handler ${HANDLER} --zip-file fileb://${ZIP_FILE}

policy:
	@aws iam create-role --role-name lambda-execution-role --assume-role-policy-document file://policy.json

update:
	@aws lambda update-function-code --function-name ${FUNCTION_NAME} --zip-file fileb://${ZIP_FILE}

build:
	set GOOS=linux&& set GOARCH=amd64&&go build -ldflags "-s -w -X main.version=1.0.0 -X main.buildTime=$(date -u+%Y-%m-%dT%H:%M:%SZ)" ./main.go

compress: build
	powershell Compress-Archive -Force ${OUT_BIN} ${ZIP_FILE}

clean:
	del $(OUT_BIN) $(ZIP_FILE)
