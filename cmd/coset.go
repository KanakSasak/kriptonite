package cmd

import (
	"Kryptonite/internal/adapter/handler"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var (
	cosetCmd = &cobra.Command{
		Use:   "coset",
		Short: "Generate Coset Groups",
		Long:  "Generate additive or multiplicative coset groups",
	}

	additiveCosetCmd = &cobra.Command{
		Use:   "additive",
		Short: "Generate Additive Coset",
		Long:  "Generate additive coset H + a = {h + a | h ∈ H} (mod n)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 3 {
				return fmt.Errorf("usage: kriptonite coset additive <subgroup> <element> <modulus>\nexample: kriptonite coset additive 0,3,6 1 9")
			}

			cliHandler := handler.NewCliServer(ctx)

			// Parse subgroup
			subgroupStr := strings.Split(args[0], ",")
			subgroup := make([]int, len(subgroupStr))
			for i, s := range subgroupStr {
				val, err := strconv.Atoi(s)
				if err != nil {
					return err
				}
				subgroup[i] = val
			}

			element, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			modulus, err := strconv.Atoi(args[2])
			if err != nil {
				return err
			}

			data := cliHandler.GenerateAdditiveCoset(subgroup, element, modulus)

			fmt.Printf("Subgroup H = %v\n", subgroup)
			fmt.Printf("Additive Coset H + %d (mod %d) = %v\n", element, modulus, data)
			return nil
		},
	}

	multiplicativeCosetCmd = &cobra.Command{
		Use:   "multiplicative",
		Short: "Generate Multiplicative Coset",
		Long:  "Generate multiplicative coset H * a = {h * a | h ∈ H} (mod n)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 3 {
				return fmt.Errorf("usage: kriptonite coset multiplicative <subgroup> <element> <modulus>\nexample: kriptonite coset multiplicative 1,2,4 3 7")
			}

			cliHandler := handler.NewCliServer(ctx)

			// Parse subgroup
			subgroupStr := strings.Split(args[0], ",")
			subgroup := make([]int, len(subgroupStr))
			for i, s := range subgroupStr {
				val, err := strconv.Atoi(s)
				if err != nil {
					return err
				}
				subgroup[i] = val
			}

			element, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			modulus, err := strconv.Atoi(args[2])
			if err != nil {
				return err
			}

			data := cliHandler.GenerateMultiplicativeCoset(subgroup, element, modulus)

			fmt.Printf("Subgroup H = %v\n", subgroup)
			fmt.Printf("Multiplicative Coset H * %d (mod %d) = %v\n", element, modulus, data)
			return nil
		},
	}
)

func init() {
	cosetCmd.AddCommand(additiveCosetCmd)
	cosetCmd.AddCommand(multiplicativeCosetCmd)
	rootCmd.AddCommand(cosetCmd)
}
