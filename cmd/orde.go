package cmd

import (
	"Kryptonite/internal/adapter/handler"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	ordeCmd = &cobra.Command{
		Use:   "orde",
		Short: "Generate Orde",
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

			data := cliHandler.Orde(args0, args1)

			fmt.Println(data)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(ordeCmd)
}
