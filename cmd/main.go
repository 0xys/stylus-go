package main

import (
	"fmt"
	"os"

	"github.com/0xys/stylus-go/lib"
	"github.com/spf13/cobra"
)

const (
	ShowTemporaryWork      = "work"
	ShowTemporaryWorkShort = "w"
	OutputPath             = "out"
	OutputPathShort        = "o"
)

var rootCmd = &cobra.Command{
	Use:  "asgo [flags] filepath",
	Long: "Arbitrum Stylus Go: go smart contract generator for Arbitrum Stylus",
	RunE: func(cmd *cobra.Command, args []string) error {
		t, err := cmd.Flags().GetBool(ShowTemporaryWork)
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
		out, err := os.Create(outpath)
		if err != nil {
			return err
		}
		defer func() {
			out.Close()
		}()

		cont, err := lib.ProcessFile(args[0], t)
		if err != nil {
			return err
		}

		err = lib.GenContract(cont, out)
		if err != nil {
			return err
		}
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
	rootCmd.Flags().BoolP(ShowTemporaryWork, ShowTemporaryWorkShort, false, "print the name of the temporary work directory and do not delete it when exiting")
	rootCmd.Flags().StringP(OutputPath, OutputPathShort, "entrypoint.go", "output path of the contract entrypoint")
}
