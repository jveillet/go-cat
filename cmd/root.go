/*
Copyright © 2023 Jérémie Veillet <jeremie.veillet@gmail.com>
*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var number bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-cat [OPTIONS]... [FILE]...",
	Short: "Concatenate FILE(s) to standard output",
	Long:  `With no FILE, or when FILE is -, read standard input.`,
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := readFile(args[0]); err != nil {
			return err
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

// Read a file on disk, then display its content in the standard output.
// Example: readFile("myfile.txt")
func readFile(filePath string) error {
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
			// Prints the line number + <TAB> + line content + CR + LF
			if number {
				fmt.Printf("%d\t%s\n\r", line, scanner.Text())
			} else {
				// Prints line content + CR + LF
				fmt.Printf("%s\n\r", scanner.Text())
			}
			line++
		}

		if err := scanner.Err(); err != nil {
			return err
		}
	}

	return nil
}
