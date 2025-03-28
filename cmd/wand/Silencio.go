package wand

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var silencioDescription = fmt.Sprintf(`%s


⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠻⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⡦⠬⠧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⢻⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢀⣠⣤⣤⣤⣀⠀⠀⠀⠙⠀⠀⠀⠀⢀⣠⣤⣤⣤⣀⠀⠀⠀⠀
⠀⠀⢀⣶⠟⠁⠀⠀⠈⠙⣷⡄⠀⠀⠀⠀⢀⣴⠟⠉⠀⠀⠈⠙⢷⡄⠀⠀
⠀⠀⣾⢣⡗⠀⠀⠀⠀⠀⠘⣷⣴⠾⠷⣶⣼⢧⡗⠀⠀⠀⠀⠀⠈⣿⠀⠀
⠘⠛⣿⢸⡄⠀⠀⠀⠀⠀⠀⣿⠁⠀⠀⠈⣿⡸⡇⠀⠀⠀⠀⠀⠀⣾⠟⠃
⠀⠀⠹⣧⡻⣦⠀⠀⠀⢀⣼⠏⠀⠀⠀⠀⠸⣧⡹⣦⠄⠀⠀⠀⣴⠏⠀⠀
⠀⠀⠀⠈⠻⢶⣦⣤⡶⠟⠉⠀⠀⠀⠀⠀⠀⠈⠻⠷⣦⣤⡶⠟⠋⠀⠀⠀


Silencio is a silencing charm that mutes the system's sound.
You can also unmute the sound using the --unmute flag.

Usage:
$ wand silencio [--unmute]

`, Yellow)

var unmute bool

var silencioCmd = &cobra.Command{
	Use:   "silencio",
	Short: "Mute or unmute the system's sound",
	Long:  silencioDescription,
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		if unmute {
			if runtime.GOOS == "darwin" {
				fmt.Println("Unmuting the system sound...")
				err = exec.Command("osascript", "-e", "set volume without output muted").Run()
			} else {
				fmt.Println("Unmuting the system sound...")
				err = exec.Command("amixer", "set", "Master", "unmute").Run()
			}
		} else {
			if runtime.GOOS == "darwin" {
				fmt.Println("Muting the system sound...")
				err = exec.Command("osascript", "-e", "set volume with output muted").Run()
			} else {
				fmt.Println("Muting the system sound...")
				err = exec.Command("amixer", "set", "Master", "mute").Run()
			}

		}

		if err != nil {
			fmt.Println("You need to install amixer to use this spell")
			fmt.Println("Error adjusting system sound:", err)
		} else {
			if unmute {
				fmt.Println("Sound unmuted successfully.")
			} else {
				fmt.Println("Sound muted successfully.")
			}
		}
	},
}

func init() {
	// Add the --unmute flag to the silencio command
	silencioCmd.Flags().BoolVarP(&unmute, "unmute", "u", false, "Unmute the system sound")

	// Add the "Silencio" command to the root command
	rootCmd.AddCommand(silencioCmd)
}
