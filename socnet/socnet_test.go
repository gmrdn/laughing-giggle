package socnet

import (
	"testing"
)

func TestReadTimeline(t *testing.T) {
	t.Run("should return a message published in a personnal timeline", func(t *testing.T) {
		message := []Message{Message{Text: "I love the weather today", Date: "19-06-2020"}}
		timeline := &Timeline{"Alice", message}
		socnet := SocialNetwork{[]Timeline{*timeline}}

		expected := []Message{Message{Text: "I love the weather today", Date: "19-06-2020"}}
		got := socnet.ReadTimeline("Alice")
		if got[0] != expected[0] {
			t.Errorf("got %q expected %q", got, expected)
		}
	})

	t.Run("should return a list of messages published in a personnal timeline", func(t *testing.T) {
		message := []Message{Message{Text: "Damn! We lost!", Date: "123"},Message{Text: "Good game though.", Date: "toto"}}
		timeline := &Timeline{"Bob", message}
		socnet := SocialNetwork{[]Timeline{*timeline}}

		expected := []Message{Message{Text: "Good game though.", Date: "toto"},Message{Text: "Damn! We lost!", Date: "123"}}
		got := socnet.ReadTimeline("Bob")

		if got[0] != expected[0] {
			t.Errorf("got %q expected %q", got, expected)
		}
	})
}
