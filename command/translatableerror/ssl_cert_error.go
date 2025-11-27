package translatableerror

type SSLCertError struct {
	Message    string
	BinaryName string
}

func (SSLCertError) Error() string {
	return "SSL Certificate Error {{.Message}}\nTIP: Use '{{.BinaryName}} api --skip-ssl-validation' to continue with an insecure API endpoint"
}

func (e SSLCertError) Translate(translate func(string, ...interface{}) string) string {
	binaryName := e.BinaryName
	if binaryName == "" {
		binaryName = "cs"
	}

	return translate(e.Error(), map[string]interface{}{
		"Message":    e.Message,
		"BinaryName": binaryName,
	})
}
