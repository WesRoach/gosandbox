package mult

import "testing"

func TestMult(t *testing.T) {
	for _, test := range multTestCases {
		res := Mult(test.input, test.multiplier)
		if res != test.expected {
			t.Errorf("Mult(%v, %v), expected: %v, actual: %v", test.input, test.multiplier, test.expected, res)
		}
	}
}

func BenchmarkMult(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range multTestCases {
			Mult(test.input, test.multiplier)
		}
	}
}
