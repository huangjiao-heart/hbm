package resource

import (
	resourceobj "github.com/kassisol/hbm/object/resource"
	"github.com/kassisol/hbm/pkg/adf"
	"github.com/kassisol/hbm/pkg/juliengk/go-utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	resourceMemberAdd    bool
	resourceMemberRemove bool
)

func newMemberCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "member [collection] [resource]",
		Short: "Manage resource membership to collection",
		Long:  memberDescription,
		Args:  cobra.ExactArgs(2),
		Run:   runMember,
	}

	flags := cmd.Flags()
	flags.BoolVarP(&resourceMemberAdd, "add", "a", false, "Add resource to collection")
	flags.BoolVarP(&resourceMemberRemove, "remove", "r", false, "Remove resource from collection")

	return cmd
}

func runMember(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	r, err := resourceobj.New("sqlite", adf.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer r.End()

	if resourceMemberAdd {
		r.AddToCollection(args[1], args[0])
	}
	if resourceMemberRemove {
		r.RemoveFromCollection(args[1], args[0])
	}
}

var memberDescription = `
Manage resource membership to collection

`
