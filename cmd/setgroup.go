package cmd

import (
	"Kryptonite/internal/adapter/handler"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	setGroupCmd = &cobra.Command{
		Use:   "setgroup",
		Short: "Generate Set Group Z_n",
		Long:  "Generate the set group Z_n = {0, 1, 2, ..., n-1}",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("usage: kriptonite setgroup <n>")
			}

			cliHandler := handler.NewCliServer(ctx)

			n, err := strconv.Atoi(args[0])
			if err != nil {
				return err
			}

			data := cliHandler.GenerateSetGroup(n)

			fmt.Printf("Set Group Z_%d = %v\n", n, data)
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(setGroupCmd)
}
