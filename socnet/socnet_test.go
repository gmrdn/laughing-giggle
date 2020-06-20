package socnet

import (
	"testing"
	"time"
)

func TestReadTimeline(t *testing.T) {
	t.Run("should return a message published in a personnal timeline", func(t *testing.T) {
		message := []Message{{Text: "I love the weather today", Date: time.Now().Add(- time.Minute * 5)}}
		timeline := &Timeline{"Alice", message}
		socnet := SocialNetwork{[]Timeline{*timeline}}

		expected := []string{"I love the weather today (5 minutes ago)"}
		got := socnet.ReadTimeline("Alice")
		if got[0] != expected[0] {
			t.Errorf("got %q expected %q", got, expected)
		}
	})

	t.Run("should return a list of messages published in a personnal timeline", func(t *testing.T) {
		message := []Message{{Text: "Damn! We lost!", Date: time.Now().Add(- time.Minute * 2)}, {Text: "Good game though.", Date: time.Now().Add(- time.Minute * 1)}}
		timeline := &Timeline{"Bob", message}
		socnet := SocialNetwork{[]Timeline{*timeline}}

		expected := []string{"Good game though. (1 minutes ago)", "Damn! We lost! (2 minutes ago)"}
		got := socnet.ReadTimeline("Bob")

		if got[0] != expected[0] {
			t.Errorf("got %q expected %q", got, expected)
		}
	})
}

