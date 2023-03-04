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
	Short: "Run the sops-predictor program",
	Long:  `Run the sops-predictor program`,
	Run: func(cmd *cobra.Command, args []string) {

		lines := ReadEachLine(args[0])

		for _, line := range lines {

			// sops puts a management object at the bottom of its yaml
			// none of this is important, so I just exit the application here.
			if line == "sops:" {
				os.Exit(1)
			}

			values := parseLine(line)

			values.unencryptedLength = dataCount(values.data)

			switch values._type {
			case "bool":
				switch values.unencryptedLength {
				case 4:
					fmt.Printf("%s:true:bool \n", values.name)
				case 5:
					fmt.Printf("%s:false:bool \n", values.name)
				}
			case "str":
				fmt.Printf("%s:%d:string \n", values.name, values.unencryptedLength)
			case "int":
				fmt.Printf("%s:%d:int \n", values.name, values.unencryptedLength)
			}

		}

	},
}

//
//		import
//
//

// ReadEachLine standard read text file function
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

//
//		Math
//
//

// dataCount simple maths function that applies the pattern observed in the unencrypted string.
func dataCount(data string) int {

	return len(data) - (getGrouping(len(data)) + getPaddingSize(data))
}

// getGrouping applies the pattern applied by sops that increased the difference between the unencrypted string
// and the encrypted one. explained in the readme.
func getGrouping(length int) int {
	return length / 4
}

// getPaddingSize the full encrypted string contains padding that helps identify where the encrypted string sits in the 3 set
func getPaddingSize(data string) int {

	if strings.HasSuffix(data, "==") {
		return 2
	}

	if strings.HasSuffix(data, "=") {
		return 1
	}

	return 0
}

//
//		Data
//
//

type SopsEncDataSet struct {
	name              string
	encryption        string
	data              string
	iv                string
	tag               string
	_type             string
	unencryptedLength int
}

// parseLine manages raw text line to object conversion.
// with below helper functions
func parseLine(dataSetString string) SopsEncDataSet {

	return SopsEncDataSet{
		name:       getKey(dataSetString),
		encryption: getEncryption(dataSetString),
		data:       getData(dataSetString, "data:"),
		iv:         getData(dataSetString, "iv:"),
		tag:        getData(dataSetString, "tag:"),
		_type:      getData(dataSetString, "type:"),
	}

}

func getKey(dataSetString string) string {
	split := strings.Split(dataSetString, ":")
	return split[0]
}

func getEncryption(dataSetString string) string {
	split := strings.Split(dataSetString, ",")
	split = strings.Split(split[0], "[")
	return split[0]
}

func getData(dataSetString string, dataToCollect string) string {
	split := strings.Split(dataSetString, ",")

	for _, i := range split {

		if strings.HasPrefix(i, dataToCollect) {

			//clean the end of the type value
			if dataToCollect == "type:" {
				i = strings.TrimSuffix(i, "]")
			}

			return strings.TrimPrefix(i, dataToCollect)
		}

	}
	return ""
}
