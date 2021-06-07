package telegram

import (
	"strconv"

	tb "gopkg.in/tucnak/telebot.v2"
)

func (b *Bot) setAdminHandlers() {
	b.Bot.Handle(&btnNewMenu, func(m *tb.Message) {
		b.authorizedAction(m.Sender, func() {
			b.draftMenu = CreateNewMenu()
			b.Bot.Send(m.Sender, "Создайте новое меню", menuEditor)
		})
	})

	b.Bot.Handle(&btnTimer, func(m *tb.Message) {
		b.authorizedAction(m.Sender, func() {
			b.Bot.Send(m.Sender, "Я буду присылать тебе результаты в "+strconv.Itoa(b.checkHour)+":"+strconv.Itoa(b.checkMinute)+"\nПредупреждать не проголосовавших буду за 12 часов до результатов\nДля выхода напиши \"отмена\"\nДля настройки времени напиши час и минуту в формате \"01:23\"")
			b.Bot.Handle(tb.OnText, func(m *tb.Message) {
				if m.Text == "отмена" {
					b.Bot.Send(m.Sender, "Выход в главное меню", menuAdmin)
					b.setDefaultEmptyTextHandler()
					return
				} else {
					if len(m.Text) == 5 {
						hourSub := string([]rune(m.Text)[:2])
						hour, err := strconv.Atoi(hourSub)

						minuteSub := string([]rune(m.Text)[3:])
						minute, err := strconv.Atoi(minuteSub)

						if err == nil && hour >= 0 && hour <= 24 && minute >= 0 && minute <= 60 {
							b.checkHour = hour
							b.checkMinute = minute
							b.Bot.Send(m.Sender, "Принято! Результаты последующих опросов будут отсылаться в "+m.Text, menuAdmin)
							b.setDefaultEmptyTextHandler()
							return
						}
					}
				}
				b.Bot.Send(m.Sender, "Неверный формат. Отмена!", menuAdmin)
				b.setDefaultEmptyTextHandler()
			})
		})
	})

	b.Bot.Handle(&btnShowResults, func(m *tb.Message) {
		b.authorizedAction(m.Sender, func() {
			b.Bot.Send(m.Sender, b.currentMenu.GetResults())
		})
	})

	b.Bot.Handle(&btnRemoveVoter, func(m *tb.Message) {
		b.authorizedAction(m.Sender, func() {
			b.Bot.Send(m.Sender, "Напишите никнейм голосующего, которого вы хотите удалить из уведомлений (формат: @someuser)")
			b.Bot.Handle(tb.OnText, func(m *tb.Message) {
				ok := b.removeVoter(m.Text)
				if ok {
					b.Bot.Send(m.Sender, "Пользователь удален из рассылки уведомлений", menuAdmin)
				} else {
					b.Bot.Send(m.Sender, "Ошибка, пользователь не найден", menuAdmin)
				}

				b.setDefaultEmptyTextHandler()
			})
		})
	})

	b.Bot.Handle(&btnListVoters, func(m *tb.Message) {
		b.authorizedAction(m.Sender, func() {
			b.Bot.Send(m.Sender, b.getAllVoters())
		})
	})

	b.Bot.Handle(&btnLogout, func(m *tb.Message) {
		b.Bot.Send(m.Sender, "Пока!", menuAuth)
		b.admin = nil
	})
}
