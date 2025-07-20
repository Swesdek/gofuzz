package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "github.com/Swesdek/gofuzz",
	Short: "Web fuzzer",
	Long:  "Fast fuzzer for web written in Go",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().IntP("threads", "t", 10, "Amount of threads used for fuzzing")
	rootCmd.PersistentFlags().StringP("wordlist", "w", "", "Path to wordlist")
	rootCmd.PersistentFlags().StringP("url", "u", "", "Url to fuzz")
	rootCmd.MarkFlagFilename("wordlist")
	rootCmd.MarkFlagRequired("wordlist")
	rootCmd.MarkFlagRequired("threads")
	rootCmd.MarkFlagRequired("target")
}
