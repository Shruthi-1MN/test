package mathutil

import "testing"

func TestCalculator_Add(t *testing.T){
	calc := Calculator{}
	result := calc.Add(1, 1)
	expected := 2
	if result != expected{
		t.Errorf("Expected %d, but got %d", expected, result)
	}

}

func TestCalculator_Subtract(t *testing.T){
	calc := Calculator{}
	result := calc.Subtract(3, 2)
	expected := 1
	if result != expected{
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestCalculator_SquareRoot(t * testing.T){
	calc := Calculator{}
	result := calc.SquareRoot(25)
	expected := 5.0
	if result != expected{
		t.Errorf("Expected %.2f, but got %.2f", expected, result)
	}

}