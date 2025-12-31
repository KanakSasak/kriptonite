package services

import (
	"Kryptonite/internal/core/domain"
	"reflect"
	"testing"
)

func setupCryptService() *cryptService {
	data := domain.Crypt{Data: ""}
	svc := NewCryptService(data)
	return svc.(*cryptService)
}

// Test Euler Function
func TestCalculateEulerFunction(t *testing.T) {
	svc := setupCryptService()

	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"φ(1)", 1, 1},
		{"φ(2)", 2, 1},
		{"φ(3)", 3, 2},
		{"φ(4)", 4, 2},
		{"φ(5)", 5, 4},
		{"φ(6)", 6, 2},
		{"φ(7)", 7, 6},
		{"φ(8)", 8, 4},
		{"φ(9)", 9, 6},
		{"φ(10)", 10, 4},
		{"φ(12)", 12, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := svc.CalculateEulerFunction(tt.input)
			if result != tt.expected {
				t.Errorf("CalculateEulerFunction(%d) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

// Test GenerateSetGroup
func TestGenerateSetGroup(t *testing.T) {
	svc := setupCryptService()

	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{"Z_3", 3, []int{0, 1, 2}},
		{"Z_5", 5, []int{0, 1, 2, 3, 4}},
		{"Z_7", 7, []int{0, 1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := svc.GenerateSetGroup(tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("GenerateSetGroup(%d) = %v; want %v", tt.n, result, tt.expected)
			}
		})
	}
}

// Test GenerateAdditiveCoset
func TestGenerateAdditiveCoset(t *testing.T) {
	svc := setupCryptService()

	tests := []struct {
		name      string
		subgroup  []int
		element   int
		modulus   int
		expected  []int
	}{
		{"H + 1 mod 9", []int{0, 3, 6}, 1, 9, []int{1, 4, 7}},
		{"H + 2 mod 9", []int{0, 3, 6}, 2, 9, []int{2, 5, 8}},
		{"H + 1 mod 6", []int{0, 2, 4}, 1, 6, []int{1, 3, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := svc.GenerateAdditiveCoset(tt.subgroup, tt.element, tt.modulus)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("GenerateAdditiveCoset(%v, %d, %d) = %v; want %v",
					tt.subgroup, tt.element, tt.modulus, result, tt.expected)
			}
		})
	}
}

// Test GenerateMultiplicativeCoset
func TestGenerateMultiplicativeCoset(t *testing.T) {
	svc := setupCryptService()

	tests := []struct {
		name      string
		subgroup  []int
		element   int
		modulus   int
		expected  []int
	}{
		{"H * 2 mod 7", []int{1, 2, 4}, 2, 7, []int{2, 4, 1}},
		{"H * 3 mod 7", []int{1, 2, 4}, 3, 7, []int{3, 6, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := svc.GenerateMultiplicativeCoset(tt.subgroup, tt.element, tt.modulus)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("GenerateMultiplicativeCoset(%v, %d, %d) = %v; want %v",
					tt.subgroup, tt.element, tt.modulus, result, tt.expected)
			}
		})
	}
}

// Test GenerateAdditiveCyclicGroup
func TestGenerateAdditiveCyclicGroup(t *testing.T) {
	svc := setupCryptService()

	tests := []struct {
		name      string
		generator int
		modulus   int
		expected  []int
	}{
		{"<3> mod 12", 3, 12, []int{0, 3, 6, 9}},
		{"<1> mod 5", 1, 5, []int{0, 1, 2, 3, 4}},
		{"<2> mod 6", 2, 6, []int{0, 2, 4}},
		{"<4> mod 6", 4, 6, []int{0, 4, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := svc.GenerateAdditiveCyclicGroup(tt.generator, tt.modulus)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("GenerateAdditiveCyclicGroup(%d, %d) = %v; want %v",
					tt.generator, tt.modulus, result, tt.expected)
			}
		})
	}
}

// Test GenerateMultiplicativeCyclicGroup
func TestGenerateMultiplicativeCyclicGroup(t *testing.T) {
	svc := setupCryptService()

	tests := []struct {
		name      string
		generator int
		modulus   int
		expected  []int
	}{
		{"<3> mod 7", 3, 7, []int{1, 3, 2, 6, 4, 5}},
		{"<2> mod 7", 2, 7, []int{1, 2, 4}},
		{"<2> mod 5", 2, 5, []int{1, 2, 4, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := svc.GenerateMultiplicativeCyclicGroup(tt.generator, tt.modulus)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("GenerateMultiplicativeCyclicGroup(%d, %d) = %v; want %v",
					tt.generator, tt.modulus, result, tt.expected)
			}
		})
	}
}

// Test FindGenerators
func TestFindGenerators(t *testing.T) {
	svc := setupCryptService()

	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{"Generators of Z_7*", 7, []int{3, 5}},
		{"Generators of Z_5*", 5, []int{2, 3}},
		{"Generators of Z_11*", 11, []int{2, 6, 7, 8}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := svc.FindGenerators(tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("FindGenerators(%d) = %v; want %v", tt.n, result, tt.expected)
			}
		})
	}
}

// Test EvaluatePolynomial
func TestEvaluatePolynomial(t *testing.T) {
	svc := setupCryptService()

	tests := []struct {
		name         string
		coefficients []int
		x            int
		expected     int
	}{
		{"P(x) = 1 + 2x + 3x^2 at x=5", []int{1, 2, 3}, 5, 86},  // 1 + 10 + 75
		{"P(x) = 5 at x=10", []int{5}, 10, 5},
		{"P(x) = 1 + x at x=3", []int{1, 1}, 3, 4},
		{"P(x) = 2x at x=4", []int{0, 2}, 4, 8},
		{"P(x) = x^2 at x=3", []int{0, 0, 1}, 3, 9},
		{"Empty polynomial", []int{}, 5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := svc.EvaluatePolynomial(tt.coefficients, tt.x)
			if result != tt.expected {
				t.Errorf("EvaluatePolynomial(%v, %d) = %d; want %d",
					tt.coefficients, tt.x, result, tt.expected)
			}
		})
	}
}

// Test AddPolynomials
func TestAddPolynomials(t *testing.T) {
	svc := setupCryptService()

	tests := []struct {
		name     string
		poly1    []int
		poly2    []int
		expected []int
	}{
		{"(1+2x) + (3+4x)", []int{1, 2}, []int{3, 4}, []int{4, 6}},
		{"(1+2x+3x^2) + (4+5x+6x^2)", []int{1, 2, 3}, []int{4, 5, 6}, []int{5, 7, 9}},
		{"(1+x) + (x^2)", []int{1, 1}, []int{0, 0, 1}, []int{1, 1, 1}},
		{"Different lengths", []int{1, 2, 3, 4}, []int{5, 6}, []int{6, 8, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := svc.AddPolynomials(tt.poly1, tt.poly2)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("AddPolynomials(%v, %v) = %v; want %v",
					tt.poly1, tt.poly2, result, tt.expected)
			}
		})
	}
}

// Test MultiplyPolynomials
func TestMultiplyPolynomials(t *testing.T) {
	svc := setupCryptService()

	tests := []struct {
		name     string
		poly1    []int
		poly2    []int
		expected []int
	}{
		{"(1+2x) * (3+4x)", []int{1, 2}, []int{3, 4}, []int{3, 10, 8}},
		{"x * x", []int{0, 1}, []int{0, 1}, []int{0, 0, 1}},
		{"(2) * (3+x)", []int{2}, []int{3, 1}, []int{6, 2}},
		{"Empty poly1", []int{}, []int{1, 2}, []int{}},
		{"Empty poly2", []int{1, 2}, []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := svc.MultiplyPolynomials(tt.poly1, tt.poly2)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MultiplyPolynomials(%v, %v) = %v; want %v",
					tt.poly1, tt.poly2, result, tt.expected)
			}
		})
	}
}

// Test GenerateAdditiveQuotientGroup
func TestGenerateAdditiveQuotientGroup(t *testing.T) {
	svc := setupCryptService()

	tests := []struct {
		name          string
		n             int
		subgroupSize  int
		expectedCount int
	}{
		{"Z_12 / H where |H|=3", 12, 3, 3},
		{"Z_6 / H where |H|=2", 6, 2, 2},
		{"Invalid: n=12, |H|=5", 12, 5, 0},
		{"Invalid: n=10, |H|=0", 10, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := svc.GenerateAdditiveQuotientGroup(tt.n, tt.subgroupSize)
			if len(result) != tt.expectedCount {
				t.Errorf("GenerateAdditiveQuotientGroup(%d, %d) returned %d cosets; want %d",
					tt.n, tt.subgroupSize, len(result), tt.expectedCount)
			}
		})
	}
}

// Test GenerateMultiplicativeQuotientGroup
func TestGenerateMultiplicativeQuotientGroup(t *testing.T) {
	svc := setupCryptService()

	tests := []struct {
		name          string
		n             int
		subgroup      []int
		expectedCount int
	}{
		{"Z_7* / H where H={1,2,4}", 7, []int{1, 2, 4}, 2},
		{"Empty subgroup", 7, []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := svc.GenerateMultiplicativeQuotientGroup(tt.n, tt.subgroup)
			if len(result) != tt.expectedCount {
				t.Errorf("GenerateMultiplicativeQuotientGroup(%d, %v) returned %d cosets; want %d",
					tt.n, tt.subgroup, len(result), tt.expectedCount)
			}
		})
	}
}

// Test IsPrime (existing function test)
func TestIsPrime(t *testing.T) {
	svc := setupCryptService()

	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Prime: 2", 2, true},
		{"Prime: 3", 3, true},
		{"Prime: 5", 5, true},
		{"Prime: 7", 7, true},
		{"Prime: 11", 11, true},
		{"Not prime: 1", 1, false},
		{"Not prime: 4", 4, false},
		{"Not prime: 6", 6, false},
		{"Not prime: 8", 8, false},
		{"Not prime: 9", 9, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := svc.IsPrime(tt.input)
			if result != tt.expected {
				t.Errorf("IsPrime(%d) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// Test AreRelativePrimes
func TestAreRelativePrimes(t *testing.T) {
	svc := setupCryptService()

	tests := []struct {
		name     string
		a        int
		b        int
		expected bool
	}{
		{"Coprime: 3, 5", 3, 5, true},
		{"Coprime: 7, 11", 7, 11, true},
		{"Not coprime: 4, 6", 4, 6, false},
		{"Not coprime: 6, 9", 6, 9, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := svc.AreRelativePrimes(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("AreRelativePrimes(%d, %d) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// Test CalculateFactorial
func TestCalculateFactorial(t *testing.T) {
	svc := setupCryptService()

	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"0!", 0, 1},
		{"1!", 1, 1},
		{"3!", 3, 6},
		{"5!", 5, 120},
		{"6!", 6, 720},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := svc.CalculateFactorial(tt.input)
			if result != tt.expected {
				t.Errorf("CalculateFactorial(%d) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

// Test CalculateFactors
func TestCalculateFactors(t *testing.T) {
	svc := setupCryptService()

	tests := []struct {
		name     string
		input    int
		expected []int
	}{
		{"Factors of 6", 6, []int{1, 2, 3, 6}},
		{"Factors of 12", 12, []int{1, 2, 3, 4, 6, 12}},
		{"Factors of 7", 7, []int{1, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := svc.CalculateFactors(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("CalculateFactors(%d) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
