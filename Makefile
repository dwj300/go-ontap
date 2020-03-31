default: netapp.yaml
	GO_POST_PROCESS_FILE="/usr/local/bin/gofmt -w" openapi-generator generate -i netapp.yaml -g go -o . --package-name ontap
