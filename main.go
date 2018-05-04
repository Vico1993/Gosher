package main

import (
	"fmt"

	"github.com/kataras/iris"
	pusher "github.com/pusher/pusher-http-go"
	"github.com/spf13/viper"
)

type message struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
}

// Message Container
var messages = make([]map[string]interface{}, 0)

func main() {
	// Initialization of the application ( Pusher Client )

	// Let's configuration about Pusher
	viper.SetConfigName("conf")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	// Config Pusher Application
	client := pusher.Client{
		AppId:   viper.GetString("pusherAppId"),
		Key:     viper.GetString("pusherKey"),
		Secret:  viper.GetString("pusherSecret"),
		Cluster: "us2",
		Secure:  true,
	}

	// Create Iris application
	app := iris.New()

	tmpl := iris.Handlebars("./templates", ".html")
	tmpl.Reload(true)

	// Add this parameter for the web assets file ( like css file / Js file )
	app.RegisterView(tmpl)
	app.StaticWeb("/templates", "./templates")

	// Racine of the application display comments + form to add one
	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("messages", messages)
		ctx.View("index.html")
	})

	// When form submited, add the message to the messages variables.
	app.Post("/message", func(ctx iris.Context) {
		var mess message
		if err := ctx.ReadJSON(&mess); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
			return
		}

		newMessage := map[string]interface{}{
			"id":      mess.ID,
			"message": mess.Body,
		}
		messages = append(messages, newMessage)

		client.Trigger("message", "new_message", newMessage)
	})

	app.Run(iris.Addr("127.0.0.1:8080"))
}
