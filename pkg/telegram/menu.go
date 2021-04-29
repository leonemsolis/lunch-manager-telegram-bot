package telegram

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
)

type Menu struct {
	title string
	items []*MenuItem
}

type MenuItem struct {
	name string
	votedUsersIDs []int
}

func CreateNewMenu() *Menu {
	return &Menu{title:""}
}

func (m *Menu) AddNewElement(name string) {
	m.items = append(m.items, &MenuItem{
		name:          name,
		votedUsersIDs: []int{},
	})
}

func (mi *MenuItem) GetVotersCount() int {
	return len(mi.votedUsersIDs)
}

func (m *Menu) AddVote(voterID int, selectedIndices []int) {
	for _, index := range selectedIndices {
		m.items[index].votedUsersIDs = append(m.items[index].votedUsersIDs, voterID)
	}
}


func (m *Menu) RemoveVote(voterID int) {
	for _, item := range m.items {
		indexInSlice := -1
		for index := range item.votedUsersIDs {
			if item.votedUsersIDs[index] == voterID {
				indexInSlice = index
				break
			}
		}
		if indexInSlice != -1 {
			item.votedUsersIDs = append(item.votedUsersIDs[:indexInSlice], item.votedUsersIDs[indexInSlice+1:]...)
		}
	}
}

func (m* Menu) GetResults() string {
	if m == nil {
		return "Меню пока что не создано"
	}
	result := "🔴 "+m.title + "\n"
	for _, mi := range m.items {
		result += fmt.Sprintf("🔵🔵 %s - %d\n", mi.name, mi.GetVotersCount())
	}
	return result
}

func (m *Menu) CreatePoll() tb.Poll {
	return tb.Poll{
		ID:              "",
		Type:            tb.PollRegular,
		Question:        m.title,
		Options:         m.createOptions(),
		VoterCount:      0,
		Closed:          false,
		CorrectOption:   0,
		MultipleAnswers: false,
		Explanation:     "",
		ParseMode:       "",
		Entities:        nil,
		Anonymous:       false,
		OpenPeriod:      0,
		CloseUnixdate:   0,
	}
}

func (m *Menu) createOptions() []tb.PollOption {
	var result []tb.PollOption
	for _, item := range m.items {
		result = append(result, tb.PollOption{
			Text: item.name,
			VoterCount: 0,
		})
	}
	return result
}

