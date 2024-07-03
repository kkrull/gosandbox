package main_test

import (
	"testing"

	"github.com/cucumber/godog"
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

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned from feature tests")
	}
}

func iHaveAGreeter() error {
	return godog.ErrPending
}

func thatGreeterGivesAGeneralGreeting() error {
	return godog.ErrPending
}

func thatGreetingShouldAddressEverybody() error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Given(`^I have a Greeter$`, iHaveAGreeter)
	ctx.When(`^that greeter gives a general greeting$`, thatGreeterGivesAGeneralGreeting)
	ctx.Then(`^that greeting should address everybody$`, thatGreetingShouldAddressEverybody)
}
