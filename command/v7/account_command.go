package v7

import (
	"fmt"

	"code.cloudfoundry.org/cli/command"
)

const accountManagementURL = "https://account.cloudsome.io/realms/cloudsome/account"

// AccountCommand opens the Cloudsom account management page.
type AccountCommand struct {
	BaseCommand

	Browser BrowserLauncher

	usage interface{} `usage:"CF_NAME account"`
}

func (cmd *AccountCommand) Setup(config command.Config, ui command.UI) error {
	cmd.UI = ui
	cmd.Config = config
	cmd.Browser = DefaultBrowserLauncher{}
	return nil
}

func (cmd *AccountCommand) Execute(args []string) error {
	cmd.UI.DisplayText("Opening {{.URL}} in your browser to manage your Cloudsom account.", map[string]interface{}{
		"URL": accountManagementURL,
	})

	browser := cmd.Browser
	if browser == nil {
		browser = DefaultBrowserLauncher{}
	}

	if err := browser.Open(accountManagementURL); err != nil {
		cmd.UI.DisplayWarning(fmt.Sprintf("Unable to open browser automatically: %s", err))
		cmd.UI.DisplayText("Please open the URL above in your browser to continue.")
	}

	return nil
}
