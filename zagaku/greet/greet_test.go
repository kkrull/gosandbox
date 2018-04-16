package greet

import "testing" // HL

func Greet() string {
	return "Hello World!"
}

func TestGreet(t *testing.T) { //It's kind of like JUnit 3 // HL
	expected := "Hello World!"
	actual := Greet()
	if actual != expected {
		t.Errorf("Expected <%v> but was <%v>", expected, actual)
	}
}
