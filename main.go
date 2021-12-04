package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-toast/toast"
)

func parseDuration(duration string) int {
	re := regexp.MustCompile(`\d+[smh]`)
	found := re.FindAll([]byte(duration), -1)

	total := 0
	for _, match := range found {
		str := string(match)

		if strings.Contains(str, "h") {
			hours, err := strconv.Atoi(strings.TrimRight(str, "h"))
			checkError(err)

			total += hours * 3600
		} else if strings.Contains(str, "m") {
			minutes, err := strconv.Atoi(strings.TrimRight(str, "m"))
			checkError(err)

			total += minutes * 60
		} else if strings.Contains(str, "s") {
			seconds, err := strconv.Atoi(strings.TrimRight(str, "s"))
			checkError(err)

			total += seconds
		}
	}

	return total
}

func usage() {
	fmt.Println("Usage: remindme about \"<message>\" in <duration>")
}

func main() {
	args := os.Args[1:]

	if len(args) != 4 || args[0] != "about" || args[2] != "in" {
		usage()
		return
	}

	message := args[1]
	duration := parseDuration(args[3])

	notification := toast.Notification{
		AppID:    "{1AC14E77-02E7-4E5D-B744-2EB1AE5198B7}\\WindowsPowerShell\\v1.0\\powershell.exe",
		Title:    "Reminder",
		Message:  message,
		Duration: toast.Long,
	}

	remaining := duration
	for range time.Tick(1 * time.Second) {
		remaining -= 1
		if remaining <= 0 {
			err := notification.Push()
			checkError(err)

			time.Sleep(1 * time.Second)
			return
		}
	}
}
