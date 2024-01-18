/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "boop",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1)),
	Short: "An easy way to create files from the terminal in Windows.",
	Long: `An easy way to create files from the terminal in Windows. 
	Similar to the 'touch' utility on Linux.`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, v := range args {
			_, err := os.Create(v)
			if err != nil {
				panic("Could not create the file.")
			}
		}
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.boop.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
