/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Ho-J/base/command/app/goproxy"
	"github.com/spf13/cobra"
)

// goproxyCmd represents the goproxy command
var goproxyCmd = &cobra.Command{
	Use:   "goproxy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		goproxy.Run()
	},
}

func init() {
	rootCmd.AddCommand(goproxyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// goproxyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// goproxyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}