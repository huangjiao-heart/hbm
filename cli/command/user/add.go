package user

import (
	userobj "github.com/kassisol/hbm/object/user"
	"github.com/kassisol/hbm/pkg/adf"
	"github.com/kassisol/hbm/pkg/juliengk/go-utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [name]",
		Short: "Add user to the whitelist",
		Long:  addDescription,
		Args:  cobra.ExactArgs(1),
		Run:   runAdd,
	}

	return cmd
}

func runAdd(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	u, err := userobj.New("sqlite", adf.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer u.End()

	if err := u.Add(args[0]); err != nil {
		log.Fatal(err)
	}
}

var addDescription = `
Add user to the whitelist

`
