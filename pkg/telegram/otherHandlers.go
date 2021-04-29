package telegram

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

func (b *Bot) setHandlers() {
	b.Bot.Handle(tb.OnAddedToGroup, func(m *tb.Message) {
		if b.chat == nil {
			b.Bot.Send(m.Chat, "Привет всем! Я Обед Менеджер v2.0")
			b.chat = m.Chat
		} else {
			b.Bot.Send(m.Chat, "Извините, у меня уже есть группа 😰")
		}
	})

	b.Bot.Handle(tb.OnPollAnswer, func(pa *tb.PollAnswer) {
		if b.currentPoolID != pa.PollID {
			return
		}

		b.addNewVoter(pa.User.Username)



		if len(pa.Options) == 0 {
			// revoke vote
			b.markVoter(pa.User.Username, false)
			b.currentMenu.RemoveVote(pa.User.ID)
			return
		}

		// Correct!
		b.currentMenu.AddVote(pa.User.ID, pa.Options)
		b.markVoter(pa.User.Username, true)

	})

	b.setDefaultEmptyTextHandler()
}