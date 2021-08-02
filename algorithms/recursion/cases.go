package recursion

var casesFactorial = []struct {
	input int
	want  int
}{
	{0, 1},
	{1, 1},
	{5, 120},
	{8, 40320},
	{10, 3628800},
	{14, 87178291200},
}

var casesFibonacci = []struct {
	input int
	want  int
}{
	{0, 0},
	{1, 1},
	{2, 1},
	{5, 5},
	{8, 21},
	{12, 144},
	{25, 75025},
}

var casesReverse = []struct {
	input string
	want  string
}{
	{"keyboard", "draobyek"},
	{"level", "level"},
	{"recursion", "noisrucer"},
}
