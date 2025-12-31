package cmd

import (
	"Kryptonite/internal/adapter/handler"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var (
	polynomialCmd = &cobra.Command{
		Use:   "polynomial",
		Short: "Polynomial Operations",
		Long:  "Perform polynomial operations (evaluate, add, multiply)",
	}

	evaluatePolyCmd = &cobra.Command{
		Use:   "evaluate",
		Short: "Evaluate Polynomial",
		Long:  "Evaluate polynomial at x: a_0 + a_1*x + a_2*x^2 + ...",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("usage: kriptonite polynomial evaluate <coefficients> <x>\nexample: kriptonite polynomial evaluate 1,2,3 5")
			}

			cliHandler := handler.NewCliServer(ctx)

			// Parse coefficients
			coeffStr := strings.Split(args[0], ",")
			coefficients := make([]int, len(coeffStr))
			for i, s := range coeffStr {
				val, err := strconv.Atoi(s)
				if err != nil {
					return err
				}
				coefficients[i] = val
			}

			x, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			result := cliHandler.EvaluatePolynomial(coefficients, x)

			// Format polynomial string
			polyStr := formatPolynomial(coefficients)
			fmt.Printf("Polynomial: %s\n", polyStr)
			fmt.Printf("P(%d) = %d\n", x, result)
			return nil
		},
	}

	addPolyCmd = &cobra.Command{
		Use:   "add",
		Short: "Add Polynomials",
		Long:  "Add two polynomials",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("usage: kriptonite polynomial add <poly1> <poly2>\nexample: kriptonite polynomial add 1,2,3 4,5")
			}

			cliHandler := handler.NewCliServer(ctx)

			// Parse first polynomial
			poly1Str := strings.Split(args[0], ",")
			poly1 := make([]int, len(poly1Str))
			for i, s := range poly1Str {
				val, err := strconv.Atoi(s)
				if err != nil {
					return err
				}
				poly1[i] = val
			}

			// Parse second polynomial
			poly2Str := strings.Split(args[1], ",")
			poly2 := make([]int, len(poly2Str))
			for i, s := range poly2Str {
				val, err := strconv.Atoi(s)
				if err != nil {
					return err
				}
				poly2[i] = val
			}

			result := cliHandler.AddPolynomials(poly1, poly2)

			fmt.Printf("P1(x) = %s\n", formatPolynomial(poly1))
			fmt.Printf("P2(x) = %s\n", formatPolynomial(poly2))
			fmt.Printf("P1(x) + P2(x) = %s\n", formatPolynomial(result))
			fmt.Printf("Result coefficients: %v\n", result)
			return nil
		},
	}

	multiplyPolyCmd = &cobra.Command{
		Use:   "multiply",
		Short: "Multiply Polynomials",
		Long:  "Multiply two polynomials",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("usage: kriptonite polynomial multiply <poly1> <poly2>\nexample: kriptonite polynomial multiply 1,2 3,4")
			}

			cliHandler := handler.NewCliServer(ctx)

			// Parse first polynomial
			poly1Str := strings.Split(args[0], ",")
			poly1 := make([]int, len(poly1Str))
			for i, s := range poly1Str {
				val, err := strconv.Atoi(s)
				if err != nil {
					return err
				}
				poly1[i] = val
			}

			// Parse second polynomial
			poly2Str := strings.Split(args[1], ",")
			poly2 := make([]int, len(poly2Str))
			for i, s := range poly2Str {
				val, err := strconv.Atoi(s)
				if err != nil {
					return err
				}
				poly2[i] = val
			}

			result := cliHandler.MultiplyPolynomials(poly1, poly2)

			fmt.Printf("P1(x) = %s\n", formatPolynomial(poly1))
			fmt.Printf("P2(x) = %s\n", formatPolynomial(poly2))
			fmt.Printf("P1(x) × P2(x) = %s\n", formatPolynomial(result))
			fmt.Printf("Result coefficients: %v\n", result)
			return nil
		},
	}
)

// formatPolynomial formats polynomial coefficients as a readable string
func formatPolynomial(coefficients []int) string {
	if len(coefficients) == 0 {
		return "0"
	}

	var terms []string
	for i, coef := range coefficients {
		if coef == 0 {
			continue
		}

		term := ""
		if i == 0 {
			term = fmt.Sprintf("%d", coef)
		} else if i == 1 {
			if coef == 1 {
				term = "x"
			} else if coef == -1 {
				term = "-x"
			} else {
				term = fmt.Sprintf("%dx", coef)
			}
		} else {
			if coef == 1 {
				term = fmt.Sprintf("x^%d", i)
			} else if coef == -1 {
				term = fmt.Sprintf("-x^%d", i)
			} else {
				term = fmt.Sprintf("%dx^%d", coef, i)
			}
		}

		if len(terms) > 0 && coef > 0 {
			terms = append(terms, "+"+term)
		} else {
			terms = append(terms, term)
		}
	}

	if len(terms) == 0 {
		return "0"
	}

	result := ""
	for _, term := range terms {
		result += term
	}
	return result
}

func init() {
	polynomialCmd.AddCommand(evaluatePolyCmd)
	polynomialCmd.AddCommand(addPolyCmd)
	polynomialCmd.AddCommand(multiplyPolyCmd)
	rootCmd.AddCommand(polynomialCmd)
}
