package set

import (
	cmdutil "github.com/robinlg/iam/internal/iamctl/cmd/util"
	"github.com/robinlg/iam/internal/iamctl/util/templates"
	"github.com/robinlg/iam/pkg/cli/genericclioptions"
	"github.com/spf13/cobra"
)

var setLong = templates.LongDesc(`
	Configure objects.

	These commands help you make changes to existing objects.`)

// NewCmdSet returns an initialized Command instance for 'set' sub command.
func NewCmdSet(f cmdutil.Factory, ioStreams genericclioptions.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "set SUBCOMMAND",
		DisableFlagsInUseLine: true,
		Short:                 "Set specific features on objects",
		Long:                  setLong,
		Run:                   cmdutil.DefaultSubCommandRun(ioStreams.ErrOut),
	}

	// add subcommands
	// cmd.AddCommand(NewCmdDB(f, ioStreams))

	return cmd
}
