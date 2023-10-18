package cmd

import (
	"Kryptonite/internal/adapter/handler"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	modCmd = &cobra.Command{
		Use:   "mod",
		Short: "Calculate Mod",
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

			data := cliHandler.Mod(float64(args0), float64(args1))

			fmt.Println(data)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(modCmd)
}
