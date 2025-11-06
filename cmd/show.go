/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"pb/internal/parser"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Shows you the playbook",
	Long: `Shows you the playbook`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("show called")
		name, _ := cmd.Flags().GetString("name")
		basePath, _ := parser.GetPlaybookBasePath() 
		fp := basePath + "/" + name

		playbook, err := parser.GetPlaybook(fp)

		if (err != nil) {
			fmt.Println("ERRROR")
			fmt.Println(err)
		}

		fmt.Println(playbook)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	showCmd.Flags().StringP("name", "n", "", "Playbook name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
