package socnet

import (
	"testing"
	"time"
)

func TestReadTimeline(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, expected string) {
		t.Helper()
		if got != expected {
			t.Errorf("got %q want %q", got, expected)
		}
	}
	
	t.Run("should return a message published in a personnal timeline", func(t *testing.T) {
		messages := []Message{{Text: "I love the weather today", Date: time.Now().Add(- time.Minute * 5)}}
		timeline := Timeline{"Alice", messages}
		socnet := SocialNetwork{[]Timeline{timeline}}

		expected := []string{"I love the weather today (5 minutes ago)"}
		got := socnet.ReadTimeline("Alice")
		assertCorrectMessage(t, got[0], expected[0])

	})

	t.Run("should return a list of messages published in a personnal timeline", func(t *testing.T) {
		messages := []Message{{Text: "Damn! We lost!", Date: time.Now().Add(- time.Minute * 2)}, {Text: "Good game though.", Date: time.Now().Add(- time.Minute * 1)}}
		timeline := Timeline{"Bob", messages}
		socnet := SocialNetwork{[]Timeline{timeline}}

		expected := []string{"Good game though. (1 minutes ago)", "Damn! We lost! (2 minutes ago)"}
		got := socnet.ReadTimeline("Bob")
		assertCorrectMessage(t, got[0], expected[0])
		assertCorrectMessage(t, got[1], expected[1])
	})
}

func TestPostedDurationAgo(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, expected string) {
		t.Helper()
		if got != expected {
			t.Errorf("got %q want %q", got, expected)
		}
	}
	
	t.Run("should return just now when posted since less than 1 second", func(t *testing.T) {
		messages := []Message{{Text: "I love the weather today", Date: time.Now()}}
		timeline := Timeline{"Alice", messages}
		socnet := SocialNetwork{[]Timeline{timeline}}

		expected := []string{"I love the weather today (just now)"}
		got := socnet.ReadTimeline("Alice")

		assertCorrectMessage(t, got[0], expected[0])

	})

	t.Run("should return X seconds ago when posted since less than 1 minute", func(t *testing.T) {
		messages := []Message{{Text: "I love the weather today", Date: time.Now().Add(- time.Second * 42)}}
		timeline := Timeline{"Alice", messages}
		socnet := SocialNetwork{[]Timeline{timeline}}

		expected := []string{"I love the weather today (42 seconds ago)"}
		got := socnet.ReadTimeline("Alice")

		assertCorrectMessage(t, got[0], expected[0])

	})
}

