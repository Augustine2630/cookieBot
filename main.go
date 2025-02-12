package main

import (
	"cookieBot/bootstrap"
	"cookieBot/usecase"
	"encoding/json"
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"os"
)

func main() {

	app := bootstrap.NewApp()

	bot, err := telego.NewBot(app.Config.App.Token, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	defer bot.StopLongPolling()

	recipeUsecase := usecase.NewRecipeUsecase()

	for update := range updates {
		recipes, err := recipeUsecase.GetRecipesByType(update.Message.Text, app.Config)
		if err != nil {
			fmt.Printf(err.Error())
		}

		marshal, err := json.Marshal(recipes)
		if err != nil {
			return
		}
		sentMessage, _ := bot.SendMessage(
			tu.Message(
				tu.ID(update.Message.Chat.ID),
				string(marshal),
			),
		)
		fmt.Printf("Sent Message: %v\n", sentMessage)
	}
}
