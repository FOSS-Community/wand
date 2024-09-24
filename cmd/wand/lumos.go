package wand

import (
	"fmt"
	"github.com/FOSS-Community/wand/pkg/wand"
	"github.com/spf13/cobra"
)

var lumosLogo string = fmt.Sprintf(`%s
   __                  
  / /_ ____ _  ___  ___
 / / // /  ' \/ _ \(_-<
/_/\_,_/_/_/_/\___/___/
                       
%s`, Yellow, Reset)

var lumosDescription string = `
Lumos is the incantation to a charm that can be used to produce a flash of bright white light from the tip of the wand.
This light can be thrown far off of the wand, illuminating the surrounding area for several minutes
after its initial casting before darkening once more.

Lumos shows the path in wizard world and so does the 'ls' command in linux world, 'ls' does directory listing.

So typing 'lumos' into terminal will display the files and folders in the current working directory.

Usage : $ wand lumos

Cheers!

`

var lumosCmd = &cobra.Command{
	Use:  "lumos",
	Long: fmt.Sprintf("%s %s", lumosLogo, lumosDescription),
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		wand.Runcommand("ls")
	},
}

func init() {
	rootCmd.AddCommand(lumosCmd)
}
