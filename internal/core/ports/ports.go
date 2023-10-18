package ports

type Crypt interface {
	GeneratePrime(min, max int) ([]int, error)
	IsPrime(num int) bool
	AreRelativePrimes(a, b int) bool
	CalculateFactorial(n int) int
	CalculateFactors(n int) []int
	CalculateEulerFunction(n int) int
	CalculateGCD(a, b int) int
	Mod(a, b float64) float64
	Pow(a, b float64) float64
	ModPow(a, b, c float64) float64
	Orde(a, b int) int
	Euclidean(a, b int) int
	ExtEuclidean(a, b int) (int, int, int)
}
