package wand

import (
	"fmt"
	"github.com/FOSS-Community/wand/pkg/wand"
	"github.com/spf13/cobra"
)

var legilimensLogo string = fmt.Sprintf(`%s 
 _         _ _ _                   
| |___ ___|_| |_|_____ ___ ___ ___ 
| | -_| . | | | |     | -_|   |_ -|
|_|___|_  |_|_|_|_|_|_|___|_|_|___|
      |___|                        

%s`, Yellow, Reset)

var legilimensDescription string = `
Legilemency or legilimens is the act of magically navigating through many layers of a person's mind and correctly interpreting one's findings. Muggles often call this as 'mind-reading'.

It reads what is beneath the flesh of the person in wizard world, and somewhat 'cat' does the same in linux world. It reads the contents hidden beneath the flesh of a file :P

Usage : $ legilimens filename

Cheers!
`

var legilimensCmd = &cobra.Command{
	Use:  "legilimens",
	Long: fmt.Sprintf("%s %s", legilimensLogo, legilimensDescription),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]
		wand.Runcommand("cat", arg)
	},
}

func init() {
	rootCmd.AddCommand(legilimensCmd)
}
