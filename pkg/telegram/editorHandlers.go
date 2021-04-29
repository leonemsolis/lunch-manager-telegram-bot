package telegram

import tb "gopkg.in/tucnak/telebot.v2"

func (b *Bot) setEditorHandlers() {

	b.Bot.Handle(&btnAdd, func(m *tb.Message) {
		b.authorizedAction(m.Sender, func() {
			b.Bot.Send(m.Sender, "Напиши название варианта ✍🏻")
			b.Bot.Handle(tb.OnText, func(m *tb.Message) {
				b.currentMenu.AddNewElement(m.Text)
				b.Bot.Send(m.Sender, "Принято!", menuEditor)
				b.setDefaultEmptyTextHandler()
			})
		})
	})

	b.Bot.Handle(&btnTitle, func(m *tb.Message) {
		b.authorizedAction(m.Sender, func() {
			b.Bot.Send(m.Sender, "Напиши новый заголовок 📝")
			b.Bot.Handle(tb.OnText, func(m *tb.Message) {
				b.currentMenu.title = m.Text
				b.Bot.Send(m.Sender, "Принято!", menuEditor)
				b.setDefaultEmptyTextHandler()
			})
		})
	})

	b.Bot.Handle(&btnTest, func(m *tb.Message) {
		b.authorizedAction(m.Sender, func() {
			poll := b.currentMenu.CreatePoll()
			poll.Send(b.Bot, m.Sender, &tb.SendOptions{})
		})
	})

	b.Bot.Handle(&btnClear, func(m *tb.Message) {
		b.authorizedAction(m.Sender, func() {
			b.Bot.Send(m.Sender, "Отчистка завершена 💥", menuAdmin)
		})
	})

	b.Bot.Handle(&btnPost, func(m *tb.Message) {
		b.authorizedAction(m.Sender, func() {
			if b.chat == nil {
				b.Bot.Send(m.Sender, "Ошибка. Нужно настройить чат 🥲", menuAdmin)
				return
			}

			if b.currentMenu.title == "" {
				b.Bot.Send(m.Sender, "Добавьте заголовок опроса 😡", menuEditor)
				return
			}

			poll := b.currentMenu.CreatePoll()
			mess, err := poll.Send(b.Bot, b.chat, &tb.SendOptions{})

			if err != nil || mess == nil {
				b.Bot.Send(m.Sender, "Произошла какая-то ошибка 😅", menuAdmin)
				return
			}



			b.currentPoolID = mess.Poll.ID

			// Clear voted_today slice
			b.voted_today = nil
			go b.timeChecker()

			b.UpdateAvailableMenus()

			b.Bot.Send(m.Sender, "Отлично, опрос отправлен 👍", menuAdmin)
		})
	})
}
