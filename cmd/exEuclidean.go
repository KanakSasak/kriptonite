package cmd

import (
	"Kryptonite/internal/adapter/handler"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	extEuclidCmd = &cobra.Command{
		Use:   "exeuclid",
		Short: "Calculate Extended Euclidean Algoritm",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliHandler := handler.NewCliServer(ctx)

			args0, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}
			args1, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}

			data1, _, data3 := cliHandler.ExtEuclidean(args0, args1)

			fmt.Printf("\nExtended GCD of %d and %d is %d - Inverse is %d\n", args0, args1, data1, data3)

			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(extEuclidCmd)
}
