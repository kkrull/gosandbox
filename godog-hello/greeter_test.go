package main_test

import (
	"testing"

	"github.com/cucumber/godog"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	//TODO KDK: Implement step definitions here with ctx.Given et al
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		Options: &godog.Options{
			Format: "pretty",
			Paths: []string{"features"},
			TestingT: t,
		},
		ScenarioInitializer: InitializeScenario,
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned from feature tests")
	}
}
