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
		// @lebay_bot
		BOT_TOKEN = "295042840:AAG1jtvqBUa9Z-MlhzSV8Wx_gq0YY0RVzfE"
		// @gianyar_bot
		//BOT_TOKEN = "295733510:AAHzYDUqIDpgdiIg0_Re1eqSKd6r8lvKCO8"
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
			bot.SendMessage(message.Chat, "Hello, "+message.Sender.FirstName+" "+message.Sender.LastName+" ! \nPerintah yang tersedia: /help, /penduduk, /ipm, /ahh, /pdrb, /lpe, /tpt\n/gini, /p0, /p1, /p2, /gk\n(nb: dalam huruf kecil)", nil)
		} else if message.Text == "/penduduk" {
			bot.SendMessage(message.Chat, "Jumlah penduduk Gianyar tahun 2015 adalah 495100 jiwa", nil)
		} else if message.Text == "/ipm" {
			bot.SendMessage(message.Chat, "IPM Gianyar\n2010:71,45\n2011:72,5\n2012:73,36\n2013:74,00\n2014:74,29\n2015:75,03", nil)
		} else if message.Text == "/ahh" {
			bot.SendMessage(message.Chat, "AHH Gianyar\n2010:72,31\n2011:72,43\n2012:72,57\n2013:72,71\n2014:72.78\n2015:72,84", nil)
		} else if message.Text == "/p0" {
			bot.SendMessage(message.Chat, "Persentase Penduduk Miskin (P0) Gianyar:\n2010:6,68\n2011:5,40\n2012:4,69\n2013:4,61\n2014:4,57\n2015:4,34**\n Angka tahun 2015 masih sangat sementara", nil)
		} else if message.Text == "/pdrb" {
			bot.SendMessage(message.Chat, "PDRB ADHB Gianyar 2015 : 20.052.659,96 juta\nPDRB ADHK Gianyar 2015 : 15.173.314,94 juta ", nil)
		} else if message.Text == "/lpe" {
			bot.SendMessage(message.Chat, "Laju Pertumbuhan Ekonomi Gianyar\n2011 : 7,15\n2012 : 7,08\n2013 : 6,82\n2014 : 6,79\n2015 : 6,34", nil)
		} else if message.Text == "/tpt" {
			bot.SendMessage(message.Chat, "Tingkat Pengangguran Gianyar:\n2010 : 2,36\n2011 : 2,11\n2012 : 1,83\n2013 : 2,23\n2014 : 1,43\n2015 : 1,93", nil)
		} else if message.Text == "pulang" {
			bot.SendMessage(message.Chat, "waktu nya pulang", nil)
		} else if message.Text == "/gini" {
			bot.SendMessage(message.Chat, "Gini Ratio Gianyar\n2011 : 0,3279\n2012:0,3362\n2013:0,3254\n2014 : 0,3774\n2015 : 0,3249", nil)
		} else if message.Text == "/p1" {
			bot.SendMessage(message.Chat, "Kedalaman Kemiskinan (P1) Gianyar\n2010 : 0,8900\n2011 : 0,4550\n2012 : 0,4700\n2013 : 0,4484\n2014 : 0,5500\n2015 : 0,5635", nil)
		} else if message.Text == "/p2" {
			bot.SendMessage(message.Chat, "Keparahan Kemiskinan (P2) Gianyar\n2010 : 0,1700\n2011 : 0,0586\n2012 : 0,0800\n2013 : 0,0857\n2014 : 0,1331\n2015 : 0,1001", nil)
		} else if message.Text == "/p2" {
			bot.SendMessage(message.Chat, "Garis Kemiskinan (Rp) Gianyar :\n2010 : 237.904\n2011 : 260.704\n2012 : 274.639\n2013 : 279.742\n2014 : 298.465\n2015 : 320.805", nil)
		} else if message.Text == "/tes" {
			bot.SendMessage(message.Chat, "Tes running bot di clever-cloud.com", nil)
		} 

	}

}
