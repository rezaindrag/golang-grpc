package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "catalina",
		Short: "Catalina is a catalog management service.",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.json", "set config file")

	rootCmd.AddCommand(grpcCmd)
}

func initConfig() {
	viper.SetConfigFile(cfgFile)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
