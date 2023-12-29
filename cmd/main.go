package main

import (
	"fmt"
	"os"

	"example.com/lib"
	"github.com/spf13/cobra"
)

const (
	KeepTemporaryWork      = "work"
	KeepTemporaryWorkShort = "w"
	OutputPath             = "out"
	OutputPathShort        = "o"
)

var rootCmd = &cobra.Command{
	Use:  "asgo [flags] filepath",
	Long: "Arbitrum Stylus Go: go smart contract generator for Arbitrum Stylus",
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := cmd.Flags().GetBool(KeepTemporaryWork)
		if err != nil {
			return err
		}
		outpath, err := cmd.Flags().GetString(OutputPath)
		if err != nil {
			return err
		}

		if len(args) > 1 {
			return fmt.Errorf("maximum of one argument is accepted")
		}

		err = lib.ProcessFile(args[0], outpath)
		if err != nil {
			return err
		}

		fmt.Printf("toggle: %t\n", t)
		return nil
	},
	Args: cobra.MinimumNArgs(1),
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP(KeepTemporaryWork, KeepTemporaryWorkShort, false, "print the name of the temporary work directory and do not delete it when exiting")
	rootCmd.Flags().StringP(OutputPath, OutputPathShort, "entrypoint.go", "output path of the contract entrypoint")
}
