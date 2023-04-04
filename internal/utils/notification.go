package utils

import (
	"fmt"
	"os/exec"
)

func ShowNotification(title, message string) {
	iconPath := "./tomato-icon.png"
	cmd := exec.Command("powershell.exe", fmt.Sprintf("New-BurntToastNotification -AppLogo '%s' -Text '%s', '%s'", iconPath, title, message))
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return

	}
}
