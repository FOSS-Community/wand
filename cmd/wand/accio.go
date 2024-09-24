package wand

import (
	"fmt"
	"github.com/FOSS-Community/wand/pkg/wand"
	"github.com/spf13/cobra"
)

var accioLogo string = fmt.Sprintf(`%s   
             _     
 ___ ___ ___|_|___ 
| .'|  _|  _| | . |
|__,|___|___|_|___|

%s`, Yellow, Reset)

var accioDescription string = `
The Summoning Charm (Accio) was a charm that caused a target at a distance from the caster to levitate or fly over to them. This spell needs thought behind it, the object must be clear in the caster's mind before trying to summon.

This spell is one of the oldest spells known to wizarding society.

Accio 'gets' the object and so does 'wget' in the linux world.

Usage : $ accio LINK

Cheers!
`

var accioCmd = &cobra.Command{
	Use:  "accio",
	Long: fmt.Sprintf("%s %s", accioLogo, accioDescription),
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := args[0]
		wand.Runcommand("wget", arg)
	},
}

func init() {
	rootCmd.AddCommand(accioCmd)
}
