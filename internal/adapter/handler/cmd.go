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
	//TODO implement me
	panic("implement me")
}

func (c cliHandler) CalculateFactorial(n int) int {
	//TODO implement me
	panic("implement me")
}

func (c cliHandler) CalculateFactors(n int) []int {
	//TODO implement me
	panic("implement me")
}

func (c cliHandler) CalculateEulerFunction(n int) int {
	//TODO implement me
	panic("implement me")
}

func (c cliHandler) CalculateGCD(a, b int) int {
	//TODO implement me
	panic("implement me")
}

func NewCliServer(ctx *adapter.HandlerContext) ports.Crypt {
	return &cliHandler{
		CryptController: ctx.CryptController,
	}
}
