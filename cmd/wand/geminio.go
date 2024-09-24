package wand

import (
	"fmt"
	"github.com/FOSS-Community/wand/pkg/wand"
	"github.com/spf13/cobra"
)

var geminoLogo string = fmt.Sprintf(`%s 
               _         
 ___ ___ _____|_|___ ___ 
| . | -_|     | |   | . |
|_  |___|_|_|_|_|_|_|___|
|___|                    

%s`, Yellow, Reset)

var geminoDescription string = `
The Geminio Curse or Doubling Charm(Geminio) is a spell used to duplicate an object. It can also be used to bewitch an object into multiplying repeatedly when touched, though how one would produce the latter effect is still unknown.

It duplicates objects, creating copiesof it and so does the 'cp' command in linux terminal, It copies files/directories in linux terminal.

Usage : $ geminio /path/to/source /path/to/destination

Cheers!
`

var geminoCmd = &cobra.Command{
	Use:  "gemino",
	Long: fmt.Sprintf("%s %s", geminoLogo, geminoDescription),
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		wand.Runcommand("cp", args...)
	},
}

func init() {
	rootCmd.AddCommand(geminoCmd)
}
