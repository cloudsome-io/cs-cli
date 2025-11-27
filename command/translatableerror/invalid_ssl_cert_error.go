package translatableerror

type InvalidSSLCertError struct {
	URL              string
	SuggestedCommand string
	BinaryName       string
}

func (InvalidSSLCertError) Error() string {
	return "Invalid SSL Cert for {{.API}}\nTIP: Use '{{.BinaryName}} {{.SuggestedCommand}} --skip-ssl-validation' to continue with an insecure API endpoint"
}

func (e InvalidSSLCertError) Translate(translate func(string, ...interface{}) string) string {
	binaryName := e.BinaryName
	if binaryName == "" {
		binaryName = "cs"
	}

	return translate(e.Error(), map[string]interface{}{
		"API":              e.URL,
		"SuggestedCommand": e.SuggestedCommand,
		"BinaryName":       binaryName,
	})
}
