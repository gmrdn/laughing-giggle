package socnet

// SocialNetwork is a collection of timelines
type SocialNetwork struct {
	Timelines		[]Timeline
} 

// Timeline is a collection of messages for a given user name
type Timeline struct {
	Name 		string
	Messages 	[]string
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

func reverseMessageLog(messages []string) []string  {
	var reversedList []string

	for i := len(messages); i > 0; i-- {
		reversedList = append(reversedList, messages[i - 1])
	}	
	return reversedList
}
