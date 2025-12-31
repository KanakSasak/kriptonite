package cmd

import (
	"Kryptonite/internal/adapter/handler"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	generatorsCmd = &cobra.Command{
		Use:   "generators",
		Short: "Find Generators of Group",
		Long:  "Find all generators of the multiplicative group Z_n*",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("usage: kriptonite generators <n>\nexample: kriptonite generators 7")
			}

			cliHandler := handler.NewCliServer(ctx)

			n, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			data := cliHandler.FindGenerators(n)

			fmt.Printf("Generators of Z_%d*: %v\n", n, data)
			fmt.Printf("Number of generators: %d\n", len(data))
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(generatorsCmd)
}
