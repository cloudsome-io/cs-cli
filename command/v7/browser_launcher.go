package v7

import (
	"os/exec"
	"runtime"
)

// BrowserLauncher opens URLs in a browser.
type BrowserLauncher interface {
	Open(url string) error
}

// DefaultBrowserLauncher opens URLs using the OS default handler.
type DefaultBrowserLauncher struct{}

func (DefaultBrowserLauncher) Open(url string) error {
	var cmdName string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmdName = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmdName = "open"
		args = []string{url}
	default:
		cmdName = "xdg-open"
		args = []string{url}
	}

	return exec.Command(cmdName, args...).Start()
}

