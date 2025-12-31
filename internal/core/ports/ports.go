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

	// Abstract Algebra operations
	GenerateSetGroup(n int) []int
	GenerateAdditiveCoset(subgroup []int, element, modulus int) []int
	GenerateMultiplicativeCoset(subgroup []int, element, modulus int) []int
	GenerateAdditiveQuotientGroup(n int, subgroupSize int) [][]int
	GenerateMultiplicativeQuotientGroup(n int, subgroup []int) [][]int
	GenerateAdditiveCyclicGroup(generator, modulus int) []int
	GenerateMultiplicativeCyclicGroup(generator, modulus int) []int
	FindGenerators(n int) []int
	EvaluatePolynomial(coefficients []int, x int) int
	AddPolynomials(poly1, poly2 []int) []int
	MultiplyPolynomials(poly1, poly2 []int) []int
}
