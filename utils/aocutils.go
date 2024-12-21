package utils

type Solution interface {
	Solve(isSample bool, dirName string) (string, string)
}
