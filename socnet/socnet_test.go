package socnet

import (
	"testing"
)


func TestReadTimeline(t *testing.T) {
	t.Run("should return a message published in a personnal timeline", func(t *testing.T)  {
		message := []string{"I love the weather today"}
		timeline := &Timeline{"Alice", message}
		socnet := SocialNetwork{[]Timeline{*timeline}}

		expected := []string{"I love the weather today"}
		got := socnet.ReadTimeline("Alice")
		if got[0] != expected[0] {
			t.Errorf("got %q expected %q", got, expected)
		}
	})

	t.Run("should return a list of messages published in a personnal timeline", func(t *testing.T) {
		message := []string{"Damn! We lost!", "Good game though."}
		timeline := &Timeline{"Bob", message}
		socnet := SocialNetwork{[]Timeline{*timeline}}

		expected := []string{"Good game though.", "Damn! We lost!"}
		got := socnet.ReadTimeline("Bob")

		if got[0] != expected[0] {
			t.Errorf("got %q expected %q", got, expected)
		}
	})
}