package services

import (
	"Kryptonite/internal/core/domain"
	"Kryptonite/internal/core/ports"
	"math"
)

type cryptService struct {
	crypt domain.Crypt
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
	//TODO implement me
	panic("implement me")
}

func (c cryptService) CalculateFactorial(n int) int {
	//TODO implement me
	panic("implement me")
}

func (c cryptService) CalculateFactors(n int) []int {
	//TODO implement me
	panic("implement me")
}

func (c cryptService) CalculateEulerFunction(n int) int {
	//TODO implement me
	panic("implement me")
}

func (c cryptService) CalculateGCD(a, b int) int {
	//TODO implement me
	panic("implement me")
}

func NewCryptService(crypt domain.Crypt) ports.Crypt {
	return &cryptService{
		crypt: crypt,
	}
}
