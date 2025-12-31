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

// Abstract Algebra operations

func (c cryptController) GenerateSetGroup(n int) []int {
	return c.c.CryptService.GenerateSetGroup(n)
}

func (c cryptController) GenerateAdditiveCoset(subgroup []int, element, modulus int) []int {
	return c.c.CryptService.GenerateAdditiveCoset(subgroup, element, modulus)
}

func (c cryptController) GenerateMultiplicativeCoset(subgroup []int, element, modulus int) []int {
	return c.c.CryptService.GenerateMultiplicativeCoset(subgroup, element, modulus)
}

func (c cryptController) GenerateAdditiveQuotientGroup(n int, subgroupSize int) [][]int {
	return c.c.CryptService.GenerateAdditiveQuotientGroup(n, subgroupSize)
}

func (c cryptController) GenerateMultiplicativeQuotientGroup(n int, subgroup []int) [][]int {
	return c.c.CryptService.GenerateMultiplicativeQuotientGroup(n, subgroup)
}

func (c cryptController) GenerateAdditiveCyclicGroup(generator, modulus int) []int {
	return c.c.CryptService.GenerateAdditiveCyclicGroup(generator, modulus)
}

func (c cryptController) GenerateMultiplicativeCyclicGroup(generator, modulus int) []int {
	return c.c.CryptService.GenerateMultiplicativeCyclicGroup(generator, modulus)
}

func (c cryptController) FindGenerators(n int) []int {
	return c.c.CryptService.FindGenerators(n)
}

func (c cryptController) EvaluatePolynomial(coefficients []int, x int) int {
	return c.c.CryptService.EvaluatePolynomial(coefficients, x)
}

func (c cryptController) AddPolynomials(poly1, poly2 []int) []int {
	return c.c.CryptService.AddPolynomials(poly1, poly2)
}

func (c cryptController) MultiplyPolynomials(poly1, poly2 []int) []int {
	return c.c.CryptService.MultiplyPolynomials(poly1, poly2)
}

func NewCryptController(c *core.Container) ports.Crypt {
	return &cryptController{c}
}
