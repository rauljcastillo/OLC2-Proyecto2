package src

var a = [4][4]int{
	{STRING, 4, 4, 4},
	{4, FLOAT, FLOAT, 4},
	{4, FLOAT, INT, 4},
	{4, 4, 4, BOOL},
}

func TablaT(tipoA, tipoB int) int {
	return a[tipoA][tipoB]
}
