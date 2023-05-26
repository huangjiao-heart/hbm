package collection

import (
	collectionobj "github.com/kassisol/hbm/object/collection"
	"github.com/kassisol/hbm/pkg/adf"
	"github.com/kassisol/hbm/pkg/juliengk/go-utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "rm [name]",
		Aliases: []string{"remove"},
		Short:   "Remove collection from the whitelist",
		Long:    removeDescription,
		Args:    cobra.ExactArgs(1),
		Run:     runRemove,
	}

	return cmd
}

func runRemove(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	c, err := collectionobj.New("sqlite", adf.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer c.End()

	if err := c.Remove(args[0]); err != nil {
		log.Fatal(err)
	}
}

var removeDescription = `
Remove a collection. You cannot remove a collection that is in use by a policy.

`
