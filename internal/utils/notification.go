package utils

import (
	"fmt"
	"os/exec"
)

func ShowNotification(title, message string) {
	iconPath := "./gopher.ico"
	cmd := exec.Command("powershell.exe", "-Command", fmt.Sprintf(`& {[reflection.assembly]::loadwithpartialname('System.Windows.Forms'); [Windows.Forms.NotifyIcon]$n = New-Object Windows.Forms.NotifyIcon; $n.Icon = [System.Drawing.Icon]::ExtractAssociatedIcon("%s"); $n.BalloonTipIcon = [System.Windows.Forms.ToolTipIcon]::Info; $n.BalloonTipTitle = "%s"; $n.BalloonTipText = "%s"; $n.Visible = $true; $n.ShowBalloonTip(5000); Start-Sleep -s 5; $n.Dispose();}`, iconPath, title, message))
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return

	}
}
