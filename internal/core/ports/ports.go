package ports

type Crypt interface {
	GeneratePrime(min, max int) ([]int, error)
	IsPrime(num int) bool
	AreRelativePrimes(a, b int) bool
	CalculateFactorial(n int) int
	CalculateFactors(n int) []int
	CalculateEulerFunction(n int) int
	CalculateGCD(a, b int) int
}
