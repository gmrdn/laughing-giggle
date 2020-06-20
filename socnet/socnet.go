package socnet

import (
	"fmt"
	"time"
)

// SocialNetwork is a collection of timelines
type SocialNetwork struct {
	Timelines []Timeline
}

// Timeline is a collection of messages for a given user name
type Timeline struct {
	Name     string
	Messages []Message
}

// Message struct is a text published at a given date
type Message struct {
	Text string
	Date time.Time
}

// ReadTimeline returns the list of messages for a given user name
func (socnet SocialNetwork) ReadTimeline(name string) []string {
	for i := 0; i < len(socnet.Timelines); i++ {
		if socnet.Timelines[i].Name == name {
			return reverseMessageLog(socnet.Timelines[i].Messages)
		}
	}
	return []string{}
}

func reverseMessageLog(messages []Message) []string {
	var reversedList []string
	for i := len(messages); i > 0; i-- {
		reversedList = append(reversedList, messages[i - 1].Text + " (" + getTimeSinceMessage(messages[i - 1]) + ")")
	}
	return reversedList
}

func getTimeSinceMessage(message Message) string {
	duration := time.Now().Sub(message.Date)
	
	formatted := formatDurationWithUnit(duration)
	return formatted
}

func formatDurationWithUnit(duration time.Duration) string {
	if time.Duration.Seconds(duration) < 1.0 {
		return "just now"
	}
	if time.Duration.Seconds(duration) <= 59 {
		return fmt.Sprintf("%d seconds ago", int64(duration.Seconds())) 
	}
	return fmt.Sprintf("%d minutes ago", int64(duration.Minutes())) 
}