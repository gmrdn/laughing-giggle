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
		messages := []Message{{Text: "I love the weather today", Date: time.Now().Add(-time.Minute * 5)}}
		timeline := Timeline{"Alice", messages}
		socnet := SocialNetwork{[]Timeline{timeline}}

		expected := []string{"I love the weather today (5 minutes ago)"}
		got := socnet.ReadTimeline("Alice")
		assertCorrectMessage(t, got[0], expected[0])

	})

	t.Run("should return a list of messages published in a personnal timeline", func(t *testing.T) {
		messages := []Message{{Text: "Damn! We lost!", Date: time.Now().Add(-time.Minute * 2)}, {Text: "Good game though.", Date: time.Now().Add(-time.Minute * 1)}}
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

	postedText := "I love the weather today"

	createPostAtTime := func(time time.Time) SocialNetwork {
		messages := []Message{{Text: postedText, Date: time}}
		timeline := Timeline{"Alice", messages}
		socnet := SocialNetwork{[]Timeline{timeline}}
		return socnet
	}

	t.Run("should return just now when posted since less than 1 second", func(t *testing.T) {
		socnet := createPostAtTime(time.Now())
		expected := []string{postedText + " (just now)"}
		got := socnet.ReadTimeline("Alice")
		assertCorrectMessage(t, got[0], expected[0])

	})

	t.Run("should return X seconds ago when posted since 1 second", func(t *testing.T) {
		socnet := createPostAtTime(time.Now().Add(-time.Second * 1))
		expected := []string{postedText + " (1 second ago)"}
		got := socnet.ReadTimeline("Alice")
		assertCorrectMessage(t, got[0], expected[0])
	})

	t.Run("should return X seconds ago when posted since less than 1 minute", func(t *testing.T) {
		socnet := createPostAtTime(time.Now().Add(-time.Second * 42))
		expected := []string{postedText + " (42 seconds ago)"}
		got := socnet.ReadTimeline("Alice")
		assertCorrectMessage(t, got[0], expected[0])
	})

	t.Run("should return 1 minute ago when posted since 1 minute", func(t *testing.T) {
		socnet := createPostAtTime(time.Now().Add(-time.Minute * 1))
		expected := []string{postedText + " (1 minute ago)"}
		got := socnet.ReadTimeline("Alice")
		assertCorrectMessage(t, got[0], expected[0])
	})

	t.Run("should return X minutes ago when posted since less than 1 hour", func(t *testing.T) {
		socnet := createPostAtTime(time.Now().Add(-time.Minute * 42))
		expected := []string{postedText + " (42 minutes ago)"}
		got := socnet.ReadTimeline("Alice")
		assertCorrectMessage(t, got[0], expected[0])
	})

	t.Run("should return 1 hour ago when posted since 1 hour", func(t *testing.T) {
		socnet := createPostAtTime(time.Now().Add(-time.Hour * 1))
		expected := []string{postedText + " (1 hour ago)"}
		got := socnet.ReadTimeline("Alice")
		assertCorrectMessage(t, got[0], expected[0])
	})

	t.Run("should return X hours ago when posted since less than 1 day", func(t *testing.T) {
		socnet := createPostAtTime(time.Now().Add(-time.Hour * 2))
		expected := []string{postedText + " (2 hours ago)"}
		got := socnet.ReadTimeline("Alice")
		assertCorrectMessage(t, got[0], expected[0])
	})

	t.Run("should return X days ago when posted since more than 48 hours", func(t *testing.T) {
		socnet := createPostAtTime(time.Now().Add(-time.Hour * 50))
		expected := []string{postedText + " (2 days ago)"}
		got := socnet.ReadTimeline("Alice")
		assertCorrectMessage(t, got[0], expected[0])
	})

	t.Run("should return X days ago when posted since more than 48 hours", func(t *testing.T) {
		socnet := createPostAtTime(time.Now().Add(-time.Hour * 50))
		expected := []string{postedText + " (2 days ago)"}
		got := socnet.ReadTimeline("Alice")
		assertCorrectMessage(t, got[0], expected[0])
	})

	t.Run("should return 1 day ago when posted since more than 24 hours but less than 48 hours", func(t *testing.T) {
		socnet := createPostAtTime(time.Now().Add(-time.Hour * 30))
		expected := []string{postedText + " (1 day ago)"}
		got := socnet.ReadTimeline("Alice")
		assertCorrectMessage(t, got[0], expected[0])
	})

}
