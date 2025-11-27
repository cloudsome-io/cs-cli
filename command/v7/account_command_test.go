package v7_test

import (
	"errors"

	"code.cloudfoundry.org/cli/command/commandfakes"
	. "code.cloudfoundry.org/cli/command/v7"
	"code.cloudfoundry.org/cli/util/ui"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = Describe("AccountCommand", func() {
	const expectedAccountURL = "https://account.cloudsome.io/realms/cloudsome/account"
	var (
		cmd         AccountCommand
		testUI      *ui.UI
		fakeConfig  *commandfakes.FakeConfig
		fakeBrowser *fakeBrowserLauncher
		executeErr  error
	)

	BeforeEach(func() {
		testUI = ui.NewTestUI(nil, NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeBrowser = &fakeBrowserLauncher{}

		cmd = AccountCommand{
			BaseCommand: BaseCommand{
				UI:     testUI,
				Config: fakeConfig,
			},
			Browser: fakeBrowser,
		}
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(nil)
	})

	It("opens the account management page in the browser", func() {
		Expect(executeErr).NotTo(HaveOccurred())
		Expect(fakeBrowser.openArgs).To(Equal([]string{expectedAccountURL}))
		Expect(testUI.Out).To(Say(expectedAccountURL))
	})

	When("opening the browser fails", func() {
		BeforeEach(func() {
			fakeBrowser.err = errors.New("browser-missing")
		})

		It("warns the user and shows next steps", func() {
			Expect(executeErr).NotTo(HaveOccurred())
			Expect(testUI.Err).To(Say("browser-missing"))
			Expect(testUI.Out).To(Say("Please open the URL above in your browser to continue."))
		})
	})
})
