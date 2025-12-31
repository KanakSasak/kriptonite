package services

import (
	"Kryptonite/internal/core/domain"
	"Kryptonite/internal/core/ports"
	"fmt"
	"math"
)

type cryptService struct {
	crypt domain.Crypt
}

func (c2 cryptService) ModPow(a, b, c float64) float64 {
	return math.Mod(math.Pow(a, b), c)
}

func (c cryptService) Mod(a, b float64) float64 {
	return math.Mod(a, b)
}

func (c cryptService) Pow(a, b float64) float64 {
	return math.Pow(a, b)
}

func (c cryptService) Orde(a, b int) int {
	for i := 1; i <= b; i++ {
		fmt.Printf("value %d Mod %d\n", c.ModPow(float64(a), float64(i), float64(b)), b)
		if c.ModPow(float64(a), float64(i), float64(b)) == 1 {
			return i
		}
	}
	return -1
}

func (c cryptService) Euclidean(a, b int) int {
	fmt.Printf("gcd(%d, %d)\n", a, b)
	for b != 0 {
		a, b = b, a%b
		fmt.Printf("   = gcd(%d, %d)\n", a, b)
	}
	fmt.Printf("   = %d\n", a)
	return a
}

func (c cryptService) ExtEuclidean(a, b int) (int, int, int) {
	fmt.Printf("extendedGCD(%d, %d)\n", a, b)
	if b == 0 {
		return a, 1, 0
	}

	gcd, x1, y1 := c.ExtEuclidean(b, a%b)
	x := y1
	y := x1 - (a/b)*y1

	return gcd, x, y
}

func (c cryptService) GeneratePrime(min, max int) ([]int, error) {
	primes := make([]int, 0)
	for num := min; num <= max; num++ {
		if c.IsPrime(num) {
			primes = append(primes, num)
		}
	}
	return primes, nil
}

func (c cryptService) IsPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func (c cryptService) AreRelativePrimes(a, b int) bool {
	for i := 2; i <= a && i <= b; i++ {
		if a%i == 0 && b%i == 0 {
			return false
		}
	}
	return true
}

func (c cryptService) CalculateFactorial(n int) int {
	result := 1
	log := ""
	for i := 1; i <= n; i++ {
		result *= i
		log += fmt.Sprintf("%d", i)
		if i < n {
			log += " × "
		}
	}
	fmt.Sprintf("%d! = %s = %d", n, log, result)
	return result
}

func (c cryptService) CalculateFactors(n int) []int {
	factors := make([]int, 0)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func (c cryptService) CalculateEulerFunction(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if c.CalculateGCD(i, n) == 1 {
			count++
		}
	}
	return count
}

func (c cryptService) CalculateGCD(a, b int) int {
	return c.Euclidean(a, b)
}

// Abstract Algebra operations

// GenerateSetGroup generates the set group Z_n = {0, 1, 2, ..., n-1}
func (c cryptService) GenerateSetGroup(n int) []int {
	group := make([]int, n)
	for i := 0; i < n; i++ {
		group[i] = i
	}
	return group
}

// GenerateAdditiveCoset generates additive coset H + a = {h + a | h ∈ H} (mod n)
func (c cryptService) GenerateAdditiveCoset(subgroup []int, element, modulus int) []int {
	coset := make([]int, len(subgroup))
	for i, h := range subgroup {
		coset[i] = (h + element) % modulus
	}
	return coset
}

// GenerateMultiplicativeCoset generates multiplicative coset H * a = {h * a | h ∈ H} (mod n)
func (c cryptService) GenerateMultiplicativeCoset(subgroup []int, element, modulus int) []int {
	coset := make([]int, len(subgroup))
	for i, h := range subgroup {
		coset[i] = (h * element) % modulus
	}
	return coset
}

// GenerateAdditiveQuotientGroup generates quotient group Z_n / H for additive group
func (c cryptService) GenerateAdditiveQuotientGroup(n int, subgroupSize int) [][]int {
	if subgroupSize <= 0 || n%subgroupSize != 0 {
		return [][]int{}
	}

	// Generate subgroup H = {0, subgroupSize, 2*subgroupSize, ...}
	subgroup := make([]int, n/subgroupSize)
	for i := 0; i < len(subgroup); i++ {
		subgroup[i] = (i * subgroupSize) % n
	}

	// Generate cosets
	quotientGroup := make([][]int, 0)
	used := make(map[int]bool)

	for i := 0; i < n; i++ {
		if !used[i] {
			coset := c.GenerateAdditiveCoset(subgroup, i, n)
			quotientGroup = append(quotientGroup, coset)
			for _, elem := range coset {
				used[elem] = true
			}
		}
	}

	return quotientGroup
}

// GenerateMultiplicativeQuotientGroup generates quotient group for multiplicative group
func (c cryptService) GenerateMultiplicativeQuotientGroup(n int, subgroup []int) [][]int {
	if len(subgroup) == 0 {
		return [][]int{}
	}

	// Generate all coprime elements (multiplicative group Z_n*)
	group := make([]int, 0)
	for i := 1; i < n; i++ {
		if c.CalculateGCD(i, n) == 1 {
			group = append(group, i)
		}
	}

	// Generate cosets
	quotientGroup := make([][]int, 0)
	used := make(map[int]bool)

	for _, g := range group {
		if !used[g] {
			coset := c.GenerateMultiplicativeCoset(subgroup, g, n)
			quotientGroup = append(quotientGroup, coset)
			for _, elem := range coset {
				used[elem] = true
			}
		}
	}

	return quotientGroup
}

// GenerateAdditiveCyclicGroup generates cyclic group <g> = {0, g, 2g, 3g, ...} (mod n)
func (c cryptService) GenerateAdditiveCyclicGroup(generator, modulus int) []int {
	cyclicGroup := make([]int, 0)
	seen := make(map[int]bool)
	current := 0

	for {
		if seen[current] {
			break
		}
		cyclicGroup = append(cyclicGroup, current)
		seen[current] = true
		current = (current + generator) % modulus
	}

	return cyclicGroup
}

// GenerateMultiplicativeCyclicGroup generates cyclic group <g> = {1, g, g^2, g^3, ...} (mod n)
func (c cryptService) GenerateMultiplicativeCyclicGroup(generator, modulus int) []int {
	cyclicGroup := make([]int, 0)
	seen := make(map[int]bool)
	current := 1

	for {
		if seen[current] {
			break
		}
		cyclicGroup = append(cyclicGroup, current)
		seen[current] = true
		current = (current * generator) % modulus
	}

	return cyclicGroup
}

// FindGenerators finds all generators of the multiplicative group Z_n*
func (c cryptService) FindGenerators(n int) []int {
	generators := make([]int, 0)
	phi := c.CalculateEulerFunction(n)

	for g := 1; g < n; g++ {
		if c.CalculateGCD(g, n) == 1 {
			// Check if g is a generator by verifying its order equals phi(n)
			order := c.Orde(g, n)
			if order == phi {
				generators = append(generators, g)
			}
		}
	}

	return generators
}

// EvaluatePolynomial evaluates polynomial at x: a_0 + a_1*x + a_2*x^2 + ...
func (c cryptService) EvaluatePolynomial(coefficients []int, x int) int {
	if len(coefficients) == 0 {
		return 0
	}

	result := 0
	power := 1

	for _, coef := range coefficients {
		result += coef * power
		power *= x
	}

	return result
}

// AddPolynomials adds two polynomials
func (c cryptService) AddPolynomials(poly1, poly2 []int) []int {
	maxLen := len(poly1)
	if len(poly2) > maxLen {
		maxLen = len(poly2)
	}

	result := make([]int, maxLen)

	for i := 0; i < maxLen; i++ {
		if i < len(poly1) {
			result[i] += poly1[i]
		}
		if i < len(poly2) {
			result[i] += poly2[i]
		}
	}

	return result
}

// MultiplyPolynomials multiplies two polynomials
func (c cryptService) MultiplyPolynomials(poly1, poly2 []int) []int {
	if len(poly1) == 0 || len(poly2) == 0 {
		return []int{}
	}

	result := make([]int, len(poly1)+len(poly2)-1)

	for i := 0; i < len(poly1); i++ {
		for j := 0; j < len(poly2); j++ {
			result[i+j] += poly1[i] * poly2[j]
		}
	}

	return result
}

func NewCryptService(crypt domain.Crypt) ports.Crypt {
	return &cryptService{
		crypt: crypt,
	}
}
