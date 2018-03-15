package greet

import "testing"

func TestGreetWithNoNameShouldGreetTheWorld(t *testing.T) {
  assertEqual(t, "Hello World!", Greet())
}

func TestGreetWithOneNameShouldGreetThatPersonByName(t *testing.T) {
  assertEqual(t, "Hello George!", Greet("George"))
}

func TestGreetWithTwoOrMoreNamesListsThosePeople(t *testing.T) {
  assertEqual(t, "Hello George, Judy, and Astro!", Greet("George", "Judy", "Astro"))
}

func assertEqual(t *testing.T, expected string, actual string) {
  if actual != expected {
    t.Errorf("Expected <%v>, but was <%v>", expected, actual)
  }
}
