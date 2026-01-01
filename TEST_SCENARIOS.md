# Kriptonite Test Scenarios

This document provides comprehensive test scenarios for all features in Kriptonite, including the newly implemented Abstract Algebra operations.

## Table of Contents
1. [Basic Number Theory Operations](#basic-number-theory-operations)
2. [Abstract Algebra Operations](#abstract-algebra-operations)
3. [Cryptographic Algorithm Simulations](#cryptographic-algorithm-simulations)
4. [Expected Test Results](#expected-test-results)

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

## Cryptographic Algorithm Simulations

This section demonstrates how to use Kriptonite's number theory and abstract algebra features to simulate and understand classic cryptographic algorithms.

### 1. RSA Encryption/Decryption Simulation

**RSA Algorithm Overview:**
- Public key cryptography system
- Based on difficulty of factoring large numbers
- Uses modular exponentiation

**Step-by-Step RSA Simulation:**

#### Step 1: Choose Two Prime Numbers (p, q)
```bash
# Generate small primes for demonstration
./kriptonite prime 10 30
# Choose two primes, e.g., p = 11, q = 13
```

#### Step 2: Calculate n = p × q
```bash
# n = 11 × 13 = 143
# Verify: 11 * 13 = 143
# n becomes part of both public and private keys
```

#### Step 3: Calculate φ(n) = (p-1)(q-1)
```bash
# φ(143) = φ(11 × 13) = φ(11) × φ(13) = 10 × 12 = 120
./kriptonite euler 143
# Expected: This will calculate φ(143) = 120
```

**Mathematical Note:** For RSA, we use φ(n) = (p-1)(q-1) instead of calculating directly.

#### Step 4: Choose Public Exponent e
```bash
# e must be coprime to φ(n) = 120
# Common choice: e = 7 (small, coprime to 120)

# Verify e is coprime to φ(n)
./kriptonite relprime 7 120
# Expected: true (gcd(7, 120) = 1)

# Alternative: Try e = 5
./kriptonite relprime 5 120
# Expected: false (gcd(5, 120) = 5, not coprime)
```

#### Step 5: Calculate Private Exponent d
```bash
# Find d such that: d × e ≡ 1 (mod φ(n))
# Using Extended Euclidean Algorithm
./kriptonite exeuclid 7 120

# This gives: gcd, x, y where 7x + 120y = gcd
# We need d such that: 7d ≡ 1 (mod 120)
# d = x (mod 120), adjusted to be positive
# Expected output will show x value, d = 103 (since 7 × 103 = 721 = 6 × 120 + 1)
```

**Manual Verification:**
```bash
# Verify: (7 × 103) mod 120 = 1
./kriptonite mod 721 120
# Expected: 1 ✓
```

#### Step 6: Encrypt a Message
```bash
# Public Key: (e=7, n=143)
# Message: m = 9 (must be < n)
# Ciphertext: c = m^e mod n = 9^7 mod 143

./kriptonite modpow 9 7 143
# Expected: c = 48
```

#### Step 7: Decrypt the Ciphertext
```bash
# Private Key: (d=103, n=143)
# Ciphertext: c = 48
# Plaintext: m = c^d mod n = 48^103 mod 143

./kriptonite modpow 48 103 143
# Expected: m = 9 (original message recovered!) ✓
```

**Complete RSA Example:**
```bash
# Setup
p=11, q=13
n=143 (public)
φ(n)=120 (private)
e=7 (public exponent)
d=103 (private exponent)

# Public Key: (7, 143)
# Private Key: (103, 143)

# Encryption: m=9
./kriptonite modpow 9 7 143
# c = 48

# Decryption: c=48
./kriptonite modpow 48 103 143
# m = 9 ✓
```

**Larger RSA Example:**
```bash
# Using larger primes
p=17, q=19
n=323
φ(n)=288

# Choose e=5, verify coprimality
./kriptonite relprime 5 288
# true

# Find d using extended Euclidean
./kriptonite exeuclid 5 288
# Calculate d from result (d=173)

# Encrypt message m=42
./kriptonite modpow 42 5 323
# c = 75

# Decrypt c=75
./kriptonite modpow 75 173 323
# m = 42 ✓
```

---

### 2. Elliptic Curve Cryptography (ECC) Concepts

**ECC Overview:**
- Based on elliptic curves over finite fields
- Point addition and scalar multiplication
- Curve equation: y² = x³ + ax + b (mod p)

**Demonstrating ECC Concepts with Kriptonite:**

#### Step 1: Choose a Prime Field
```bash
# Use prime p for the finite field F_p
./kriptonite isprime 23
# true - We'll use F_23
```

#### Step 2: Define Elliptic Curve Parameters
```bash
# Curve: y² = x³ + x + 1 (mod 23)
# Parameters: a=1, b=1, p=23

# Points on the curve satisfy: y² ≡ x³ + x + 1 (mod 23)
```

#### Step 3: Verify Point on Curve
```bash
# Check if point (3, 10) is on the curve
# Left side: y² mod p
./kriptonite modpow 10 2 23
# 100 mod 23 = 8

# Right side: (x³ + x + 1) mod p
# x³ = 3³ = 27
./kriptonite modpow 3 3 23
# 27 mod 23 = 4
# Add: 4 + 3 + 1 = 8
./kriptonite mod 8 23
# 8 ✓

# Point (3, 10) is on the curve since 8 ≡ 8 (mod 23)
```

#### Step 4: Order of Elliptic Curve Group
```bash
# The order of E(F_23) for y² = x³ + x + 1
# For small fields, we can find generators

# Example: Using base point G and finding its order
# This demonstrates the cyclic group structure
./kriptonite generators 23
# Shows generators of multiplicative group Z_23*
```

#### Step 5: Demonstrate Scalar Multiplication (Simplified)
```bash
# Scalar multiplication: kG where k is scalar, G is base point
# In practice: repeatedly add point to itself k times

# Example: 5G means G + G + G + G + G
# This uses the group operation (point addition on the curve)

# For demonstration with cyclic groups:
./kriptonite cyclicgroup multiplicative 2 23
# Shows <2> in Z_23*, analogous to scalar multiplication
```

**ECC Key Generation Simulation:**
```bash
# Private key: Random integer d (1 < d < n, where n is order)
# Public key: Q = dG (scalar multiplication of base point)

# Example with d=7, using generator in F_23
./kriptonite modpow 2 7 23
# Shows concept of "raising to power" similar to scalar mult
```

---

### 3. Digital Signature Algorithm (DSA) Simulation

**DSA Overview:**
- Used for digital signatures (authentication, non-repudiation)
- Based on discrete logarithm problem
- Uses modular arithmetic and group theory

**Step-by-Step DSA Simulation:**

#### Step 1: Choose Domain Parameters
```bash
# Choose prime p and prime q where q divides (p-1)
# Example: p=23, q=11 where 11 divides 22

# Verify q is prime
./kriptonite isprime 11
# true

# Verify q divides (p-1)
./kriptonite mod 22 11
# 0 ✓ (22 is divisible by 11)
```

#### Step 2: Find Generator g of Order q
```bash
# Find g such that g^q ≡ 1 (mod p)
# Generator of subgroup of order q in Z_p*

# Find generators of Z_23*
./kriptonite generators 23
# Shows generators [5, 7, 10, 11, 14, 15, 17, 19, 20, 21]

# For subgroup of order 11, we need element with order 11
# Check order of element 2
./kriptonite orde 2 23
# If order is 11, then g=2 works

# Alternative: Use generator and raise to (p-1)/q
# h = generator, g = h^((p-1)/q) mod p
# g = 2^(22/11) = 2^2 mod 23 = 4
./kriptonite modpow 2 2 23
# g = 4

# Verify: g^q mod p should be 1
./kriptonite modpow 4 11 23
# Expected: 1 ✓
```

#### Step 3: Generate Private Key
```bash
# Private key x: random integer (0 < x < q)
# Example: x = 6 (chosen randomly from 1 to 10)
x=6
```

#### Step 4: Generate Public Key
```bash
# Public key y = g^x mod p
./kriptonite modpow 4 6 23
# y = 13 (public key)

# Verify
./kriptonite modpow 4 6 23
# Expected: 13
```

#### Step 5: Signing Process (Simplified)
```bash
# To sign message m:
# 1. Choose random k (0 < k < q)
k=7

# 2. Calculate r = (g^k mod p) mod q
./kriptonite modpow 4 7 23
# = 16
./kriptonite mod 16 11
# r = 5

# 3. Calculate s = k^(-1) × (H(m) + x×r) mod q
#    (Simplified: H(m) = message hash, we'll use m=3 for demo)
m=3

# First: x × r mod q = 6 × 5 mod 11
./kriptonite mod 30 11
# = 8

# Second: (m + x×r) mod q = (3 + 8) mod 11
./kriptonite mod 11 11
# = 0 (need to choose different k or m)

# Try again with k=5
./kriptonite modpow 4 5 23
# = 1
./kriptonite mod 1 11
# r = 1

# Calculate x × r = 6 × 1 = 6
# m + x×r = 3 + 6 = 9
# Need k^(-1) mod 11 for k=5
./kriptonite exeuclid 5 11
# Find multiplicative inverse

# Signature: (r, s)
```

#### Step 6: Verification Process
```bash
# To verify signature (r, s) on message m with public key y:
# 1. Calculate w = s^(-1) mod q
# 2. Calculate u1 = (H(m) × w) mod q
# 3. Calculate u2 = (r × w) mod q
# 4. Calculate v = ((g^u1 × y^u2) mod p) mod q
# 5. Signature is valid if v = r

# This demonstrates the mathematical verification
```

**Complete DSA Example (Simplified):**
```bash
# Domain Parameters:
p=23, q=11, g=4

# Key Generation:
x=6 (private)
./kriptonite modpow 4 6 23
y=13 (public)

# Sign message m=3 with k=7:
./kriptonite modpow 4 7 23  # g^k mod p
./kriptonite mod 16 11       # r = (g^k mod p) mod q
# r = 5

# Verify using public key y=13
```

---

### 4. Diffie-Hellman Key Exchange Simulation

**DH Overview:**
- Allows two parties to establish shared secret over insecure channel
- Based on discrete logarithm problem

**Step-by-Step Simulation:**

#### Step 1: Agree on Public Parameters
```bash
# Prime p and generator g
p=23
./kriptonite generators 23
# Choose g=5 (a generator of Z_23*)
g=5
```

#### Step 2: Alice's Private Key
```bash
# Alice chooses private key a
a=6

# Alice computes public key A = g^a mod p
./kriptonite modpow 5 6 23
# A = 8 (Alice's public key)
```

#### Step 3: Bob's Private Key
```bash
# Bob chooses private key b
b=15

# Bob computes public key B = g^b mod p
./kriptonite modpow 5 15 23
# B = 19 (Bob's public key)
```

#### Step 4: Compute Shared Secret
```bash
# Alice computes: s = B^a mod p
./kriptonite modpow 19 6 23
# s = 2

# Bob computes: s = A^b mod p
./kriptonite modpow 8 15 23
# s = 2

# Both get the same shared secret s=2 ✓
```

**Verification:**
```bash
# Shared secret: s = g^(a×b) mod p
# a × b = 6 × 15 = 90
./kriptonite modpow 5 90 23
# s = 2 ✓
```

---

### 5. Discrete Logarithm Problem Demonstration

**Problem:** Given g, p, and y = g^x mod p, find x

```bash
# Known: g=3, p=17, y=12
# Find x such that 3^x ≡ 12 (mod 17)

# Brute force approach using order calculation
./kriptonite cyclicgroup multiplicative 3 17
# Shows all powers of 3 mod 17
# Find which power gives 12

# We can verify solutions:
./kriptonite modpow 3 1 17  # = 3
./kriptonite modpow 3 2 17  # = 9
./kriptonite modpow 3 3 17  # = 10
./kriptonite modpow 3 4 17  # = 13
./kriptonite modpow 3 5 17  # = 5
./kriptonite modpow 3 6 17  # = 15
./kriptonite modpow 3 7 17  # = 11
./kriptonite modpow 3 8 17  # = 16
./kriptonite modpow 3 9 17  # = 14
./kriptonite modpow 3 10 17 # = 8
./kriptonite modpow 3 11 17 # = 7
./kriptonite modpow 3 12 17 # = 4
./kriptonite modpow 3 13 17 # = 12 ✓

# Answer: x = 13
```

This demonstrates why discrete log is hard for large numbers!

---

### 6. Chinese Remainder Theorem (CRT) Application

**CRT in RSA Optimization:**

```bash
# RSA decryption can be sped up using CRT
# Given: c=48, d=103, n=143, p=11, q=13

# Compute dp = d mod (p-1)
./kriptonite mod 103 10
# dp = 3

# Compute dq = d mod (q-1)
./kriptonite mod 103 12
# dq = 7

# Compute mp = c^dp mod p
./kriptonite modpow 48 3 11
# mp = 9

# Compute mq = c^dq mod q
./kriptonite modpow 48 7 13
# mq = 9

# Both give 9, which is our original message! ✓
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
