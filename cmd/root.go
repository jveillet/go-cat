/*
Copyright © 2023 Jérémie Veillet <jeremie.veillet@gmail.com>
*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// Holds the value of the -n command option
var number bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cat [OPTIONS]... [FILE]...",
	Short: "Concatenate FILE(s) to standard output",
	Long:  `With no FILE, or when FILE is -, read standard input.`,
	Args:  cobra.MinimumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 || args[0] == "-" {
			if err := readStandardInput(); err != nil {
				return err
			}
		} else {
			if err := readFile(args[0], os.Stdout, number); err != nil {
				return err
			}
		}

		return nil
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-cat.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolVarP(&number, "number", "n", false, "number all output lines")
}

// Read a file on disk, then display its content in a data stream (example Stdout).
func readFile(filePath string, writer io.Writer, displayNumbers bool) error {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return err
	} else {
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		// Reads file into a buffer
		scanner := bufio.NewScanner(file)

		var line int = 1
		for scanner.Scan() {
			// If the command line argument -n is given to the command
			// Prints the line number + <TAB> + line content + <LF>
			if displayNumbers {
				fmt.Fprintf(writer, "%d\t%s\n", line, scanner.Text())
			} else {
				// Prints line content + <LF>
				fmt.Fprintf(writer, "%s\n", scanner.Text())
			}
			line++
		}

		if err := scanner.Err(); err != nil {
			return err
		}
	}

	return nil
}

// Read from the standard OS input (STDIN) and write back to the standard OS output (STDOUT)
// More info on why this function exists in the original `cat` unix command:
// https://retrocomputing.stackexchange.com/questions/26641/why-does-cat-with-no-argument-read-from-standard-input
func readStandardInput() error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
