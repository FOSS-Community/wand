package wand

import (
	"fmt"
	"github.com/FOSS-Community/wand/pkg/wand"
	"github.com/spf13/cobra"
)

var obliviateLogo string = fmt.Sprintf(`%s
     _   _ _     _     _       
 ___| |_| |_|_ _|_|___| |_ ___ 
| . | . | | | | | | .'|  _| -_|
|___|___|_|_|\_/|_|__,|_| |___|

%s`, Yellow, Reset)

var obliviateDescription string = `
Originating from Harry Potter and the Chamber of Secrets, Obliviate is a term meaning 'forget'. Obliviate is a memory charm, resulting in the erasure of the recipients memory. One's memory vanishes as soon as this charm is casted.

It makes one forgetful of things or memories and so does the 'clear' command in linux terminal, It clears the output screen or 'obliviates' it :P

Usage : $ obliviate

Cheers!

`

var obliviateCmd = &cobra.Command{
	Use:  "obliviate",
	Long: fmt.Sprintf("%s %s", obliviateLogo, obliviateDescription),
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		wand.Runcommand("clear")
	},
}

func init() {
	rootCmd.AddCommand(obliviateCmd)
}
