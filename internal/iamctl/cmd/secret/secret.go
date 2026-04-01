package secret

import (
	"github.com/olekukonko/tablewriter"
	cmdutil "github.com/robinlg/iam/internal/iamctl/cmd/util"
	"github.com/robinlg/iam/internal/iamctl/util/templates"
	"github.com/robinlg/iam/pkg/cli/genericclioptions"
	"github.com/spf13/cobra"
)

var secretLong = templates.LongDesc(`
	Secret management commands.

	This commands allow you to manage your secret on iam platform.`)

// NewCmdSecret returns new initialized instance of 'secret' sub command.
func NewCmdSecret(f cmdutil.Factory, ioStreams genericclioptions.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "secret SUBCOMMAND",
		DisableFlagsInUseLine: true,
		Short:                 "Manage secrets on iam platform",
		Long:                  secretLong,
		Run:                   cmdutil.DefaultSubCommandRun(ioStreams.ErrOut),
	}

	cmd.AddCommand(NewCmdCreate(f, ioStreams))
	cmd.AddCommand(NewCmdGet(f, ioStreams))
	cmd.AddCommand(NewCmdList(f, ioStreams))
	cmd.AddCommand(NewCmdDelete(f, ioStreams))
	cmd.AddCommand(NewCmdUpdate(f, ioStreams))

	return cmd
}

// setHeader set headers for secret commands.
func setHeader(table *tablewriter.Table) *tablewriter.Table {
	table.SetHeader([]string{"Name", "SecretID", "SecretKey", "Expires", "Created"})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.FgRedColor},
		tablewriter.Colors{tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.FgMagentaColor},
		tablewriter.Colors{tablewriter.FgGreenColor})

	return table
}
