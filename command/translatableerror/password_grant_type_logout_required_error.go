package translatableerror

type PasswordGrantTypeLogoutRequiredError struct {
	BinaryName string
}

func (PasswordGrantTypeLogoutRequiredError) Error() string {
	return "Service account currently logged in. Use '{{.BinaryName}} logout' to log out service account and try again."
}

func (e PasswordGrantTypeLogoutRequiredError) Translate(translate func(string, ...interface{}) string) string {
	binaryName := e.BinaryName
	if binaryName == "" {
		binaryName = "cs"
	}

	return translate(e.Error(), map[string]interface{}{
		"BinaryName": binaryName,
	})
}
