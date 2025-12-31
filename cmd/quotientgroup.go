package cmd

import (
	"Kryptonite/internal/adapter/handler"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

var (
	quotientGroupCmd = &cobra.Command{
		Use:   "quotientgroup",
		Short: "Generate Quotient Groups",
		Long:  "Generate additive or multiplicative quotient groups",
	}

	additiveQuotientCmd = &cobra.Command{
		Use:   "additive",
		Short: "Generate Additive Quotient Group",
		Long:  "Generate additive quotient group Z_n / H",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("usage: kriptonite quotientgroup additive <n> <subgroup_size>\nexample: kriptonite quotientgroup additive 12 3")
			}

			cliHandler := handler.NewCliServer(ctx)

			n, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			subgroupSize, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			data := cliHandler.GenerateAdditiveQuotientGroup(n, subgroupSize)

			fmt.Printf("Additive Quotient Group Z_%d / H (where |H| = %d):\n", n, subgroupSize)
			for i, coset := range data {
				fmt.Printf("  Coset %d: %v\n", i, coset)
			}
			return nil
		},
	}

	multiplicativeQuotientCmd = &cobra.Command{
		Use:   "multiplicative",
		Short: "Generate Multiplicative Quotient Group",
		Long:  "Generate multiplicative quotient group Z_n* / H",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("usage: kriptonite quotientgroup multiplicative <n> <subgroup>\nexample: kriptonite quotientgroup multiplicative 7 1,2,4")
			}

			cliHandler := handler.NewCliServer(ctx)

			n, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			// Parse subgroup
			subgroupStr := strings.Split(args[1], ",")
			subgroup := make([]int, len(subgroupStr))
			for i, s := range subgroupStr {
				val, err := strconv.Atoi(s)
				if err != nil {
					return err
				}
				subgroup[i] = val
			}

			data := cliHandler.GenerateMultiplicativeQuotientGroup(n, subgroup)

			fmt.Printf("Multiplicative Quotient Group Z_%d* / H (where H = %v):\n", n, subgroup)
			for i, coset := range data {
				fmt.Printf("  Coset %d: %v\n", i, coset)
			}
			return nil
		},
	}
)

func init() {
	quotientGroupCmd.AddCommand(additiveQuotientCmd)
	quotientGroupCmd.AddCommand(multiplicativeQuotientCmd)
	rootCmd.AddCommand(quotientGroupCmd)
}
