package main

import (
	"fmt"
	"github.com/tucnak/telebot"
	"log"
	//"os"
	"time"
)

func main() {
	const (
                BOT_TOKEN = "TO:KEN"
	)
	bot, err := telebot.NewBot(BOT_TOKEN)
	if err != nil {
		log.Println(err)
		fmt.Println(err)
	}

	messages := make(chan telebot.Message)
	bot.Listen(messages, 1*time.Second)
	// bisa dikembangkan querry ke database, dll
	for message := range messages {
		if message.Text == "/help" {
			bot.SendMessage(message.Chat, "Hello, "+message.Sender.FirstName+" "+message.Sender.LastName+" ! \nPerintah yang tersedia: /help, /penduduk, /ipm, /ahh, /poverty, /pdrb, /lpe, /tpt \n(nb: dalam huruf kecil)", nil)
		} else if message.Text == "/penduduk" {
			bot.SendMessage(message.Chat, "Jumlah penduduk Gianyar tahun 2015 adalah 495100 jiwa", nil)
		} else if message.Text == "/ipm" {
			bot.SendMessage(message.Chat, "IPM Gianyar\n2010:71,45\n2011:72,5\n2012:73,36\n2013:74,00\n2014:74,29\n2015:75,03", nil)
		} else if message.Text == "/ahh" {
			bot.SendMessage(message.Chat, "AHH Gianyar\n2010:72,31\n2011:72,43\n2012:72,57\n2013:72,71\n2014:72.78\n2015:72,84", nil)
		} else if message.Text == "/poverty" {
			bot.SendMessage(message.Chat, "Persentase Penduduk Kemiskinan Gianyar:\n2010:6,68\n2011:5,40\n2012:4,69\n2013:4,27\n2014:4,57\n2015:4,34**\n Angka tahun 2015 masih sangat sementara", nil)
		} else if message.Text == "/pdrb" {
			bot.SendMessage(message.Chat, "PDRB ADHB Gianyar 2015 : 20.052.659,96 juta\nPDRB ADHK Gianyar 2015 : 15.173.314,94 juta ", nil)
		} else if message.Text == "/lpe" {
			bot.SendMessage(message.Chat, "Laju Pertumbuhan Ekonomi Gianyar\n2011 : 7,15\n2012 : 7,08\n2013 : 6,82\n2014 : 6,79\n2015 : 6,34", nil)
		} else if message.Text == "/tpt" {
			bot.SendMessage(message.Chat, "Tingkat Pengangguran Gianyar:\n2010 : 2,36\n2011 : 2,11\n2012 : 1,83\n2013 : 2,23\n2014 : 1,43\n2015 : 1,93", nil)
		} else if message.Text == "pulang" {
			bot.SendMessage(message.Chat, "waktu nya pulang", nil)
		}

	}

}
