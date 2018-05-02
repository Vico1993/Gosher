package main

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/spf13/viper"
)

func main() {

	// TEST DATA
	messages := make([]map[string]string, 0)

	message1 := map[string]string{
		"id":      "1",
		"message": "bonjour",
	}

	message2 := map[string]string{
		"id":      "2",
		"message": "AuRevoir",
	}

	messages = append(messages, message1)
	messages = append(messages, message2)

	// Let's configuration about Pusher
	viper.SetConfigName("conf")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	// client := pusher.Client{
	// 	AppId:   viper.GetString("pusherAppId"),
	// 	Key:     viper.GetString("pusherKey"),
	// 	Secret:  viper.GetString("pusherSecret"),
	// 	Cluster: "us2",
	// 	Secure:  true,
	// }

	app := iris.New()

	tmpl := iris.Handlebars("./templates", ".html")
	tmpl.Reload(true)

	app.RegisterView(tmpl)
	app.StaticWeb("/templates", "./templates")

	app.Get("/", func(ctx iris.Context) {
		fmt.Println("New Connection")

		ctx.ViewData("messages", messages)
		ctx.View("index.html")
	})

	app.Run(iris.Addr("127.0.0.1:8080"))
}
