package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func init() {

	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Print the version number of sops-predictor",
	Long:  `Print the version number of sops-predictor`,
	Run: func(cmd *cobra.Command, args []string) {

		out := ReadEachLine(args[0])

		for _, i := range out {
			if i == "sops:" {
				os.Exit(1)
			}
			sizeofdata := dataCount(i)
			if itemType(i) == "boolean" {
				switch sizeofdata {
				case 4:
					fmt.Printf("%s : True\n", getKey(i))
				case 5:
					fmt.Printf("%s : False\n", getKey(i))
				}

			} else {
				fmt.Printf("%s : %d \n", getKey(i), sizeofdata)
			}

		}

	},
}

func itemType(field string) string {
	split := strings.Split(field, ",")

	for _, i := range split {

		if strings.HasPrefix(i, "type:") {

			secondSplit := strings.Split(i, ":")

			switch secondSplit[1] {
			case "bool]":
				return "boolean"
			case "str]":
				return "string"
			case "int]":
				return "int"
			}

		}
	}
	return ""
}

func ReadEachLine(filepath string) (fileLines []string) {

	readFile, err := os.Open(filepath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines
}

func dataCount(field string) int {
	split := strings.Split(field, ",")

	for _, i := range split {

		if strings.HasPrefix(i, "data:") {

			encryptedString := strings.Trim(i, "data:")

			item := item{
				fullLength:  len(encryptedString),
				paddingSize: getPaddingSize(encryptedString),
				grouping:    getGrouping(len(encryptedString)),
			}

			return item.fullLength - (item.grouping + item.paddingSize)

		}
	}
	return 0
}

func getGrouping(length int) int {
	return length / 4
}

type item struct {
	fullLength      int
	paddingSize     int
	paddinglessSize int
	grouping        int
}

func getPaddingSize(encryptedString string) int {

	if strings.HasSuffix(encryptedString, "==") {
		return 2
	}

	if strings.HasSuffix(encryptedString, "=") {
		return 1
	}

	return 0
}

func getKey(encryptedString string) string {
	split := strings.Split(encryptedString, ":")
	return split[0]
}
