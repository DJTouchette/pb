/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"pb/internal/parser"
	"os"
	"path/filepath"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your playbooks",
	Long: `List the playbooks in this git repo or directory`,
	Run: func(cmd *cobra.Command, args []string) {
		pwd, _ := os.Getwd()
		fps, err := parser.Discover(pwd + "/playbooks")

		if err != nil {
			fmt.Println("Something went wront")
			return
		}

		for _, file := range fps {
			fmt.Println(filepath.Base(file))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
