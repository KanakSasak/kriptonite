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
	//TODO implement me
	panic("implement me")
}

func (c cryptController) CalculateFactorial(n int) int {
	//TODO implement me
	panic("implement me")
}

func (c cryptController) CalculateFactors(n int) []int {
	//TODO implement me
	panic("implement me")
}

func (c cryptController) CalculateEulerFunction(n int) int {
	//TODO implement me
	panic("implement me")
}

func (c cryptController) CalculateGCD(a, b int) int {
	//TODO implement me
	panic("implement me")
}

func NewCryptController(c *core.Container) ports.Crypt {
	return &cryptController{c}
}
