package cmd

import (
	"Kryptonite/internal/adapter/handler"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	cyclicGroupCmd = &cobra.Command{
		Use:   "cyclicgroup",
		Short: "Generate Cyclic Groups",
		Long:  "Generate additive or multiplicative cyclic groups",
	}

	additiveCyclicCmd = &cobra.Command{
		Use:   "additive",
		Short: "Generate Additive Cyclic Group",
		Long:  "Generate additive cyclic group <g> = {0, g, 2g, 3g, ...} (mod n)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("usage: kriptonite cyclicgroup additive <generator> <modulus>\nexample: kriptonite cyclicgroup additive 3 12")
			}

			cliHandler := handler.NewCliServer(ctx)

			generator, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			modulus, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			data := cliHandler.GenerateAdditiveCyclicGroup(generator, modulus)

			fmt.Printf("Additive Cyclic Group <%d> (mod %d) = %v\n", generator, modulus, data)
			fmt.Printf("Order: %d\n", len(data))
			return nil
		},
	}

	multiplicativeCyclicCmd = &cobra.Command{
		Use:   "multiplicative",
		Short: "Generate Multiplicative Cyclic Group",
		Long:  "Generate multiplicative cyclic group <g> = {1, g, g^2, g^3, ...} (mod n)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("usage: kriptonite cyclicgroup multiplicative <generator> <modulus>\nexample: kriptonite cyclicgroup multiplicative 3 7")
			}

			cliHandler := handler.NewCliServer(ctx)

			generator, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			modulus, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			data := cliHandler.GenerateMultiplicativeCyclicGroup(generator, modulus)

			fmt.Printf("Multiplicative Cyclic Group <%d> (mod %d) = %v\n", generator, modulus, data)
			fmt.Printf("Order: %d\n", len(data))
			return nil
		},
	}
)

func init() {
	cyclicGroupCmd.AddCommand(additiveCyclicCmd)
	cyclicGroupCmd.AddCommand(multiplicativeCyclicCmd)
	rootCmd.AddCommand(cyclicGroupCmd)
}
