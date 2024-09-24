package wand

import (
	"fmt"
	"github.com/FOSS-Community/wand/pkg/wand"
	"github.com/spf13/cobra"
)

var mobiliarbusLogo string = fmt.Sprintf(`%s
           _   _ _ _         _ 
 _____ ___| |_|_| |_|___ ___| |_ _ _ ___ 
|     | . | . | | | | .'|  _| . | | |_ -|
|_|_|_|___|___|_|_|_|__,|_| |___|___|___|

%s`, Yellow, Reset)

var mobiliarbusDescription string = `
Mobiliarbus is the incantation to a charm used to levitate and move plants and trees, as well as the materials made of wood. The Latin term mobilis, meaning "movable" , and arbor means "tree".

It moves objects from one place to another, and so does the 'mv' command in linux world. It moves files and directories from one place to another.

Usage : $ mobiliarbus /path/to/source /path/to/destination

Cheers!
`

var mobiliarbusCmd = &cobra.Command{
	Use:  "mobiliarbus",
	Long: fmt.Sprintf("%s %s", mobiliarbusLogo, mobiliarbusDescription),
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		wand.Runcommand("mv", args...)
	},
}

func init() {
	rootCmd.AddCommand(mobiliarbusCmd)
}
