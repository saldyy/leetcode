/*
Copyright © 2024 NAME HERE png9981@gmail.com
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

const executionCodeTemplate = `package problem{problemName}

func Execute({input}) {expectingReturnType} {

  return 0
}
`

const testingTemplate = `package problem{problemName}

import "testing"

type Input struct {
}

type Case struct {
	name     string
	input    Input
	expected {expectingReturnType}
}

func TestExecute(t *testing.T) {
	cases := []Case{}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			input := c.input
			result := Execute(input)
			if result != c.expected {
				t.Errorf("\x1b[31mFAILED %s: got=%d, want=%d\x1b[0m",
					c.name, result, c.expected)
			} else {
				t.Logf("\x1b[32mPASSED %s: got=%d\x1b[0m",
					c.name, result)
			}
		})
	}
}
`

var problemName string
var expectingReturnType string
var numberOfInputParams int
var inputArgsStr string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a template to solve a Leetcode problem",
	Long: `
The generate command is used to create a boilerplate template for solving LeetCode problems. It generates two files for a given problem:

Execution File: Contains the solution template, including function stubs and imports.
Test File: Includes a basic test structure to validate the solution.
This command helps streamline the workflow for solving LeetCode problems by providing a consistent and reusable structure

`,
	Run: func(cmd *cobra.Command, args []string) {

		problemName = promptInteractiveInput("Enter the problem's name (required): ", "").(string)
		expectingReturnType = promptInteractiveInput("Enter the expecting return type (required): ", "").(string)
		numberOfInputParams = promptInteractiveInput("Enter the number of params(required): ", 0).(int)
		var b strings.Builder
		var argsBuilder strings.Builder

		for i := 1; i <= numberOfInputParams; i++ {
			b.WriteString("Enter ")
			b.WriteString(string(i))
			b.WriteString(" type: ")
			it := promptInteractiveInput(b.String(), "").(string)
			b.Reset()

			argsBuilder.WriteString(string((i - 1) + 'a'))
			argsBuilder.WriteString(" ")
			argsBuilder.WriteString(it)
			if i != numberOfInputParams {
				argsBuilder.WriteString(", ")
			}
		}

		inputArgsStr = argsBuilder.String()

		fmt.Printf("P: %s\n", problemName)
		fmt.Printf("E: %s\n", expectingReturnType)
		fmt.Printf("%d\n", numberOfInputParams)
		fmt.Printf("%s\n", inputArgsStr)

		writeExecutionTemplate()
		writeTestingTemplate()
	},
}

func promptInteractiveInput(prompt string, inputType interface{}) interface{} {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf(prompt)
		rawInput, _ := reader.ReadString('\n')
		rawInput = strings.TrimSpace(rawInput)

		switch reflect.TypeOf(inputType).Kind() {
		case reflect.String:
			fmt.Printf("%s\n", rawInput)
			if rawInput == "" {
				fmt.Printf("Input cannot be empty. Please try again.\n")
				break
			}
			return rawInput
		case reflect.Int:
			parsed, err := strconv.Atoi(rawInput)
			if err != nil {
				fmt.Printf("Input must be an integer. Please try again.\n")
				break
			}
			return parsed
		default:
			fmt.Printf("Unsupported input type\n")
			panic(1)
		}
	}
}

func writeExecutionTemplate() {

	folderPath := fmt.Sprintf("./problems/%04s", problemName)
	fileName := fmt.Sprintf("%04s.go", problemName)

	content := strings.Replace(executionCodeTemplate, "{problemName}", problemName, -1)
	content = strings.Replace(content, "{input}", inputArgsStr, -1)
	content = strings.Replace(content, "{expectingReturnType}", expectingReturnType, -1)

  checkFolder(folderPath)
	writeFile(folderPath, fileName, content)
}

func writeTestingTemplate() {

	folderPath := fmt.Sprintf("./problems/%04s", problemName)
	fileName := fmt.Sprintf("%04s_test.go", problemName)

	content := strings.Replace(testingTemplate, "{problemName}", problemName, -1)
	content = strings.Replace(content, "{input}", inputArgsStr, -1)
	content = strings.Replace(content, "{expectingReturnType}", expectingReturnType, -1)

	writeFile(folderPath, fileName, content)
}

func checkFolder(folderPath string) {
	fileCheck := "Y"

	if info, err := os.Stat(folderPath); err == nil {
		fmt.Printf("Size: %d", info.Size())
		for {
			if info.Size() == 0 {
				break
			}

			o := promptInteractiveInput("There is an existing files. Overwrite it [Y/N]: ", "").(string)
			if o == "Y" || o == "N" {
				fileCheck = o
				break
			}
			fmt.Printf("Input must be Y or N.\n")

		}
	}
	if fileCheck == "N" {
		fmt.Printf("Cancelled.\n")
		os.Exit(0)
	}

	os.Mkdir(folderPath, os.ModeDir)
}

func writeFile(folderPath string, fileName string, content string) {
	fullPath := folderPath + "/" + fileName

	f, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY, 0644)

	defer f.Close()

	if err != nil {
		log.Fatal("Cannot create file", err)
	}

	_, err = f.WriteString(content)

	if err != nil {
		log.Fatal("Cannot write file", err)
	}
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
