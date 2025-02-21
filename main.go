package main

import (
	"cookieBot/bootstrap"
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"math/rand"
	"os"
	"time"
)

func main() {
	app := bootstrap.NewApp()

	os.Setenv("BOT_TOKEN", app.Config.App.Token)

	bot, err := telego.NewBot(app.Config.App.Token, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	defer bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil {
			switch update.Message.Text {
			case "/start":
				// Send a message with a reply keyboard to roll a D20
				_, _ = bot.SendMessage(
					tu.Message(
						tu.ID(update.Message.Chat.ID),
						"Press the button to roll a D20!",
					).WithReplyMarkup(
						tu.Keyboard(
							tu.KeyboardRow(
								tu.KeyboardButton("Roll D20"),
								tu.KeyboardButton("Roll D100"),
							),
						).WithResizeKeyboard(), // Resize the keyboard to fit the button
					),
				)
			case "Roll D20":
				// Handle the button press to roll a D20
				roll := RollD20(20)
				_, _ = bot.SendMessage(
					tu.Message(
						tu.ID(update.Message.Chat.ID),
						fmt.Sprintf("Your roll is: %d", roll),
					),
				)
			case "Roll D100":
				// Handle the button press to roll a D20
				roll := RollD20(100)
				_, _ = bot.SendMessage(
					tu.Message(
						tu.ID(update.Message.Chat.ID),
						fmt.Sprintf("Your roll is: %d", roll),
					),
				)
			default:
				// Handle other messages
				_, _ = bot.SendMessage(
					tu.Message(
						tu.ID(update.Message.Chat.ID),
						"Press the 'Roll D20' button to roll!",
					),
				)
			}
		}
	}
}

func RollD20(rollSize int) int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(rollSize) + 1
}
