package handler

import (
	"Kryptonite/internal/adapter"
	"Kryptonite/internal/core/ports"
)

type cliHandler struct {
	CryptController ports.Crypt
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
