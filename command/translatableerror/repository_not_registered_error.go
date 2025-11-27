package translatableerror

type RepositoryNotRegisteredError struct {
	Name       string
	BinaryName string
}

func (RepositoryNotRegisteredError) Error() string {
	return "Plugin repository {{.Name}} not found.\nUse '{{.BinaryName}} list-plugin-repos' to list registered repos."
}

func (e RepositoryNotRegisteredError) Translate(translate func(string, ...interface{}) string) string {
	binaryName := e.BinaryName
	if binaryName == "" {
		binaryName = "cs"
	}

	return translate(e.Error(), map[string]interface{}{
		"Name":       e.Name,
		"BinaryName": binaryName,
	})
}
