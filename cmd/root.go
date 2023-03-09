/*
Copyright © 2023 yuanjun <imoowi@qq.com>

*/
package cmd

import (
	"os"

	"github.com/imoowi/goRESTApiGen/app"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "goRESTApiGen",
	Short:   "goRESTApiGen -a",
	Example: "goRESTApiGen -a [appname]",
	Long:    ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var generator = app.Generator{}
		generator.Gen(cmd, args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goRESTApiGen.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("appname", "a", "", "api path name, such as :'/api/[appname]'")
	rootCmd.Flags().StringP("path", "p", "", "app folder path")
	rootCmd.Flags().StringP("service", "s", "", "service name")
	rootCmd.Flags().StringP("model", "m", "", "model name")
}
