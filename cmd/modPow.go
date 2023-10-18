package cmd

import (
	"Kryptonite/internal/adapter/handler"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	modPowCmd = &cobra.Command{
		Use:   "modpow",
		Short: "Calculate Power Modulus",
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
			args2, err := strconv.Atoi(args[2])
			if err != nil {
				return err
			}

			data := cliHandler.ModPow(float64(args0), float64(args1), float64(args2))

			fmt.Println(data)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(modPowCmd)
}
