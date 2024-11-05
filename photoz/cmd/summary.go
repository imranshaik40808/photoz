/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"photoz/helper"
	option "photoz/option"

	"github.com/spf13/cobra"
)

// summaryCmd represents the info command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Get summary of the directory",
	Long: `Get the summary of the directory. For example:

./photoz summary -p /Users/imranshaik/Documents -e="/Users/imranshaik/Documents/Projects/photoz/model/file" -e="/Users/imranshaik/Documents/Projects/photoz/option"   
`,
	Run: func(cmd *cobra.Command, args []string) {
		path := summaryOption.Path
		if len(path) > 0 {
			path = helper.ChangePathSeparator(path)
			summary := helper.GetSummary(path, summaryOption.ExcludePath, summaryOption.ExcludeExtension)
			if summary != nil {
				v, err := json.Marshal(summary)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(v))
			} else {
				println("No data found")
			}
		}
	},
}

var summaryOption option.SummaryOption

func init() {
	rootCmd.AddCommand(summaryCmd)
	f := summaryCmd.Flags()
	f.StringVarP(&summaryOption.Path, "path", "p", "", "specifiy the folder path")
	f.StringArrayVarP(&summaryOption.ExcludePath, "exclude-path", "e", nil, "specifiy the folder path to exclude")
	f.StringArrayVarP(&summaryOption.ExcludeExtension, "exclude-ext", "x", nil, "specifiy the extension to exclude")
}
