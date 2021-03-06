package cmd

import (
	"fmt"

	"github.com/boson-project/faas"
	"github.com/boson-project/faas/knative"
	"github.com/ory/viper"
	"github.com/spf13/cobra"
)

func init() {
	root.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:        "list",
	Short:      "Lists deployed Functions",
	Long:       `Lists deployed Functions`,
	SuggestFor: []string{"ls"},
	RunE:       list,
}

func list(cmd *cobra.Command, args []string) (err error) {
	verbose := viper.GetBool("verbose")

	lister, err := knative.NewLister(faas.DefaultNamespace)
	if err != nil {
		return
	}
	lister.Verbose = verbose

	client, err := faas.New(
		faas.WithVerbose(verbose),
		faas.WithLister(lister),
	)
	if err != nil {
		return
	}

	names, err := client.List()
	if err != nil {
		return
	}
	for _, name := range names {
		fmt.Printf("%s\n", name)
	}
	return
}
