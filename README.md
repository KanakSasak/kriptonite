# Kriptonite

```bash
██╗  ██╗██████╗ ██╗██████╗ ████████╗ ██████╗ ███╗   ██╗██╗████████╗███████╗
██║ ██╔╝██╔══██╗██║██╔══██╗╚══██╔══╝██╔═══██╗████╗  ██║██║╚══██╔══╝██╔════╝
█████╔╝ ██████╔╝██║██████╔╝   ██║   ██║   ██║██╔██╗ ██║██║   ██║   █████╗  
██╔═██╗ ██╔══██╗██║██╔═══╝    ██║   ██║   ██║██║╚██╗██║██║   ██║   ██╔══╝  
██║  ██╗██║  ██║██║██║        ██║   ╚██████╔╝██║ ╚████║██║   ██║   ███████╗
╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝        ╚═╝    ╚═════╝ ╚═╝  ╚═══╝╚═╝   ╚═╝   ╚══════╝
                                                                           
```
### Educational Cryptography Toolkit

Welcome to Kriptonite! A comprehensive suite of utilities designed to assist cryptography students in mastering the mathematical foundations of modern cryptography. This toolkit provides hands-on experience with number theory, abstract algebra, and cryptographic algorithms.

## Features

### Number Theory Operations
1. Prime Number Generation
2. Prime Number Checking
3. Relative Prime (Coprime) Verification
4. Integer Factorization
5. Factorial Calculation
6. Modular Arithmetic
7. Power Calculation
8. Modular Exponentiation (with optimized algorithm)
9. GCD Calculation
10. Order Calculation
11. Euler's Totient Function (φ)
12. Euclidean Algorithm
13. Extended Euclidean Algorithm

### Abstract Algebra Operations ✨ NEW
1. **Set Group Generation** - Generate Z_n groups
2. **Coset Groups** (additive & multiplicative)
3. **Quotient Groups** (additive & multiplicative)
4. **Cyclic Groups** (additive & multiplicative)
5. **Generator Finding** - Find generators of multiplicative groups
6. **Polynomial Operations** - Evaluate, add, and multiply polynomials

### Cryptographic Algorithm Simulations 🔐 NEW

Kriptonite now includes interactive demonstrations of classic cryptographic algorithms:

1. **RSA Encryption/Decryption**
   - Complete key generation workflow
   - Encryption and decryption demonstration
   - Step-by-step mathematical verification

2. **Diffie-Hellman Key Exchange**
   - Shared secret establishment
   - Public/private key operations
   - Security demonstration

3. **Elliptic Curve Cryptography (ECC) Concepts**
   - Point verification on elliptic curves
   - Finite field operations
   - Group structure exploration

4. **Digital Signature Algorithm (DSA)**
   - Parameter selection
   - Key generation
   - Signing process walkthrough

5. **Discrete Logarithm Problem**
   - Demonstrates cryptographic hardness
   - Educational brute-force examples

6. **Chinese Remainder Theorem (CRT)**
   - RSA optimization techniques
   - Faster decryption methods

## Quick Start

### Installation
```bash
git clone https://github.com/KanakSasak/kriptonite.git
cd kriptonite
go build -o kriptonite
```

### Basic Usage

**Number Theory:**
```bash
# Generate primes
./kriptonite prime 1 100

# Check if number is prime
./kriptonite isprime 17

# Modular exponentiation
./kriptonite modpow 9 7 143
```

**Abstract Algebra:**
```bash
# Generate cyclic group
./kriptonite cyclicgroup multiplicative 3 7

# Find generators
./kriptonite generators 11

# Polynomial operations
./kriptonite polynomial evaluate 1,2,3 5
```

**Cryptographic Simulations:**
```bash
# Run all crypto simulations (interactive demo)
./examples_crypto.sh

# Or test individual components
./kriptonite modpow 9 7 143    # RSA encryption
./kriptonite modpow 48 103 143 # RSA decryption
```

## Documentation

- **[TEST_SCENARIOS.md](TEST_SCENARIOS.md)** - Comprehensive test scenarios for all features
  - Detailed usage examples
  - Expected outputs
  - Cryptographic algorithm walkthroughs

- **[examples_crypto.sh](examples_crypto.sh)** - Interactive cryptographic simulation script
  - Automated demonstrations of RSA, DH, ECC, DSA
  - Color-coded output
  - Educational explanations

## Testing

Run unit tests:
```bash
go test ./internal/core/services/... -v
```

Run integration tests:
```bash
./test_integration.sh
```

## Architecture

Kriptonite follows a clean hexagonal architecture:
- **Domain Layer** - Core mathematical models
- **Ports** - Interface definitions
- **Services** - Business logic implementation
- **Controllers** - Operation orchestration
- **Handlers** - CLI request handling
- **Commands** - Cobra CLI commands

## Educational Use Cases

### For Students
- Understand RSA encryption from first principles
- Explore group theory concepts hands-on
- Visualize discrete logarithm problem
- Learn modular arithmetic practically

### For Educators
- Demonstrate cryptographic concepts in class
- Provide homework verification tools
- Create custom examples with real calculations
- Show security properties through mathematics

## Examples

### Complete RSA Example
```bash
# 1. Generate primes
./kriptonite prime 10 30
# Choose p=11, q=13

# 2. Calculate φ(n)
./kriptonite euler 143
# φ(143) = 120

# 3. Verify e is coprime
./kriptonite relprime 7 120
# true

# 4. Find private key
./kriptonite exeuclid 7 120
# d = 103

# 5. Encrypt
./kriptonite modpow 9 7 143
# c = 48

# 6. Decrypt
./kriptonite modpow 48 103 143
# m = 9 ✓
```

### Diffie-Hellman Key Exchange
```bash
# Alice's public key
./kriptonite modpow 5 6 23
# A = 8

# Bob's public key
./kriptonite modpow 5 15 23
# B = 19

# Shared secret (computed by both)
./kriptonite modpow 19 6 23  # Alice
./kriptonite modpow 8 15 23  # Bob
# Both get s = 2 ✓
```

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues for bugs and feature requests.

## License

This project is open source and available for educational purposes.

---

**Note:** This tool is designed for educational purposes to help students understand cryptographic concepts. For production cryptography, always use well-tested libraries.