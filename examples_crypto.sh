#!/bin/bash

# Kriptonite Cryptographic Algorithm Simulations
# Demonstrates RSA, ECC concepts, DSA, and Diffie-Hellman

# Colors for output
BLUE='\033[0;34m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

echo -e "${CYAN}╔════════════════════════════════════════════════════════╗${NC}"
echo -e "${CYAN}║   KRIPTONITE - Cryptographic Algorithm Simulations   ║${NC}"
echo -e "${CYAN}╔════════════════════════════════════════════════════════╗${NC}\n"

# ============================================================================
# 1. RSA ENCRYPTION/DECRYPTION SIMULATION
# ============================================================================

echo -e "${BLUE}═══════════════════════════════════════════════════${NC}"
echo -e "${BLUE}  1. RSA Encryption/Decryption Simulation${NC}"
echo -e "${BLUE}═══════════════════════════════════════════════════${NC}\n"

echo -e "${YELLOW}Step 1: Generate Prime Numbers${NC}"
echo "Finding primes between 10 and 30..."
./kriptonite prime 10 30
echo -e "\n${GREEN}Choosing p=11, q=13${NC}\n"

echo -e "${YELLOW}Step 2: Calculate n = p × q${NC}"
echo "n = 11 × 13 = 143"
n=143
echo -e "${GREEN}n = $n (modulus)${NC}\n"

echo -e "${YELLOW}Step 3: Calculate φ(n) = (p-1)(q-1)${NC}"
echo "φ(143) = 10 × 12 = 120"
echo "Verification using Euler function:"
./kriptonite euler 143
echo ""

echo -e "${YELLOW}Step 4: Choose Public Exponent e${NC}"
echo "Checking if e=7 is coprime to φ(n)=120..."
./kriptonite relprime 7 120
echo -e "${GREEN}e = 7 (public exponent)${NC}\n"

echo -e "${YELLOW}Step 5: Calculate Private Exponent d${NC}"
echo "Finding d such that 7d ≡ 1 (mod 120)..."
./kriptonite exeuclid 7 120
echo ""
echo -e "${GREEN}d = 103 (private exponent)${NC}"
echo "Verification: (7 × 103) mod 120 ="
./kriptonite mod 721 120
echo ""

echo -e "${YELLOW}Step 6: Encrypt Message${NC}"
echo "Message: m = 9"
echo "Encryption: c = 9^7 mod 143"
echo -n "Ciphertext c = "
./kriptonite modpow 9 7 143
echo ""

echo -e "${YELLOW}Step 7: Decrypt Ciphertext${NC}"
echo "Ciphertext: c = 48"
echo "Decryption: m = 48^103 mod 143"
echo -n "Recovered message m = "
./kriptonite modpow 48 103 143
echo ""

echo -e "${GREEN}✓ RSA Encryption/Decryption Successful!${NC}"
echo -e "${GREEN}  Original: 9 → Encrypted: 48 → Decrypted: 9${NC}\n"

sleep 2

# ============================================================================
# 2. DIFFIE-HELLMAN KEY EXCHANGE
# ============================================================================

echo -e "${BLUE}═══════════════════════════════════════════════════${NC}"
echo -e "${BLUE}  2. Diffie-Hellman Key Exchange${NC}"
echo -e "${BLUE}═══════════════════════════════════════════════════${NC}\n"

echo -e "${YELLOW}Setup: Public Parameters${NC}"
echo "Prime p = 23"
echo "Finding generators of Z_23*..."
./kriptonite generators 23
echo -e "${GREEN}Choosing g = 5 (generator)${NC}\n"

echo -e "${YELLOW}Alice's Side:${NC}"
echo "Alice chooses private key a = 6"
echo "Alice computes public key A = 5^6 mod 23"
echo -n "Alice's public key A = "
./kriptonite modpow 5 6 23
echo ""

echo -e "${YELLOW}Bob's Side:${NC}"
echo "Bob chooses private key b = 15"
echo "Bob computes public key B = 5^15 mod 23"
echo -n "Bob's public key B = "
./kriptonite modpow 5 15 23
echo ""

echo -e "${YELLOW}Shared Secret Computation:${NC}"
echo "Alice computes: s = B^a mod p = 19^6 mod 23"
echo -n "Alice's result: s = "
ALICE_SECRET=$(./kriptonite modpow 19 6 23)
echo "$ALICE_SECRET"

echo "Bob computes: s = A^b mod p = 8^15 mod 23"
echo -n "Bob's result: s = "
BOB_SECRET=$(./kriptonite modpow 8 15 23)
echo "$BOB_SECRET"
echo ""

if [ "$ALICE_SECRET" = "$BOB_SECRET" ]; then
    echo -e "${GREEN}✓ Shared secrets match! s = $ALICE_SECRET${NC}"
    echo -e "${GREEN}  Alice and Bob can now use this for symmetric encryption!${NC}\n"
else
    echo -e "✗ Shared secrets don't match (should not happen)\n"
fi

sleep 2

# ============================================================================
# 3. ELLIPTIC CURVE CRYPTOGRAPHY (ECC) CONCEPTS
# ============================================================================

echo -e "${BLUE}═══════════════════════════════════════════════════${NC}"
echo -e "${BLUE}  3. Elliptic Curve Cryptography (ECC) Concepts${NC}"
echo -e "${BLUE}═══════════════════════════════════════════════════${NC}\n"

echo -e "${YELLOW}Curve: y² = x³ + x + 1 (mod 23)${NC}\n"

echo -e "${YELLOW}Verifying Point (3, 10) on Curve:${NC}"
echo "Left side: y² mod 23 = 10² mod 23"
echo -n "Result: "
LEFT=$(./kriptonite modpow 10 2 23)
echo "$LEFT"

echo "Right side: (x³ + x + 1) mod 23"
echo "x³ mod 23 = 3³ mod 23"
echo -n "Result: "
X_CUBED=$(./kriptonite modpow 3 3 23)
echo "$X_CUBED"
echo "Adding: $X_CUBED + 3 + 1 = 8"
echo ""

if [ "$LEFT" = "8" ]; then
    echo -e "${GREEN}✓ Point (3, 10) is on the curve!${NC}"
    echo -e "${GREEN}  8 ≡ 8 (mod 23) ✓${NC}\n"
fi

echo -e "${YELLOW}Cyclic Group Structure (analogous to scalar multiplication):${NC}"
echo "Generating multiplicative cyclic group <2> in Z_23:"
./kriptonite cyclicgroup multiplicative 2 23
echo ""

sleep 2

# ============================================================================
# 4. DIGITAL SIGNATURE ALGORITHM (DSA) CONCEPTS
# ============================================================================

echo -e "${BLUE}═══════════════════════════════════════════════════${NC}"
echo -e "${BLUE}  4. Digital Signature Algorithm (DSA) Concepts${NC}"
echo -e "${BLUE}═══════════════════════════════════════════════════${NC}\n"

echo -e "${YELLOW}Domain Parameters:${NC}"
echo "Prime p = 23"
echo "Prime q = 11 (divides p-1)"
echo ""
echo "Verifying q | (p-1):"
echo "(p-1) mod q = 22 mod 11 ="
./kriptonite mod 22 11
echo -e "${GREEN}✓ q divides (p-1)${NC}\n"

echo -e "${YELLOW}Finding Generator:${NC}"
echo "Computing g = 2^((p-1)/q) mod p = 2^2 mod 23"
echo -n "g = "
./kriptonite modpow 2 2 23
echo ""
echo "Verifying g^q mod p = 4^11 mod 23:"
./kriptonite modpow 4 11 23
echo -e "${GREEN}✓ g is valid generator of subgroup${NC}\n"

echo -e "${YELLOW}Key Generation:${NC}"
echo "Private key: x = 6"
echo "Public key: y = g^x mod p = 4^6 mod 23"
echo -n "Public key y = "
./kriptonite modpow 4 6 23
echo -e "${GREEN}✓ Key pair generated${NC}\n"

sleep 2

# ============================================================================
# 5. DISCRETE LOGARITHM PROBLEM
# ============================================================================

echo -e "${BLUE}═══════════════════════════════════════════════════${NC}"
echo -e "${BLUE}  5. Discrete Logarithm Problem Demonstration${NC}"
echo -e "${BLUE}═══════════════════════════════════════════════════${NC}\n"

echo -e "${YELLOW}Problem: Find x such that 3^x ≡ 12 (mod 17)${NC}\n"

echo "Generating all powers of 3 mod 17:"
./kriptonite cyclicgroup multiplicative 3 17
echo ""

echo "Testing specific values:"
echo -n "3^13 mod 17 = "
./kriptonite modpow 3 13 17
echo ""

echo -e "${GREEN}✓ Answer: x = 13${NC}"
echo -e "${YELLOW}Note: This is easy for small numbers, but extremely hard for large ones!${NC}"
echo -e "${YELLOW}This is the basis of cryptographic security.${NC}\n"

sleep 2

# ============================================================================
# 6. CHINESE REMAINDER THEOREM IN RSA
# ============================================================================

echo -e "${BLUE}═══════════════════════════════════════════════════${NC}"
echo -e "${BLUE}  6. Chinese Remainder Theorem (CRT) in RSA${NC}"
echo -e "${BLUE}═══════════════════════════════════════════════════${NC}\n"

echo -e "${YELLOW}RSA-CRT Optimization for Decryption${NC}"
echo "Using previous RSA example: c=48, d=103, n=143, p=11, q=13"
echo ""

echo "Computing dp = d mod (p-1) = 103 mod 10:"
echo -n "dp = "
./kriptonite mod 103 10
echo ""

echo "Computing dq = d mod (q-1) = 103 mod 12:"
echo -n "dq = "
./kriptonite mod 103 12
echo ""

echo "Computing mp = c^dp mod p = 48^3 mod 11:"
echo -n "mp = "
./kriptonite modpow 48 3 11
echo ""

echo "Computing mq = c^dq mod q = 48^7 mod 13:"
echo -n "mq = "
./kriptonite modpow 48 7 13
echo ""

echo -e "${GREEN}✓ Both give 9 (the original message)!${NC}"
echo -e "${YELLOW}CRT allows faster RSA decryption by working with smaller numbers.${NC}\n"

# ============================================================================
# SUMMARY
# ============================================================================

echo -e "${CYAN}╔════════════════════════════════════════════════════════╗${NC}"
echo -e "${CYAN}║                      SUMMARY                          ║${NC}"
echo -e "${CYAN}╚════════════════════════════════════════════════════════╝${NC}\n"

echo -e "${GREEN}✓ RSA Encryption/Decryption${NC} - Demonstrated full cycle"
echo -e "${GREEN}✓ Diffie-Hellman Key Exchange${NC} - Shared secret established"
echo -e "${GREEN}✓ ECC Concepts${NC} - Point verification and group structure"
echo -e "${GREEN}✓ DSA Concepts${NC} - Key generation and parameters"
echo -e "${GREEN}✓ Discrete Logarithm${NC} - Problem difficulty demonstrated"
echo -e "${GREEN}✓ CRT in RSA${NC} - Optimization technique shown"

echo -e "\n${CYAN}All cryptographic simulations completed successfully!${NC}"
echo -e "${CYAN}Kriptonite provides the mathematical foundation for understanding cryptography.${NC}\n"
