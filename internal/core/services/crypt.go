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
	//TODO implement me
	return 0
}

func (c cryptService) CalculateGCD(a, b int) int {
	return c.Euclidean(a, b)
}

func NewCryptService(crypt domain.Crypt) ports.Crypt {
	return &cryptService{
		crypt: crypt,
	}
}
