#!/bin/bash

# Kriptonite Integration Test Script
# Tests all CLI commands with various scenarios

set -e  # Exit on any error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

PASSED=0
FAILED=0

# Build the project first
echo -e "${YELLOW}Building Kriptonite...${NC}"
go build -o kriptonite

echo -e "\n${YELLOW}Starting Integration Tests...${NC}\n"

# Helper function to run a test
run_test() {
    local test_name=$1
    local command=$2
    local expected_in_output=$3

    echo -e "Testing: ${test_name}"

    if output=$(eval "./kriptonite $command 2>&1"); then
        if [[ -n "$expected_in_output" ]] && echo "$output" | grep -q "$expected_in_output"; then
            echo -e "${GREEN}✓ PASSED${NC}"
            ((PASSED++))
        elif [[ -z "$expected_in_output" ]]; then
            echo -e "${GREEN}✓ PASSED${NC}"
            ((PASSED++))
        else
            echo -e "${RED}✗ FAILED${NC} - Expected '$expected_in_output' not found in output"
            echo "Output: $output"
            ((FAILED++))
        fi
    else
        echo -e "${RED}✗ FAILED${NC} - Command failed with exit code $?"
        echo "Output: $output"
        ((FAILED++))
    fi
    echo ""
}

# Test Basic Number Theory Operations
echo -e "${YELLOW}=== Basic Number Theory Tests ===${NC}\n"

run_test "Prime Generation (1-20)" "prime 1 20" "2"
run_test "Check Prime (17)" "isprime 17" ""
run_test "Check Non-Prime (24)" "isprime 24" ""
run_test "Relative Prime (8, 15)" "relprime 8 15" ""
run_test "Factorial (5)" "factorial 5" "120"
run_test "Factor (12)" "factor 12" "1"
run_test "Modulus (17, 5)" "mod 17 5" "2"
run_test "Power (2, 8)" "pow 2 8" "256"
run_test "Modular Power (2, 10, 13)" "modpow 2 10 13" "10"
run_test "Euler Function (12)" "euler 12" "4"
run_test "Euclidean (48, 18)" "euclid 48 18" "6"
run_test "Extended Euclidean (35, 15)" "exeuclid 35 15" "5"

# Test Abstract Algebra Operations
echo -e "${YELLOW}=== Abstract Algebra Tests ===${NC}\n"

# Set Group Tests
run_test "Set Group Z_5" "setgroup 5" "[0 1 2 3 4]"
run_test "Set Group Z_10" "setgroup 10" "[0 1 2 3 4 5 6 7 8 9]"

# Coset Tests
run_test "Additive Coset (H+1 mod 9)" "coset additive 0,3,6 1 9" "[1 4 7]"
run_test "Additive Coset (H+2 mod 9)" "coset additive 0,3,6 2 9" "[2 5 8]"
run_test "Multiplicative Coset (H*3 mod 7)" "coset multiplicative 1,2,4 3 7" "[3 6 5]"

# Quotient Group Tests
run_test "Additive Quotient Group (Z_12/H)" "quotientgroup additive 12 3" "Coset"
run_test "Additive Quotient Group (Z_6/H)" "quotientgroup additive 6 2" "Coset"

# Cyclic Group Tests
run_test "Additive Cyclic Group <3> mod 12" "cyclicgroup additive 3 12" "[0 3 6 9]"
run_test "Additive Cyclic Group <1> mod 5" "cyclicgroup additive 1 5" "[0 1 2 3 4]"
run_test "Multiplicative Cyclic Group <3> mod 7" "cyclicgroup multiplicative 3 7" "[1 3 2 6 4 5]"
run_test "Multiplicative Cyclic Group <2> mod 7" "cyclicgroup multiplicative 2 7" "[1 2 4]"

# Generator Tests
run_test "Generators of Z_7*" "generators 7" "[3 5]"
run_test "Generators of Z_5*" "generators 5" "[2 3]"

# Polynomial Tests
run_test "Polynomial Evaluate (1+2x+3x^2 at x=5)" "polynomial evaluate 1,2,3 5" "86"
run_test "Polynomial Evaluate (5 at x=10)" "polynomial evaluate 5 10" "5"
run_test "Polynomial Add (1+2x)+(3+4x)" "polynomial add 1,2 3,4" "[4 6]"
run_test "Polynomial Add (1+2x+3x^2)+(4+5x+6x^2)" "polynomial add 1,2,3 4,5,6" "[5 7 9]"
run_test "Polynomial Multiply (1+2x)×(3+4x)" "polynomial multiply 1,2 3,4" "[3 10 8]"
run_test "Polynomial Multiply x×x" "polynomial multiply 0,1 0,1" "[0 0 1]"

# Edge Cases
echo -e "${YELLOW}=== Edge Case Tests ===${NC}\n"

run_test "Empty Polynomial Evaluation" "polynomial evaluate '' 5" "0"
run_test "Set Group Z_1" "setgroup 1" "[0]"
run_test "Cyclic Group Order 1" "cyclicgroup additive 5 5" "[0]"

# Summary
echo -e "${YELLOW}=== Test Summary ===${NC}"
TOTAL=$((PASSED + FAILED))
echo -e "Total Tests: $TOTAL"
echo -e "${GREEN}Passed: $PASSED${NC}"
echo -e "${RED}Failed: $FAILED${NC}"

if [ $FAILED -eq 0 ]; then
    echo -e "\n${GREEN}All tests passed! ✓${NC}"
    exit 0
else
    echo -e "\n${RED}Some tests failed! ✗${NC}"
    exit 1
fi
