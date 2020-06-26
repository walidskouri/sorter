package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"training.go/sortlines/sorter"
)

const version = "0.0.1"

type Options struct {
	fileName string
	suffix   string
	version  bool
}

var opts = &Options{
	fileName: "file.txt",
	suffix:   "sorted",
}

func Run() {
	cmd := &cobra.Command{}
	cmd.Use = "sort"
	cmd.Short = "Sort the input file lines alphabetically"
	cmd.Flags().StringVarP(&opts.fileName, "file-name", "f", opts.fileName, "The input file name")
	cmd.Flags().StringVarP(&opts.suffix, "sorted-file-suffix", "s", opts.suffix, "The output file name suffix")
	cmd.Flags().BoolVarP(&opts.version, "version", "v", opts.version, "Print the version and exit")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {


		if opts.version {
			fmt.Printf("sorter version %s\n", version)
			return nil
		}
		narg := len(args)
		if (narg > 1) || (narg == 0 && opts.fileName == "") {
			return cmd.Help()
		}
		config := parseConfig()

		err := sorter.Run(config)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		return nil
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func parseConfig() *sorter.Config {
	return &sorter.Config{
		FileName: opts.fileName,
		Suffix:   opts.suffix,
	}

}
