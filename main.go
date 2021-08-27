package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("DATA"),
		tgbotapi.NewKeyboardButton("IMAGES"),
		tgbotapi.NewKeyboardButton("TELNUMBER"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("PASSPORT"),
		tgbotapi.NewKeyboardButton("DIPLOM"),
		tgbotapi.NewKeyboardButton("HOMENUMBER"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1922492591:AAGmrQOmxHPA_OaM9wSSmdN2nQdlKd_RyEc")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch update.Message.Text {
		case "/start":
			msg.ReplyMarkup = numericKeyboard
		case "close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}

		switch update.Message.Text {
		case "DATA":
			msgg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello my name is Eldor. I am 18 age. I study is Tashkent arxitektony inistitut")
			bot.Send(msgg)
		case "IMAGES":
			file := "./eldorchik.jpg"
			imsg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, file)
			bot.Send(imsg)
		case "TELNUMBER":
			tel := tgbotapi.NewMessage(update.Message.Chat.ID, "+998973122902")
			bot.Send(tel)
		case "PASSPORT":
			tel := tgbotapi.NewMessage(update.Message.Chat.ID, "Nima deganiz u internetga quymaganim bitta passport qoldi uzi")
			bot.Send(tel)
		case "DIPLOM":
			tel := tgbotapi.NewMessage(update.Message.Chat.ID, "Nima deganiz u hali olmadimku")
			bot.Send(tel)
		case "HOMENUMBER":
			tel := tgbotapi.NewMessage(update.Message.Chat.ID, "Nima deganiz u endi birkami sizga uyimni raqamini berganim qolgandi")
			bot.Send(tel)
		}

	}
}
