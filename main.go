package main

import (
	"fmt"
	"time"

	"github.com/gen2brain/beeep"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
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
	rounds := 4

	// Create widgets
	timer := widgets.NewParagraph()
	timer.Title = "Pomodoro Timer \U0001F345"
	timer.Text = formatTime(workDuration)
	timer.TextStyle.Fg = ui.ColorGreen
	timer.SetRect(0, 0, 50, 3)

	status := widgets.NewParagraph()
	status.Title = "Status"
	status.Text = "Working"
	status.TextStyle.Fg = ui.ColorYellow
	status.SetRect(0, 3, 50, 6)

	progress := widgets.NewGauge()
	progress.Title = "Progress"
	progress.Percent = 0
	progress.BarColor = ui.ColorBlue
	progress.SetRect(0, 6, 50, 9)

	ui.Render(timer, status, progress)

	for round := 1; round <= rounds; round++ {
		runPomodoro(timer, status, progress, workDuration, "Working")
		showNotification("Pomodoro Timer", "Time for a break!")
		playSound()
		runPomodoro(timer, status, progress, breakDuration, "Break")
		showNotification("Pomodoro Timer", "Time to get back to work!")
		playSound()
	}

	// Show completed message
	status.Text = "Done"
	ui.Render(timer, status, progress)
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

func playSound() {
	for i := 0; i < 3; i++ {
		err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(500 * time.Millisecond) // Wait for 500 milliseconds between each beep
	}
}

func showNotification(title, message string) {
	iconPath := "./tomato-icon.png"
	err := beeep.Notify(title, message, iconPath)
	if err != nil {
		return
	}
}

func formatTime(duration time.Duration) string {
	minutes := int(duration.Minutes())
	seconds := int(duration.Seconds()) % 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
