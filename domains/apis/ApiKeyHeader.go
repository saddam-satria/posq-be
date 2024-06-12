package apis

type ApiKeyHeader struct {
	XApiKey string `reqHeader:"x-api-key"`
}
