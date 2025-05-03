package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var showLines, showWords, showBytes, showChars bool

var wcCmd = &cobra.Command{
	Use:   "wc [file]",
	Short: "Counts lines, words, bytes, or characters from a file or stdin",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var reader io.Reader
		filename := ""
		if len(args) == 1 {
			filename = args[0]
			file, err := os.Open(args[0])
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
				os.Exit(1)
			}
			defer file.Close()
			reader = file
		} else {
			reader = os.Stdin
		}
		lines, words, bytes, characters := count(reader)

		if showLines {
			fmt.Printf("%d %s\n", lines, filename)
		} else if showWords {
			fmt.Printf("%d %s\n", words, filename)
		} else if showBytes {
			fmt.Printf("%d %s\n", bytes, filename)
		} else if showChars {
			fmt.Printf("%d %s\n", characters, filename)
		} else {
			fmt.Printf("%d %d %d %s\n", lines, words, bytes, filename)
		}
	},
}

func count(r io.Reader) (int, int, int, int) {
	scanner := bufio.NewScanner(r)
	lines, words, bytes, characters := 0, 0, 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		lines++
		words += len(strings.Fields(line))
		for _, ch := range line {
			if ch != ' ' {
				characters++
			}
		}
		bytes += len(line) + 1 // assuming newline
	}
	return lines, words, bytes, characters
}

func init() {
	rootCmd.AddCommand(wcCmd)

	// Flags
	wcCmd.Flags().BoolVarP(&showLines, "lines", "l", false, "Print the newline counts")
	wcCmd.Flags().BoolVarP(&showWords, "words", "w", false, "Print the word counts")
	wcCmd.Flags().BoolVarP(&showBytes, "bytes", "c", false, "Print the byte counts")
	wcCmd.Flags().BoolVarP(&showChars, "chars", "m", false, "Print the character counts (excluding spaces)")
}
