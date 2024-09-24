package wand

import (
	"fmt"
	"github.com/FOSS-Community/wand/pkg/wand"
	"github.com/spf13/cobra"
)

var avadakedavraLogo string = fmt.Sprintf(`%s  
               _         _         _                 
 ___ _ _ ___ _| |___ ___| |_ ___ _| |___ _ _ ___ ___ 
| .'| | | .'| . | .'|___| '_| -_| . | .'| | |  _| .'|
|__,|\_/|__,|___|__,|   |_,_|___|___|__,|\_/|_| |__,|

%s`, Yellow, Reset)

var avadakedavraDescription string = `
The Killing Curse (Avada Kedavra) is a tool of the Dark Arts and one of the three Unforgivable Curses. It is one of the most powerful and sinister spells known to wizardkind. When cast successfully on a living person or creature the curse causes instantaneous and painless death, without any signs of violence on body.

It 'deletes' the person or creature from the living world, poof! And so does the 'rm' command in linux world. 'rm' causes painless death of files and folders :P

Usage : $ avada-kedavra filename

Cheers!
`

var avadakedavraCmd = &cobra.Command{
	Use:  "avada-kedavra",
	Long: fmt.Sprintf("%s %s", avadakedavraLogo, avadakedavraDescription),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]
		wand.Runcommand("rm", arg)
	},
}

func init() {
	rootCmd.AddCommand(avadakedavraCmd)
}
