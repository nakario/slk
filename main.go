package main

import (
	"flag"
	"log"

	"github.com/nlopes/slack"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var mes = flag.String("m", "", "message to post")

func main() {
	flag.Parse()

	viper.SetDefault("prefix", "")
	viper.SetDefault("channel", "")
	viper.SetDefault("token", "")
	viper.SetConfigName(".slk")
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalln("Failed to detect the home directory:", err)
	}
	viper.AddConfigPath(home)
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Failed to read the config file:", err)
	}
	prefix := viper.GetString("prefix")
	channel := viper.GetString("channel")
	token := viper.GetString("token")

	cli := slack.New(token)
	params := slack.NewPostMessageParameters()
	params.EscapeText = false
	_, _, err = cli.PostMessage(channel, prefix + *mes, params)
	if err != nil {
		log.Fatalln("Failed to post the message:", err)
	}
}
