package controller

import (
	"Kryptonite/internal/core"
	"Kryptonite/internal/core/ports"
)

type cryptController struct {
	c *core.Container
}

func (c cryptController) ModPow(a, b, c2 float64) float64 {
	data := c.c.CryptService.ModPow(a, b, c2)
	return data
}

func (c cryptController) Mod(a, b float64) float64 {
	data := c.c.CryptService.Mod(a, b)
	return data
}

func (c cryptController) Pow(a, b float64) float64 {
	data := c.c.CryptService.Pow(a, b)
	return data
}

func (c cryptController) Orde(a, b int) int {
	data := c.c.CryptService.Orde(a, b)
	return data
}

func (c cryptController) Euclidean(a, b int) int {
	data := c.c.CryptService.Euclidean(a, b)
	return data
}

func (c cryptController) ExtEuclidean(a, b int) (int, int, int) {
	data1, data2, data3 := c.c.CryptService.ExtEuclidean(a, b)
	return data1, data2, data3
}

func (c cryptController) GeneratePrime(min, max int) ([]int, error) {
	data, err := c.c.CryptService.GeneratePrime(min, max)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c cryptController) IsPrime(num int) bool {
	data := c.c.CryptService.IsPrime(num)
	return data
}

func (c cryptController) AreRelativePrimes(a, b int) bool {
	data := c.c.CryptService.AreRelativePrimes(a, b)
	return data
}

func (c cryptController) CalculateFactorial(n int) int {
	data := c.c.CryptService.CalculateFactorial(n)
	return data
}

func (c cryptController) CalculateFactors(n int) []int {
	data := c.c.CryptService.CalculateFactors(n)
	return data
}

func (c cryptController) CalculateEulerFunction(n int) int {
	data := c.c.CryptService.CalculateEulerFunction(n)
	return data
}

func (c cryptController) CalculateGCD(a, b int) int {
	data := c.c.CryptService.CalculateGCD(a, b)
	return data
}

func NewCryptController(c *core.Container) ports.Crypt {
	return &cryptController{c}
}
