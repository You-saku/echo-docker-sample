package test

import (
	// テストのパッケージ
	"testing"
)

func Greet(str string) string {
	return "Hello"+str
}

func TestCalculate(t *testing.T) {
    result := 1+1
    if result != 2 {
        t.Errorf("1+1 = %d; want 2", result)
    }
}

func TestString(t *testing.T) {
    result := "HelloWorld";
    if Greet("World") != "HelloWorld" {
        t.Errorf("Greet(str) = %s; want World", result)
    }
}
