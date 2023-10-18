package cmd

import (
	"Kryptonite/internal/adapter/handler"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	euclidCmd = &cobra.Command{
		Use:   "euclid",
		Short: "Calculate Euclidean Algoritm",
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

			data := cliHandler.Euclidean(args0, args1)

			fmt.Println(data)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(euclidCmd)
}
