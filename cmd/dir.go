package cmd

import (
	"time"

	"github.com/Swesdek/gofuzz/dirfuzz"
	"github.com/Swesdek/gofuzz/fuzzlib"

	"github.com/spf13/cobra"
)

func fuzzDir(cmd *cobra.Command, args []string) error {
	threads, err := cmd.Flags().GetInt("threads")
	if err != nil {
		return err
	}
	wordlistFile, err := cmd.Flags().GetString("wordlist")
	if err != nil {
		return err
	}
	url, err := cmd.Flags().GetString("url")
	if err != nil {
		return err
	}
	timeoutSeconds, err := cmd.Flags().GetInt("timeout")
	if err != nil {
		return err
	}

	timeout := time.Duration(timeoutSeconds) * time.Second

	wordlist, err := fuzzlib.WordListParser(wordlistFile)
	if err != nil {
		return err
	}

	dirConfig := dirfuzz.NewDirConfig(url, timeout)

	config := fuzzlib.NewConfig(threads, wordlist, dirConfig)
	config.Run()

	return nil
}

func init() {
	command := &cobra.Command{
		Use:   "dir",
		Short: "Directory fuzzing",
		RunE:  fuzzDir,
	}
	command.Flags().Int("timeout", 2, "Timeout in seconds ")
	rootCmd.AddCommand(command)
}
