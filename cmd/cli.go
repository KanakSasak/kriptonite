package cmd

import (
	"Kryptonite/internal/adapter"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var top = ".xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx\n\n"
var banner = "██╗  ██╗██████╗ ██╗██████╗ ████████╗ ██████╗ ███╗   ██╗██╗████████╗███████╗\n██║ ██╔╝██╔══██╗██║██╔══██╗╚══██╔══╝██╔═══██╗████╗  ██║██║╚══██╔══╝██╔════╝\n█████╔╝ ██████╔╝██║██████╔╝   ██║   ██║   ██║██╔██╗ ██║██║   ██║   █████╗  \n██╔═██╗ ██╔══██╗██║██╔═══╝    ██║   ██║   ██║██║╚██╗██║██║   ██║   ██╔══╝  \n██║  ██╗██║  ██║██║██║        ██║   ╚██████╔╝██║ ╚████║██║   ██║   ███████╗\n╚═╝  ╚═╝╚═╝  ╚═╝╚═╝╚═╝        ╚═╝    ╚═════╝ ╚═╝  ╚═══╝╚═╝   ╚═╝   ╚══════╝\n                                                                           \n"
var foot = ".xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx.xOx\n"
var by = "Created By : KanakSasak\n"
var repo = "Repo : https://github.com/KanakSasak/kriptonite\n"
var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "kriptonite",
		Short: top + banner + foot + by + repo,
	}
	ctx *adapter.HandlerContext
)

// Executes the command line interface
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is ./config/config.yaml)")
}

func initContext() {
	ctx = adapter.NewContext()
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./config")
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")

	viper.SetEnvPrefix("cli")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Failed to load config file:", viper.ConfigFileUsed())
	}

	initContext()
}
