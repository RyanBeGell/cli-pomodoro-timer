package main

import (
	"cli-pomodoro-timer/internal/utils"
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"os"
	"time"
)

const (
	workMinutes      = 1
	breakMinutes     = 1
	longBreakMinutes = 1
)

func main() {

	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	workDuration := time.Duration(workMinutes) * time.Minute
	breakDuration := time.Duration(breakMinutes) * time.Minute
	longBreakDuration := time.Duration(longBreakMinutes) * time.Minute
	// Create widgets
	timer := widgets.NewParagraph()
	timer.Title = " Pomodoro Timer \U0001F345 "
	timer.SetRect(0, 0, 51, 12)

	status := widgets.NewParagraph()
	status.Title = " Status "
	status.Text = "Working"
	status.TextStyle.Fg = ui.ColorGreen
	status.SetRect(2, 2, 49, 5)

	progress := widgets.NewGauge()
	progress.Title = " Progress "
	progress.Percent = 0
	progress.BarColor = ui.ColorBlue
	progress.SetRect(2, 5, 49, 8)

	dataBox := widgets.NewParagraph()
	dataBox.Border = false
	dataBox.Text = "PRESS q TO CLOSE"
	dataBox.TextStyle.Fg = ui.ColorRed
	dataBox.SetRect(2, 8, 49, 11)

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

	cycleCount := 0
	longBreakCount := 0
	for {
		runPomodoro(timer, status, progress, dataBox, workDuration, "Working")
		if (cycleCount+1)%4 == 0 {
			longBreakCount++
			utils.ShowNotification("Pomodoro Timer", "Time for a long break!")
			runPomodoro(timer, status, progress, dataBox, longBreakDuration, "Long Break")
		} else {
			utils.ShowNotification("Pomodoro Timer", "Time for a break!")
			runPomodoro(timer, status, progress, dataBox, breakDuration, "Break")
		}
		utils.ShowNotification("Pomodoro Timer", "Time to get back to work!")
		cycleCount++

	}
}

func runPomodoro(timer *widgets.Paragraph, status *widgets.Paragraph, progress *widgets.Gauge, dataBox *widgets.Paragraph, duration time.Duration, statusText string) {
	startTime := time.Now()
	endTime := startTime.Add(duration)
	ticker := time.NewTicker(time.Second)

	for range ticker.C {
		remaining := endTime.Sub(time.Now())
		if remaining <= 0 {
			break
		}
		// Modify the status text to include the timer display string
		status.Text = fmt.Sprintf("%s - %s", statusText, utils.FormatTime(remaining))
		progress.Percent = int((duration - remaining) * 100 / duration)
		ui.Render(timer, status, progress, dataBox)
	}

	// Show completed message

	status.Text = fmt.Sprintf("%s - Completed", statusText)
	progress.Percent = 100
	ui.Render(timer, status, progress, dataBox)

	// Wait for a second to show the completed message
	time.Sleep(1 * time.Second)
}
