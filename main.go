package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type TimerData struct {
	EndTime time.Time `json:"end_time"`
}

func timerFilePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".sleep_champ_timer.json")
}

func saveTimer(delayMinutes int) error {
	data := TimerData{EndTime: time.Now().Add(time.Duration(delayMinutes) * time.Minute)}
	bytes, _ := json.Marshal(data)
	return os.WriteFile(timerFilePath(), bytes, 0644)
}

func loadTimer() (*TimerData, error) {
	bytes, err := os.ReadFile(timerFilePath())
	if err != nil {
		return nil, err
	}
	var data TimerData
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func deleteTimer() error {
	return os.Remove(timerFilePath())
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected 'start', 'status' or 'delete' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "start":
		startCmd := flag.NewFlagSet("start", flag.ExitOnError)
		delay := startCmd.Int("delay", 60, "Delay in minutes before shutdown")
		startCmd.Parse(os.Args[2:])

		if *delay <= 0 {
			fmt.Println("Please provide a delay > 0 minutes")
			os.Exit(1)
		}

		if err := saveTimer(*delay); err != nil {
			fmt.Println("Error saving timer:", err)
			os.Exit(1)
		}

		fmt.Printf("Timer set for %d minutes. Shutdown scheduled at %s\n", *delay, time.Now().Add(time.Duration(*delay)*time.Minute).Format(time.RFC1123))

	case "status":
		data, err := loadTimer()
		if err != nil {
			fmt.Println("No active timer found.")
			return
		}
		remaining := time.Until(data.EndTime)
		if remaining <= 0 {
			fmt.Println("Timer expired.")
			return
		}
		fmt.Printf("Shutdown scheduled at %s (%v remaining)\n", data.EndTime.Format(time.RFC1123), remaining.Round(time.Second))

	case "delete":
		err := deleteTimer()
		if err != nil {
			fmt.Println("No active timer to delete.")
			return
		}
		fmt.Println("Timer deleted.")

	default:
		fmt.Println("expected 'start', 'status' or 'delete' subcommands")
		os.Exit(1)
	}
}
