/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"path/filepath"
	"photoz/helper"
	model "photoz/model/file"
	option "photoz/option"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// arrangeCmd represents the arrange command
var arrangeCmd = &cobra.Command{
	Use:   "arrange",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		path := arrangeOption.Path
		if len(path) > 0 {
			path = helper.ChangePathSeparator(path)
			summary := helper.GetSummary(path, arrangeOption.ExcludePath, arrangeOption.ExcludeExtension)

			duplicateFileCount := 0
			duplicateFileSize := 0
			duplicateFiles := make([]string, 0)

			successfulFileCount := 0
			successfulFileSize := 0
			successfulFiles := make([]string, 0)

			failedFileCount := 0
			failedFileSize := 0
			failedFiles := make([]string, 0)
			hashMap := make(map[string]string)
			folderMap := make(map[int]map[time.Month]map[int][]string)
			rootPath := filepath.Join(arrangeOption.DestinationPath)
			helper.CreateRootDirectory(rootPath)
			if summary != nil {
				for _, info := range summary.Files {
					year := info.CreatedDate.Year()
					month := info.CreatedDate.Month()
					day := info.CreatedDate.Day()
					if monthMap, ok := folderMap[year]; ok {
						if dayMap, ok := monthMap[month]; ok {
							if _, ok := dayMap[day]; ok {
								//data = append(data, info.Path)
								//dayMap[day] = data
								if _, hashPresent := hashMap[info.Hash]; hashPresent {
									duplicateFiles = append(duplicateFiles, info.Path)
									duplicateFileCount += 1
									duplicateFileSize += int(info.Size)
								} else {
									hashMap[info.Hash] = ""
									fileName := filepath.Base(info.Path)
									err := helper.MoveFile(info.Path, filepath.Join(arrangeOption.DestinationPath, helper.ROOT, strconv.Itoa(year), strconv.Itoa(int(month)), strconv.Itoa(day), fileName))
									if err != nil {
										failedFileCount += 1
										failedFiles = append(failedFiles, info.Path)
										failedFileSize += int(info.Size)
									} else {
										successfulFileCount += 1
										successfulFiles = append(successfulFiles, info.Path)
										successfulFileSize += int(info.Size)
									}
								}
							} else {
								dayMap[day] = make([]string, 0)
								helper.CreateDirectory(filepath.Join(arrangeOption.DestinationPath, helper.ROOT, strconv.Itoa(year), strconv.Itoa(int(month)), strconv.Itoa(day)))
							}
						} else {
							monthMap[month] = make(map[int][]string)
							helper.CreateDirectory(filepath.Join(arrangeOption.DestinationPath, helper.ROOT, strconv.Itoa(year), strconv.Itoa(int(month))))
						}
					} else {
						folderMapYear := make(map[time.Month]map[int][]string)
						folderMap[year] = folderMapYear
						helper.CreateDirectory(filepath.Join(arrangeOption.DestinationPath, helper.ROOT, strconv.Itoa(year)))
					}
				}
				arrangeSummary := model.ArrangeSummary{
					FailedFiles:         failedFiles,
					FailedFileSize:      failedFileSize,
					FailedFileCount:     failedFileCount,
					SuccessfulFiles:     successfulFiles,
					SuccessfulFileSize:  successfulFileSize,
					SuccessfulFileCount: successfulFileCount,
					DuplicateFiles:      duplicateFiles,
					DuplicateFileSize:   duplicateFileSize,
					DuplicateFileCount:  duplicateFileCount,
				}
				data, err := json.Marshal(arrangeSummary)
				if err != nil {
					panic(err)
				}
				println(string(data))
			}
		}
	},
}

var arrangeOption option.ArrangeOption

func init() {
	rootCmd.AddCommand(arrangeCmd)
	f := arrangeCmd.Flags()
	f.StringVarP(&arrangeOption.Path, "path", "p", "", "specifiy the folder path")
	f.StringArrayVarP(&arrangeOption.ExcludePath, "exclude-path", "e", nil, "specifiy the folder path to exclude")
	f.StringArrayVarP(&arrangeOption.ExcludeExtension, "exclude-ext", "x", nil, "specifiy the extension to exclude")
	f.StringVarP(&arrangeOption.DestinationPath, "destination", "d", "", "specifiy the destination folder path")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// arrangeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// arrangeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
