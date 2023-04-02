package main

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"os"
	"os/exec"
	"time"
)

const (
	workMinutes  = 1
	breakMinutes = 1
)

func main() {

	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	workDuration := time.Duration(workMinutes) * time.Minute
	breakDuration := time.Duration(breakMinutes) * time.Minute
	rounds := 72 // 24-hour cycle

	// Create widgets
	timer := widgets.NewParagraph()
	timer.Title = " Pomodoro Timer \U0001F345 "
	timer.Text = formatTime(workDuration)
	timer.TextStyle.Fg = ui.ColorGreen
	timer.SetRect(0, 0, 51, 10)

	status := widgets.NewParagraph()
	status.Title = " Status "
	status.Text = "Working"
	status.TextStyle.Fg = ui.ColorYellow
	status.SetRect(1, 3, 50, 6)

	progress := widgets.NewGauge()
	progress.Title = " Progress "
	progress.Percent = 0
	progress.BarColor = ui.ColorBlue
	progress.SetRect(1, 6, 50, 9)

	ui.Render(timer)

	events := ui.PollEvents()
	go func() {
		for {
			e := <-events
			if e.Type == ui.KeyboardEvent {
				switch e.ID {
				case "q", "<C-c>":
					ui.Close()
					fmt.Println("Exiting Pomodoro CLI...")
					os.Exit(0)
					return
				}
			}
		}
	}()

	for round := 1; round <= rounds; round++ {
		runPomodoro(timer, status, progress, workDuration, "Working")
		showNotification("Pomodoro Timer", "Time for a break!")
		runPomodoro(timer, status, progress, breakDuration, "Break")
		showNotification("Pomodoro Timer", "Time to get back to work!")
	}

	// Show completed message
	status.Text = "Done"
	ui.Render(timer)
	time.Sleep(3 * time.Second)
}

func runPomodoro(timer *widgets.Paragraph, status *widgets.Paragraph, progress *widgets.Gauge, duration time.Duration, statusText string) {
	startTime := time.Now()
	endTime := startTime.Add(duration)
	ticker := time.NewTicker(time.Second)

	for range ticker.C {
		remaining := endTime.Sub(time.Now())
		if remaining <= 0 {
			break
		}
		timer.Text = formatTime(remaining)
		progress.Percent = int((duration - remaining) * 100 / duration)
		status.Text = statusText
		ui.Render(timer, status, progress)
	}

	// Show completed message
	timer.Text = formatTime(0)
	status.Text = fmt.Sprintf("%s - Completed", statusText)
	progress.Percent = 100
	ui.Render(timer, status, progress)

	// Wait for a second to show the completed message
	time.Sleep(1 * time.Second)
}

func showNotification(title, message string) {
	iconPath := "./tomato-icon.png"
	cmd := exec.Command("powershell.exe", fmt.Sprintf("New-BurntToastNotification -AppLogo '%s' -Text '%s', '%s'", iconPath, title, message))
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func formatTime(duration time.Duration) string {
	minutes := int(duration.Minutes())
	seconds := int(duration.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
