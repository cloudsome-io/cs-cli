package translatableerror

type ManualClientCredentialsError struct {
	BinaryName string
}

func (e ManualClientCredentialsError) Error() string {
	return "Error: Support for manually writing your client credentials to config.json has been removed. For similar functionality please use `{{.BinaryName}} auth --client-credentials`."
}

func (e ManualClientCredentialsError) Translate(translate func(string, ...interface{}) string) string {
	binaryName := e.BinaryName
	if binaryName == "" {
		binaryName = "cs"
	}

	return translate(e.Error(), map[string]interface{}{
		"BinaryName": binaryName,
	})
}
