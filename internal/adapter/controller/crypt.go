package controller

import (
	"Kryptonite/internal/core"
	"Kryptonite/internal/core/ports"
)

type cryptController struct {
	c *core.Container
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
