package handler

import (
	"Kryptonite/internal/adapter"
	"Kryptonite/internal/core/ports"
)

type cliHandler struct {
	CryptController ports.Crypt
}

func (c cliHandler) ModPow(a, b, c2 float64) float64 {
	data := c.CryptController.ModPow(a, b, c2)

	return data
}

func (c cliHandler) Mod(a, b float64) float64 {
	data := c.CryptController.Mod(a, b)

	return data
}

func (c cliHandler) Pow(a, b float64) float64 {
	data := c.CryptController.Pow(a, b)

	return data
}

func (c cliHandler) Orde(a, b int) int {
	data := c.CryptController.Orde(a, b)

	return data
}

func (c cliHandler) Euclidean(a, b int) int {
	data := c.CryptController.Euclidean(a, b)

	return data
}

func (c cliHandler) ExtEuclidean(a, b int) (int, int, int) {
	data1, data2, data3 := c.CryptController.ExtEuclidean(a, b)

	return data1, data2, data3
}

func (c cliHandler) GeneratePrime(min, max int) ([]int, error) {
	data, err := c.CryptController.GeneratePrime(min, max)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c cliHandler) IsPrime(num int) bool {
	data := c.IsPrime(num)

	return data
}

func (c cliHandler) AreRelativePrimes(a, b int) bool {
	data := c.CryptController.AreRelativePrimes(a, b)

	return data
}

func (c cliHandler) CalculateFactorial(n int) int {
	data := c.CryptController.CalculateFactorial(n)

	return data
}

func (c cliHandler) CalculateFactors(n int) []int {
	data := c.CryptController.CalculateFactors(n)

	return data
}

func (c cliHandler) CalculateEulerFunction(n int) int {
	data := c.CryptController.CalculateEulerFunction(n)
	return data
}

func (c cliHandler) CalculateGCD(a, b int) int {
	data := c.CryptController.CalculateGCD(a, b)
	return data
}

func NewCliServer(ctx *adapter.HandlerContext) ports.Crypt {
	return &cliHandler{
		CryptController: ctx.CryptController,
	}
}
