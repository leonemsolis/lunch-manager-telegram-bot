package telegram

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

func (b *Bot) setAuthHandlers() {
	b.Bot.Handle("/start", func(m *tb.Message) {
		//b.Bot.Send(m.Sender, "Hello there! Send me your password using '/auth password'")
		b.Bot.Send(m.Sender, "Привет!", menuAuth)
	})

	b.Bot.Handle(&btnMenu, func(m *tb.Message) {
		b.Bot.Send(m.Sender, b.availableMenus)
	})

	b.Bot.Handle(&btnAuth, func(m *tb.Message) {
		if b.admin != nil {
			b.Bot.Send(m.Sender, "Администратор уже зашел")
			return
		}
		b.Bot.Send(m.Sender, "Теперь напиши мне пароль")
		b.Bot.Handle(tb.OnText, func(m *tb.Message) {
			if m.Text == b.adminKey {
				b.admin = m.Sender
				b.Bot.Send(m.Sender, "Привет менеджер!", menuAdmin)
			} else {
				b.Bot.Send(m.Sender, "Неверный пароль, попробуй еще раз", menuAuth)
			}
			b.setDefaultEmptyTextHandler()
		})
	})

	//b.Bot.Handle(&btnCheck, func(m *tb.Message) {
	//	menuSelect := &tb.ReplyMarkup{
	//		InlineKeyboard:      [][]tb.InlineButton{
	//			[]tb.InlineButton{
	//				{
	//					Unique: "HELLOASODJASDAOSDJ",
	//					Text:            "1",
	//					URL:             "",
	//					Data:            "",
	//					InlineQuery:     "",
	//					InlineQueryChat: "",
	//				},
	//			},
	//		},
	//	}
	//
	//
	//	message, _ := b.Bot.Send(b.chat, "Hello world", menuSelect)
	//	b.message = message
	//})
}
