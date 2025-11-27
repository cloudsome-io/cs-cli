package actionerror

import "fmt"

// InvalidCommandError represents an error that happens when help is called
// with an invalid command.
type InvalidCommandError struct {
	CommandName string
	BinaryName  string
}

func (err InvalidCommandError) Error() string {
	binaryName := err.BinaryName
	if binaryName == "" {
		binaryName = "cf"
	}

	return fmt.Sprintf("'%s' is not a registered command. See '%s help -a'", err.CommandName, binaryName)
}
