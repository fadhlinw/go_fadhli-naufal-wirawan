package Simple

import (
	"testing"
)

func TestAdditon(t *testing.T) {
	if Additon(1, 2) != 3 {
		t.Error("1 +2 Harusnya 3")
	}
	if Additon(-1, -2) != -3 {
		t.Error("-1 + -2 Harusnya -3")
	}
}

func TestSubstract(t *testing.T) {
	if Substract(1, 2) != -1 {
		t.Error("1 - 2 Harusnya -1")
	}
	if Substract(2, 1) != 1 {
		t.Error("-1 - -2 Harusnya 1")
	}
}

func TestMult(t *testing.T) {
	if Mult(1, 2) != 2 {
		t.Error("1 * 2 Harusnya 2")
	}
	if Mult(-1, -2) != 2 {
		t.Error("-1 * -2 Harusnya 2")
	}
}

func TestDiv(t *testing.T) {
	if Div(2, 1) != 2 {
		t.Error("2 / 1 Harusnya 2")
	}
	if Div(-2, -1) != 2 {
		t.Error("-2 * -1 Harusnya 2")
	}
}
