package cmd

import (
	"context"

	"github.com/Swesdek/gofuzz/dnsfuzz"
	"github.com/Swesdek/gofuzz/fuzzlib"

	"github.com/spf13/cobra"
)

func fuzzDns(cmd *cobra.Command, args []string) error {
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
	wordlist, err := fuzzlib.WordListParser(wordlistFile)
	if err != nil {
		return err
	}

	dnsConfig := dnsfuzz.NewDnsConfig(context.Background(), url)

	config := fuzzlib.NewConfig(threads, wordlist, dnsConfig)
	config.Run()

	return nil
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "dns",
		Short: "Subdomain enumeration",
		RunE:  fuzzDns,
	})
}
