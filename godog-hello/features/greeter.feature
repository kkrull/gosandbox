Feature: Greeter
  In order to use Cucumber for Go
  As a developer
  I want to write a Hello World scenario in godog

  Scenario: Greeter's gotta greet
    Given I have a Greeter
    When that greeter gives a general greeting
    Then that greeting should address everybody
