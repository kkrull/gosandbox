package greet

import "testing"

func Greet() string {
  return "Hello World!"
}

func TestGreet(t *testing.T) {
  expected := "Hello World!"
  actual := Greet()
  if actual != expected {
    t.Errorf("Expected <%v> but was <%v>", expected, actual)
  }
}

