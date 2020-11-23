package mult

type inputTest struct {
	input      int
	multiplier int
	expected   int
}

var multTestCases = []inputTest{
	{
		input:      1,
		multiplier: 2,
		expected:   2,
	},
	{
		input:      9,
		multiplier: 9,
		expected:   81,
	},
}
