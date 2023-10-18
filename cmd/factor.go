package cmd

import (
	"Kryptonite/internal/adapter/handler"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	factorCmd = &cobra.Command{
		Use:   "factor",
		Short: "Calculate Factor",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliHandler := handler.NewCliServer(ctx)

			args0, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			data := cliHandler.CalculateFactors(args0)
			fmt.Println(data)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(factorCmd)
}
