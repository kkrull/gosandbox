package main_test

import (
	"testing"

	"github.com/cucumber/godog"
	. "github.com/kkrull/godog-hello"
	. "github.com/onsi/gomega"
)

var (
	thatGreeter  *Greeter
	thatGreeting string
)

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
		ScenarioInitializer: InitializeScenario,
	}
	RegisterTestingT(t)

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned from feature tests")
	}
}

func iHaveAGreeter() {
	thatGreeter = NewGreeter()
}

func thatGreeterGivesAGeneralGreeting() error {
	var err error
	thatGreeting, err = thatGreeter.Greet()
	return err
}

func thatGreetingShouldAddressEverybody() {
	Expect(thatGreeting).To(Equal("Hello World!"))
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Given(`^I have a Greeter$`, iHaveAGreeter)
	ctx.When(`^that greeter gives a general greeting$`, thatGreeterGivesAGeneralGreeting)
	ctx.Then(`^that greeting should address everybody$`, thatGreetingShouldAddressEverybody)
}
