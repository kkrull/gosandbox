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
	ctx.Step(`^I have a Greeter$`, iHaveAGreeter)
	ctx.Step(`^that greeter gives a general greeting$`, thatGreeterGivesAGeneralGreeting)
	ctx.Step(`^that greeting should address everybody$`, thatGreetingShouldAddressEverybody)
}
