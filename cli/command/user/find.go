package user

import (
	"fmt"

	userobj "github.com/kassisol/hbm/object/user"
	"github.com/kassisol/hbm/pkg/adf"
	"github.com/kassisol/hbm/pkg/juliengk/go-utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newFindCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "find [name]",
		Short: "Verify if user exists in the whitelist",
		Long:  findDescription,
		Args:  cobra.ExactArgs(1),
		Run:   runFind,
	}

	return cmd
}

func runFind(cmd *cobra.Command, args []string) {
	defer utils.RecoverFunc()

	u, err := userobj.New("sqlite", adf.AppPath)
	if err != nil {
		log.Fatal(err)
	}
	defer u.End()

	fmt.Println(u.Find(args[0]))
}

var findDescription = `
Verify if user exists in the whitelist

`
