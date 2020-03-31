package math

import "testing"

func TestAdd(t *testing.T) {
	var a, b int = 100, 200
	var val = Add(a, b)

	if val != a+b {
		t.Error("Test Case [", "TestAdd", "] Failed!")
	}
}

func TestSubtract(t *testing.T) {
	var a, b int = 100, 200
	var val = Subtract(a, b)

	if val != a-b {
		t.Error("Test Case [", "TestSubtract", "] Failed!")
	}
}

func TestMultiply(t *testing.T) {
	var a, b int = 100, 200
	var val = Multiply(a, b)

	if val != a*b {
		t.Error("Test Case [", "TestMultiply", "] Failed!")
	}
}

func TestDivide(t *testing.T) {
	var a, b int = 100, 200
	var val = Divide(a, b)

	if val != a/b {
		t.Error("Test Case [", "TestDivide", "] Failed!")
	}
}
