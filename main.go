package main

import (
        "fmt"
        "github.com/tucnak/telebot"
        "log"
        "os"
        "time"
)

func main() {
        bot, err := telebot.NewBot(os.Getenv("BOT_TOKEN"))
        if err != nil {
                log.Println(err)
                fmt.Println(err)
        }

        messages := make(chan telebot.Message)
        bot.Listen(messages, 1*time.Second)

        for message := range messages {
                //      bot.sendMessage(message.Chat, "ketik /help untuk perintah yang tersedia", nil)
                if message.Text == "/hi" {
                        bot.SendMessage(message.Chat, "Hello, "+message.Sender.FirstName+" !", nil)
                } else if message.Text == "/penduduk" {
                        bot.SendMessage(message.Chat, "Jumlah penduduk Gianyar tahun 2015 adalah 495100 jiwa", nil)
                } else if message.Text == "/help" {
                        bot.SendMessage(message.Chat, "perintah yang tersedia /hi, /penduduk, dan /help", nil)
                } else {
                        bot.SendMessage(message.Chat, "Maaf perintah belum tersedia", nil)
                }

        }

}
