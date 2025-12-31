# Kriptonite Test Scenarios

This document provides comprehensive test scenarios for all features in Kriptonite, including the newly implemented Abstract Algebra operations.

## Table of Contents
1. [Basic Number Theory Operations](#basic-number-theory-operations)
2. [Abstract Algebra Operations](#abstract-algebra-operations)
3. [Expected Test Results](#expected-test-results)

---

## Basic Number Theory Operations

### 1. Prime Number Generator
**Command:** `kriptonite prime <min> <max>`

```bash
# Test Case 1: Generate primes between 1 and 20
./kriptonite prime 1 20
# Expected: [2, 3, 5, 7, 11, 13, 17, 19]

# Test Case 2: Generate primes between 50 and 70
./kriptonite prime 50 70
# Expected: [53, 59, 61, 67]
```

### 2. Check Prime Number
**Command:** `kriptonite isprime <number>`

```bash
# Test Case 1: Check if 17 is prime
./kriptonite isprime 17
# Expected: true

# Test Case 2: Check if 24 is prime
./kriptonite isprime 24
# Expected: false
```

### 3. Relative Prime Check
**Command:** `kriptonite relprime <a> <b>`

```bash
# Test Case 1: Check if 8 and 15 are coprime
./kriptonite relprime 8 15
# Expected: true

# Test Case 2: Check if 12 and 18 are coprime
./kriptonite relprime 12 18
# Expected: false (gcd = 6)
```

### 4. Factorial Calculation
**Command:** `kriptonite factorial <n>`

```bash
# Test Case 1: Calculate 5!
./kriptonite factorial 5
# Expected: 120

# Test Case 2: Calculate 7!
./kriptonite factorial 7
# Expected: 5040
```

### 5. Factor Calculation
**Command:** `kriptonite factor <n>`

```bash
# Test Case 1: Find factors of 12
./kriptonite factor 12
# Expected: [1, 2, 3, 4, 6, 12]

# Test Case 2: Find factors of 28
./kriptonite factor 28
# Expected: [1, 2, 4, 7, 14, 28]
```

### 6. Modulus Operation
**Command:** `kriptonite mod <a> <b>`

```bash
# Test Case 1: 17 mod 5
./kriptonite mod 17 5
# Expected: 2

# Test Case 2: 100 mod 7
./kriptonite mod 100 7
# Expected: 2
```

### 7. Power Operation
**Command:** `kriptonite pow <base> <exponent>`

```bash
# Test Case 1: 2^8
./kriptonite pow 2 8
# Expected: 256

# Test Case 2: 3^4
./kriptonite pow 3 4
# Expected: 81
```

### 8. Modular Power
**Command:** `kriptonite modpow <base> <exponent> <modulus>`

```bash
# Test Case 1: 2^10 mod 13
./kriptonite modpow 2 10 13
# Expected: 1024 mod 13 = 10

# Test Case 2: 3^7 mod 11
./kriptonite modpow 3 7 11
# Expected: 2187 mod 11 = 9
```

### 9. Order Calculation
**Command:** `kriptonite orde <a> <n>`

```bash
# Test Case 1: Order of 2 in Z_7
./kriptonite orde 2 7
# Expected: 3 (since 2^3 ≡ 1 mod 7)

# Test Case 2: Order of 3 in Z_7
./kriptonite orde 3 7
# Expected: 6 (since 3^6 ≡ 1 mod 7, 3 is a generator)
```

### 10. Euler's Totient Function
**Command:** `kriptonite euler <n>`

```bash
# Test Case 1: φ(12)
./kriptonite euler 12
# Expected: 4 (numbers coprime to 12: 1, 5, 7, 11)

# Test Case 2: φ(7)
./kriptonite euler 7
# Expected: 6 (7 is prime, so φ(7) = 7-1)
```

### 11. Euclidean Algorithm
**Command:** `kriptonite euclid <a> <b>`

```bash
# Test Case 1: gcd(48, 18)
./kriptonite euclid 48 18
# Expected: 6

# Test Case 2: gcd(100, 35)
./kriptonite euclid 100 35
# Expected: 5
```

### 12. Extended Euclidean Algorithm
**Command:** `kriptonite exeuclid <a> <b>`

```bash
# Test Case 1: Extended GCD of 35 and 15
./kriptonite exeuclid 35 15
# Expected: gcd=5, x=1, y=-2 (where 35*1 + 15*(-2) = 5)

# Test Case 2: Extended GCD of 240 and 46
./kriptonite exeuclid 240 46
# Expected: gcd=2, x=-9, y=47
```

---

## Abstract Algebra Operations

### 1. Set Group Generation
**Command:** `kriptonite setgroup <n>`

**Description:** Generates the additive group Z_n = {0, 1, 2, ..., n-1}

```bash
# Test Case 1: Generate Z_5
./kriptonite setgroup 5
# Expected: [0, 1, 2, 3, 4]

# Test Case 2: Generate Z_10
./kriptonite setgroup 10
# Expected: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

# Test Case 3: Generate Z_3
./kriptonite setgroup 3
# Expected: [0, 1, 2]
```

### 2. Additive Coset Generation
**Command:** `kriptonite coset additive <subgroup> <element> <modulus>`

**Description:** Generates additive coset H + a = {h + a | h ∈ H} (mod n)

```bash
# Test Case 1: Z_9 coset
# Subgroup H = {0, 3, 6}, element = 1, modulus = 9
./kriptonite coset additive 0,3,6 1 9
# Expected: [1, 4, 7]

# Test Case 2: Z_9 coset
# Subgroup H = {0, 3, 6}, element = 2, modulus = 9
./kriptonite coset additive 0,3,6 2 9
# Expected: [2, 5, 8]

# Test Case 3: Z_6 coset
# Subgroup H = {0, 2, 4}, element = 1, modulus = 6
./kriptonite coset additive 0,2,4 1 6
# Expected: [1, 3, 5]
```

### 3. Multiplicative Coset Generation
**Command:** `kriptonite coset multiplicative <subgroup> <element> <modulus>`

**Description:** Generates multiplicative coset H * a = {h * a | h ∈ H} (mod n)

```bash
# Test Case 1: Z_7* coset
# Subgroup H = {1, 2, 4}, element = 3, modulus = 7
./kriptonite coset multiplicative 1,2,4 3 7
# Expected: [3, 6, 5]

# Test Case 2: Z_7* coset
# Subgroup H = {1, 2, 4}, element = 2, modulus = 7
./kriptonite coset multiplicative 1,2,4 2 7
# Expected: [2, 4, 1]

# Test Case 3: Z_11* coset
# Subgroup H = {1, 3, 9}, element = 2, modulus = 11
./kriptonite coset multiplicative 1,3,9 2 11
# Expected: [2, 6, 7]
```

### 4. Additive Quotient Group
**Command:** `kriptonite quotientgroup additive <n> <subgroup_size>`

**Description:** Generates quotient group Z_n / H where H is the subgroup of multiples of subgroup_size

```bash
# Test Case 1: Z_12 / H where |H| = 3
./kriptonite quotientgroup additive 12 3
# Expected:
# Coset 0: [0, 3, 6, 9]
# Coset 1: [1, 4, 7, 10]
# Coset 2: [2, 5, 8, 11]

# Test Case 2: Z_6 / H where |H| = 2
./kriptonite quotientgroup additive 6 2
# Expected:
# Coset 0: [0, 2, 4]
# Coset 1: [1, 3, 5]

# Test Case 3: Z_8 / H where |H| = 4
./kriptonite quotientgroup additive 8 4
# Expected:
# Coset 0: [0, 4]
# Coset 1: [1, 5]
# Coset 2: [2, 6]
# Coset 3: [3, 7]
```

### 5. Multiplicative Quotient Group
**Command:** `kriptonite quotientgroup multiplicative <n> <subgroup>`

**Description:** Generates quotient group Z_n* / H for multiplicative group

```bash
# Test Case 1: Z_7* / H where H = {1, 2, 4}
./kriptonite quotientgroup multiplicative 7 1,2,4
# Expected: 2 cosets
# Coset 0: elements from H
# Coset 1: elements from {3, 5, 6}

# Test Case 2: Z_13* / H where H = {1, 3, 9}
./kriptonite quotientgroup multiplicative 13 1,3,9
# Expected: 4 cosets
```

### 6. Additive Cyclic Group
**Command:** `kriptonite cyclicgroup additive <generator> <modulus>`

**Description:** Generates cyclic group <g> = {0, g, 2g, 3g, ...} (mod n)

```bash
# Test Case 1: <3> in Z_12
./kriptonite cyclicgroup additive 3 12
# Expected: [0, 3, 6, 9]
# Order: 4

# Test Case 2: <1> in Z_5
./kriptonite cyclicgroup additive 1 5
# Expected: [0, 1, 2, 3, 4]
# Order: 5

# Test Case 3: <2> in Z_6
./kriptonite cyclicgroup additive 2 6
# Expected: [0, 2, 4]
# Order: 3

# Test Case 4: <5> in Z_15
./kriptonite cyclicgroup additive 5 15
# Expected: [0, 5, 10]
# Order: 3
```

### 7. Multiplicative Cyclic Group
**Command:** `kriptonite cyclicgroup multiplicative <generator> <modulus>`

**Description:** Generates cyclic group <g> = {1, g, g^2, g^3, ...} (mod n)

```bash
# Test Case 1: <3> in Z_7
./kriptonite cyclicgroup multiplicative 3 7
# Expected: [1, 3, 2, 6, 4, 5]
# Order: 6 (3 is a generator of Z_7*)

# Test Case 2: <2> in Z_7
./kriptonite cyclicgroup multiplicative 2 7
# Expected: [1, 2, 4]
# Order: 3 (2 is not a generator)

# Test Case 3: <2> in Z_5
./kriptonite cyclicgroup multiplicative 2 5
# Expected: [1, 2, 4, 3]
# Order: 4

# Test Case 4: <2> in Z_13
./kriptonite cyclicgroup multiplicative 2 13
# Expected: [1, 2, 4, 8, 3, 6, 12, 11, 9, 5, 10, 7]
# Order: 12 (2 is a generator of Z_13*)
```

### 8. Find Generators
**Command:** `kriptonite generators <n>`

**Description:** Finds all generators of the multiplicative group Z_n*

```bash
# Test Case 1: Generators of Z_7*
./kriptonite generators 7
# Expected: [3, 5]
# Number of generators: 2

# Test Case 2: Generators of Z_5*
./kriptonite generators 5
# Expected: [2, 3]
# Number of generators: 2

# Test Case 3: Generators of Z_11*
./kriptonite generators 11
# Expected: [2, 6, 7, 8]
# Number of generators: 4

# Test Case 4: Generators of Z_13*
./kriptonite generators 13
# Expected: [2, 6, 7, 11]
# Number of generators: 4
```

**Note:** The number of generators of Z_p* (where p is prime) equals φ(φ(p)) = φ(p-1)

### 9. Polynomial Evaluation
**Command:** `kriptonite polynomial evaluate <coefficients> <x>`

**Description:** Evaluates polynomial P(x) = a₀ + a₁x + a₂x² + ... at given x

```bash
# Test Case 1: P(x) = 1 + 2x + 3x² at x = 5
./kriptonite polynomial evaluate 1,2,3 5
# Expected: P(5) = 1 + 2(5) + 3(25) = 1 + 10 + 75 = 86

# Test Case 2: P(x) = 5 at x = 10
./kriptonite polynomial evaluate 5 10
# Expected: P(10) = 5

# Test Case 3: P(x) = 1 + x + x² at x = 2
./kriptonite polynomial evaluate 1,1,1 2
# Expected: P(2) = 1 + 2 + 4 = 7

# Test Case 4: P(x) = 2x³ + 3x² + 4x + 5 at x = 3
./kriptonite polynomial evaluate 5,4,3,2 3
# Expected: P(3) = 5 + 12 + 27 + 54 = 98
```

### 10. Polynomial Addition
**Command:** `kriptonite polynomial add <poly1> <poly2>`

**Description:** Adds two polynomials

```bash
# Test Case 1: (1 + 2x) + (3 + 4x)
./kriptonite polynomial add 1,2 3,4
# Expected: 4 + 6x
# Coefficients: [4, 6]

# Test Case 2: (1 + 2x + 3x²) + (4 + 5x + 6x²)
./kriptonite polynomial add 1,2,3 4,5,6
# Expected: 5 + 7x + 9x²
# Coefficients: [5, 7, 9]

# Test Case 3: (1 + x) + (x²)
./kriptonite polynomial add 1,1 0,0,1
# Expected: 1 + x + x²
# Coefficients: [1, 1, 1]

# Test Case 4: (x² + 2x³) + (1 + x)
./kriptonite polynomial add 0,0,1,2 1,1
# Expected: 1 + x + x² + 2x³
# Coefficients: [1, 1, 1, 2]
```

### 11. Polynomial Multiplication
**Command:** `kriptonite polynomial multiply <poly1> <poly2>`

**Description:** Multiplies two polynomials

```bash
# Test Case 1: (1 + 2x) × (3 + 4x)
./kriptonite polynomial multiply 1,2 3,4
# Expected: 3 + 10x + 8x²
# Coefficients: [3, 10, 8]
# Explanation: 1×3 + (1×4 + 2×3)x + 2×4x²

# Test Case 2: x × x
./kriptonite polynomial multiply 0,1 0,1
# Expected: x²
# Coefficients: [0, 0, 1]

# Test Case 3: (1 + x) × (1 - x) (using 1,1 and 1,-1)
./kriptonite polynomial multiply 1,1 1,-1
# Expected: 1 - x²
# Coefficients: [1, 0, -1]

# Test Case 4: (x + 1) × (x² + x + 1)
./kriptonite polynomial multiply 1,1 1,1,1
# Expected: 1 + 2x + 2x² + x³
# Coefficients: [1, 2, 2, 1]
```

---

## Expected Test Results

### Unit Test Execution

Run all unit tests:
```bash
go test ./internal/core/services/... -v
```

Expected output:
```
=== RUN   TestCalculateEulerFunction
--- PASS: TestCalculateEulerFunction (0.00s)
=== RUN   TestGenerateSetGroup
--- PASS: TestGenerateSetGroup (0.00s)
=== RUN   TestGenerateAdditiveCoset
--- PASS: TestGenerateAdditiveCoset (0.00s)
=== RUN   TestGenerateMultiplicativeCoset
--- PASS: TestGenerateMultiplicativeCoset (0.00s)
=== RUN   TestGenerateAdditiveCyclicGroup
--- PASS: TestGenerateAdditiveCyclicGroup (0.00s)
=== RUN   TestGenerateMultiplicativeCyclicGroup
--- PASS: TestGenerateMultiplicativeCyclicGroup (0.00s)
=== RUN   TestFindGenerators
--- PASS: TestFindGenerators (0.00s)
=== RUN   TestEvaluatePolynomial
--- PASS: TestEvaluatePolynomial (0.00s)
=== RUN   TestAddPolynomials
--- PASS: TestAddPolynomials (0.00s)
=== RUN   TestMultiplyPolynomials
--- PASS: TestMultiplyPolynomials (0.00s)
PASS
ok      Kryptonite/internal/core/services
```

### Coverage Report

Generate coverage report:
```bash
go test ./internal/core/services/... -cover
```

Expected coverage: > 85%

---

## Common Test Patterns

### Testing Group Properties

1. **Closure**: All operations produce elements within the group
2. **Identity**: Group contains identity element (0 for additive, 1 for multiplicative)
3. **Inverses**: Every element has an inverse
4. **Associativity**: (a ⊕ b) ⊕ c = a ⊕ (b ⊕ c)

### Testing Cosets

1. **Disjoint cosets**: No overlap between different cosets
2. **Partition**: Cosets partition the entire group
3. **Same cardinality**: All cosets have the same number of elements

### Testing Generators

1. **Order verification**: Generator's order equals φ(n)
2. **Completeness**: <g> generates all elements of Z_n*
3. **Primitive root**: gᵏ ≢ 1 (mod n) for all 1 ≤ k < φ(n)

---

## Performance Benchmarks

For performance testing, you can benchmark polynomial operations:

```bash
go test -bench=. ./internal/core/services/...
```

Expected scenarios:
- Polynomial evaluation: O(n) where n is degree
- Polynomial addition: O(max(deg(p1), deg(p2)))
- Polynomial multiplication: O(n×m) where n, m are degrees
- Generator finding: O(n×φ(n)) where n is the modulus

---

## Notes

- All algebraic operations assume valid mathematical inputs
- Modulus operations work with positive integers
- Polynomial coefficients are in order: [a₀, a₁, a₂, ...] for a₀ + a₁x + a₂x² + ...
- For coset/quotient group commands, subgroups must be valid subgroups of the parent group
