//annotated_greet_test.go -- test sources end in _test.go // HL
package annotated_greet

import "testing"

func Greet() string { //Function syntax // HL
	return "Hello World!"
}

func TestGreet(t *testing.T) { //Test functions start with "Test" and accept *testing.T // HL
	expected := "Hello World!" //Variable declaration and initialization // HL
	actual := Greet()
	if actual != expected {
		t.Errorf("Expected <%v> but was <%v>", expected, actual) //How to fail the test // HL
	}
}
