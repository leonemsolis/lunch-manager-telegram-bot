package telegram

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
)

var menuAuth = &tb.ReplyMarkup{ResizeReplyKeyboard: true, OneTimeKeyboard: true}
var btnAuth = menuAuth.Text("Войти 🔒")

var menuAdmin = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
var btnNewMenu = menuAdmin.Text("Создать новое меню 🍲")
var btnShowResults = menuAdmin.Text("Посмотреть результаты 👀")
var btnTimer = menuAdmin.Text("Настройки уведомленияw ⏱")
var btnLogout = menuAdmin.Text("Выйти 🚪")

var menuEditor = &tb.ReplyMarkup{ResizeReplyKeyboard: true}
var btnAdd = menuEditor.Text("Добавить вариант 🍜")
var btnTest = menuEditor.Text("Посмотреть опрос 👀")
var btnPost = menuEditor.Text("Отправить опрос ✅")
var btnClear = menuEditor.Text("Отчистить все 🗑")

type Bot struct {
	Bot      *tb.Bot
	adminKey string
	admin    *tb.User
	chat     *tb.Chat

	total_voters []string
	voted_today []string

	checkHour int
	checkMinute int

	currentPoolID string

	currentMenu *Menu

	message tb.Editable
}

func NewBot(bot *tb.Bot, adminKey string) *Bot {
	return &Bot{Bot: bot, adminKey: adminKey, checkHour: 9, checkMinute: 20}
}

func (b *Bot) addNewVoter(voter string) {
	if !isIn(b.total_voters, voter) {
		b.total_voters = append(b.total_voters, voter)
	}
}

func (b *Bot) markVoter(voter string, voted bool) {
	if voted {
		if !isIn(b.voted_today, voter) {
			b.voted_today = append(b.voted_today, voter)
		}
	} else {
		for index, name := range b.voted_today {
			if name == voter {
				b.voted_today = remove(b.voted_today, index)
				return
			}
		}
	}
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func isIn(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func (b *Bot) timeChecker() {
	notificationSent := false
	for true {
		location, _ := time.LoadLocation("UTC")
		hour := time.Now().In(location).Add(time.Hour * 6).Hour()
		minute := time.Now().In(location).Add(time.Hour * 6).Minute()
		if hour == b.checkHour - 4 && !notificationSent {
			notificationSent = true
			b.sendNotification()
		}
		if hour == b.checkHour && minute == b.checkMinute {
			b.sendResults()
			return
		}
		time.Sleep(time.Minute)
	}
}

func (b *Bot) sendNotification() {
	if b.chat != nil {
		b.Bot.Send(b.chat, "Кандидаты на голодовку: \n"+b.getNonVoted())
	}
}

func (b *Bot) getNonVoted() string {
	var nonVoted []string
	for _, t := range b.total_voters {
		voted := false
		for _, v := range b.voted_today {
			if t == v {
				voted = true
				break
			}
		}
		if !voted {
			nonVoted = append(nonVoted, t)
		}
	}
	result := ""
	for _, n := range nonVoted {
		result += "@"+n+" "
	}
	return result
}

func (b *Bot) sendResults() {
	if b.admin != nil {
		b.Bot.Send(b.admin, "Время опроса вышло, вот результаты: \n"+b.currentMenu.GetResults())
	}
}

func (b *Bot) Init() {
	b.initMenus()

	b.setHandlers()
	b.setAuthHandlers()
	b.setAdminHandlers()
	b.setEditorHandlers()
}


func (b *Bot) initMenus() {
	menuAuth.Reply(
		menuAuth.Row(btnAuth),
	)

	menuAdmin.Reply(
		menuAdmin.Row(btnNewMenu),
		menuAdmin.Row(btnShowResults),
		menuAdmin.Row(btnTimer),
		menuAdmin.Row(btnLogout),
	)

	menuEditor.Reply(
		menuEditor.Row(btnAdd, btnTest),
		menuEditor.Row(btnClear, btnPost),
	)
}

func (b *Bot) authorizedAction(user *tb.User, callback func()) {
	if b.admin != nil && user.ID == b.admin.ID {
		callback()
	} else {
		b.Bot.Send(user, "Ты не админ 😡")
	}
}

func (b *Bot) setDefaultEmptyTextHandler() {
	b.Bot.Handle(tb.OnText, func(m *tb.Message) {
		b.Bot.Send(m.Sender, "Я не знаю такой команды 🥲")
	})
}
