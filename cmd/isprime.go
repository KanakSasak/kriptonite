package cmd

import (
	"Kryptonite/internal/adapter/handler"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	isPrimeCmd = &cobra.Command{
		Use:   "isprime",
		Short: "Check Prime Number",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliHandler := handler.NewCliServer(ctx)

			args0, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			data := cliHandler.IsPrime(args0)
			fmt.Println(data)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(isPrimeCmd)
}
