default: netapp.yaml
	openapi-generator generate -i netapp.yaml -g go -o . --package-name ontap
