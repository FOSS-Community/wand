package wand

import (
	"fmt"
	"github.com/FOSS-Community/wand/pkg/wand"
	"github.com/spf13/cobra"
)

var mobiliarbusLogo string = fmt.Sprintf(`%s
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣰⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⢿⡿⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⠃⢸⡇⠙⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢀⣼⠃⠀⢸⡇⠀⠘⣧⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢀⡾⣁⣤⡴⢾⡷⢦⣤⡘⣷⡀⠀⠀⠀⠀⠀⠀⠀⢀⣴⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⢀⣾⡿⠋⠀⠀⢸⡇⠀⠀⠙⢿⣷⡀⠀⠀⠀⠀⠀⠀⢸⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢠⣿⡿⠀⠀⠀⠀⢸⡇⠀⠀⠀⠈⢿⣿⡄⠀⠀⠀⠀⠀⢸⣿⡇⢀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣶⣿⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣴⣿⣦⡀
⠀⠀⠀⠀⢠⡟⢹⡇⠀⠀⠀⠀⢸⡇⠀⠀⠀⠀⢸⡏⢿⡄⠀⠀⠀⠀⢸⣿⡇⠈⣿⣿⠀⠀⣠⠀⠈⣿⣿⠃⠻⡏⠀⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⣿⡀⠀⣿⣷
⠀⠀⠀⣠⡟⠀⢸⣇⠀⠀⠀⠀⢸⡇⠀⠀⠀⠀⣸⠇⠀⢻⣆⠀⠀⠀⢸⣿⡇⠀⣿⣿⠀⢸⣿⠀⠀⣿⣿⠀⢀⣤⠆⣿⣿⠀⢠⣤⣤⠀⠀⢰⣶⠆⠘⢿⣿⣄⠁⠀
⠀⠀⣰⠟⠀⠀⠀⢻⣄⠀⠀⠀⢸⡇⠀⠀⠀⣠⡟⠀⠀⠀⢻⣆⠀⠀⢸⣿⡇⠀⣿⣿⠀⢸⣿⠀⠀⣿⣿⠀⣿⣿⠀⣿⣿⠀⠈⢿⣿⠀⢠⣿⠏⠀⣀⡀⠻⣿⣷⡀
⠀⣰⠏⠀⠀⠀⠀⠀⠙⠷⣤⣀⣸⣇⣀⣤⠾⠋⠀⠀⠀⠀⠀⠹⣆⠀⢸⣿⡇⠀⣿⣿⣀⣾⣿⣀⣠⡿⠃⠀⢹⣿⣧⣿⣿⠦⠀⠸⣿⣷⣿⠏⠀⠀⣿⣷⡀⣹⣿⠟
⠼⠿⠶⠶⠶⠶⠶⠶⠶⠶⠶⠿⠿⠿⠿⠶⠶⠶⠶⠶⠶⠶⠶⠶⠿⠆⢸⣿⡇⠀⠈⠻⠟⠁⠈⠻⠋⠀⠀⠀⠀⠛⠉⠛⠉⠀⠀⠀⣿⣿⠏⠀⠀⠀⠈⢻⡿⠟⠁⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⠻⢇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣴⡶⠒⠀⠀⠀⣼⣿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⠻⠿⢿⣷⣾⡿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
%s`, Green, Reset)

var mobiliarbusDescription string = fmt.Sprintf(`%s
%sMobiliarbus is the incantation to a charm used to levitate and move plants and trees, as well as materials made of wood.%s

The Latin term 'mobilis', meaning "movable," and 'arbor', meaning "tree."

%sIt moves objects from one place to another, and so does the 'mv' command in the Linux world.%s %s'mv' moves files and directories from one location to another.%s

Usage : %s$ wand mobiliarbus /path/to/source /path/to/destination%s

%sCheers!%s
`, Cyan, Green, Reset, Yellow, Reset, Blue, Reset, Red, Reset, Magenta, Reset)

var mobiliarbusCmd = &cobra.Command{
	Use:  "mobiliarbus",
	Long: fmt.Sprintf("%s %s", mobiliarbusLogo, mobiliarbusDescription),
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		source := args[0]
		destination := args[1]
		wand.Runcommand("mv", source, destination)
	},
}

func init() {
	rootCmd.AddCommand(mobiliarbusCmd)
}
