package greet

import "testing"

func TestGreetWithNoNameShouldGreetTheWorld(t *testing.T) {
  expected := "Hello World!"
  actual := Greet()
  if actual != expected {
    t.Errorf("Expected <%v> but was <%v>", expected, actual)
  }
}

func TestGreetWithOneNameShouldGreetThatPersonByName(t *testing.T) {
  expected := "Hello George!"
  actual := Greet("George")
  if actual != expected {
    t.Errorf("Expected <%v> but was <%v>", expected, actual)
  }
}

func TestGreetWithTwoOrMoreNamesListsThosePeople(t *testing.T) {
  expected := "Hello George, Judy, and Astro!"
  actual := Greet("George", "Judy", "Astro")
  if actual != expected {
    t.Errorf("Expected <%v> but was <%v>", expected, actual)
  }
}

