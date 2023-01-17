/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-app",
	Short: "Go App",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-app.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(versionCmd)

	rootCmd.AddCommand(removeSpecials)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Go App",
	Long:  `All software has versions. This is Go Apps's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.2")
	},
}

var removeSpecials = &cobra.Command{
	Use:   "removespecials",
	Short: "Remove Special characters from the first argument",
	Long: `This removes below special characters from the first argument.
			"*", "/","-".`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			input := args[0]
			output := strings.ReplaceAll(input, "*", "")
			output = strings.ReplaceAll(output, "/", "")
			output = strings.ReplaceAll(output, "-", "")
			output = strings.TrimSpace(output)
			fmt.Println(output)
		} else {
			fmt.Println("No Input")
		}
	},
}
