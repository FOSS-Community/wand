package wand

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var reparoDescription = fmt.Sprintf(`%s

⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣀⣤⣤⣶⣶⣶⣿⣿⣿⣿⣿⣿⣷⣶⣶⣦⣤⣄⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣴⣶⣿⠿⠛⠛⠉⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠉⠛⠻⠿⣿⣷⣦⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⣶⡿⠟⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠻⣿⣷⣤⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣶⡿⠟⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠻⣿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⡿⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣴⠖⠛⠛⢿⣷⣶⣄⠀⠀⠀⠙⢿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⣠⣾⡿⠋⠀⠀⠀⠀⢀⣀⣤⣤⣤⣤⣤⣤⣄⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⣿⣿⣷⡄⠀⠀⢻⣿⣿⣷⠀⠀⠀⠀⠙⢿⣷⡄⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣴⣿⠟⠀⠀⠀⠀⣠⣾⣿⣿⣿⣿⠟⠛⠛⢿⣿⣿⣿⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠻⠿⠟⠁⠀⠀⢸⣿⣿⣿⠀⠀⠀⠀⠀⠈⠻⣿⣆⠀⠀⠀⠀⠀
⠀⠀⠀⢀⣾⣿⠃⠀⠀⠀⢠⣾⣿⣿⣿⣿⣿⠁⠀⠀⠀⠀⢻⣿⣿⣿⣿⣦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⣿⡿⠋⠀⠀⠀⠀⠀⠀⠀⠘⣿⣧⡀⠀⠀⠀
⠀⠀⢀⣾⡿⠁⠀⠀⠀⢠⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠘⣿⣿⣿⣿⣿⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠐⠶⠶⣿⣏⣁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⣷⡀⠀⠀
⠀⠀⣾⣿⠃⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⣿⣿⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⣧⠀⠀
⠀⣸⣿⠇⠀⠀⠀⠀⢰⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⡀⠀⠀⠀⠀⢀⣤⣤⣄⠀⠀⠀⢸⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⣿⡆⠀
⢀⣿⡟⠀⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⣾⣿⣿⣿⠀⠀⠀⣸⣿⣿⣿⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻⣿⠀
⢸⣿⡇⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⡀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣷⠀⠀⠀⠀⠘⢿⣿⣁⣀⣀⣴⣿⣿⠿⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠸⣿⡇
⣾⣿⠀⠀⠀⠀⠀⠀⠀⠹⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⢰⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠉⠉⠉⠉⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣧
⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠹⣿⣿⣿⣿⣿⣿⡄⠀⠀⠀⢀⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠰⠶⠶⠶⠶⠶⠶⠶⠶⠶⠶⠶⠆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿
⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠛⢿⣿⣿⣿⣿⣷⣶⣶⡿⠿⣿⣿⣿⣿⣿⣿⣿⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣤⣤⣤⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿
⢿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠛⠛⠛⠉⠁⠀⠀⣿⣿⣿⣿⣿⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⡿
⢸⣿⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⠁⠀⠀⠀⠀⠀⠀⠀⢀⣾⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⣿⡇
⠈⣿⣧⠀⠀⠀⠀⠀⠀⠀⢀⣠⣴⣶⣶⣦⣄⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⡏⠀⠀⠀⠀⠀⠀⠀⢠⣾⡟⢹⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⣿⠁
⠀⢹⣿⡄⠀⠀⠀⠀⠀⠀⣾⣿⣿⣿⣿⣿⣿⡆⠀⠀⠀⠀⢸⣿⣿⣿⣿⣿⡟⠀⠀⠀⠀⠀⠀⠀⣠⣿⡟⠀⢸⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⣿⡏⠀
⠀⠀⢿⣿⡀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⠇⠀⠀⠀⠀⣾⣿⣿⣿⣿⠟⠀⠀⠀⠀⠀⠀⠀⣴⣿⠏⠀⠀⢸⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⣿⡿⠀⠀
⠀⠀⠈⢿⣷⡀⠀⠀⠀⠀⢻⣿⣿⣿⣿⠋⠀⠀⠀⠀⢀⣼⣿⣿⣿⡿⠋⠀⠀⠀⠀⠀⠀⠀⢰⣿⣏⣀⣀⣀⣼⣿⣿⣿⣤⣤⡀⠀⠀⠀⠀⠀⠀⠀⢀⣾⡿⠁⠀⠀
⠀⠀⠀⠈⢿⣿⡄⠀⠀⠀⠀⠙⠻⢿⣿⣷⣶⣶⣶⣾⣿⣿⡿⠟⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⠉⠉⠉⠉⢻⣿⣿⣿⠋⠉⠁⠀⠀⠀⠀⠀⠀⢠⣿⡿⠁⠀⠀⠀
⠀⠀⠀⠀⠀⠻⣿⣦⠀⠀⠀⠀⠀⠀⠀⠉⠉⠉⠉⠉⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⣿⣿⣿⡄⠀⠀⠀⠀⠀⠀⢀⣴⣿⠟⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠙⢿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠐⠛⠿⠿⠿⠿⠿⠿⠿⠃⠀⠀⠀⣠⣾⡿⠋⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠙⢿⣷⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⣾⡿⠋⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢿⣿⣦⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣴⣿⠿⠋⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠛⢿⣿⣶⣤⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣶⣿⠿⠛⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠙⠻⢿⣿⣶⣦⣤⣀⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣀⣤⣴⣶⣿⠿⠟⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠙⠛⠻⠿⠿⣿⣿⣿⣿⣿⣿⣿⡿⠿⠿⠟⠛⠋⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀


Reparo is used to fix broken objects, such as repairing package dependencies.

This command will attempt to repair broken packages or dependencies based on your system's package manager.

Supported package managers:
- apt (Debian/Ubuntu)
- yum (CentOS/Fedora)
- pacman (Arch Linux)
- brew (macOS)
`, Yellow)

var reparoCmd = &cobra.Command{
	Use:   "reparo",
	Short: "Repair broken packages or dependencies",
	Long:  reparoDescription,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		packageManager := detectPackageManager()

		switch packageManager {
		case "apt":
			fmt.Println("Repairing broken packages using apt...")
			err = exec.Command("sudo", "apt-get", "install", "--fix-broken").Run()
		case "yum":
			fmt.Println("Repairing broken packages using yum...")
			err = exec.Command("sudo", "yum", "check").Run()
		case "pacman":
			fmt.Println("Repairing broken packages using pacman...")
			err = exec.Command("sudo", "pacman", "-Syu").Run()
		case "brew":
			fmt.Println("Repairing broken packages using brew...")
			err = exec.Command("brew", "doctor").Run()
		default:
			fmt.Println("Unsupported package manager. Please use apt, yum, pacman, or brew.")
		}

		if err != nil {
			fmt.Println("Error repairing packages:", err)
		} else {
			fmt.Println("Package repair completed successfully.")
		}
	},
}

func detectPackageManager() string {
	// Detect the package manager based on the system type
	if runtime.GOOS == "linux" {
		// Check for apt, yum, or pacman
		if _, err := exec.LookPath("apt-get"); err == nil {
			return "apt"
		} else if _, err := exec.LookPath("yum"); err == nil {
			return "yum"
		} else if _, err := exec.LookPath("pacman"); err == nil {
			return "pacman"
		}
	} else if runtime.GOOS == "darwin" {
		// Check for Homebrew on macOS
		if _, err := exec.LookPath("brew"); err == nil {
			return "brew"
		}
	}

	return ""
}

func init() {
	// Add the "Reparo" command to the root command
	rootCmd.AddCommand(reparoCmd)
}
