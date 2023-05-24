package cmd

import (
	"github.com/luis-souza-dev/go-playground/dev_env_setup/pkg"
	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use: "stupen",
	Short: "initial setup for dev environment",
	Long: "This will install some tools that I've personally make use of during my day-to-day activities",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute());
	pkg.GetCode();
}