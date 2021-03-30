package examples

import (
	"fmt"
	"github.com/MathisBurger/commander"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Init discord session
	session, err := discordgo.New("Bot <TOKEN>")
	if err != nil {
		fmt.Println("Error while creating discord session")
		return
	}

	// Init command handler with (prefix, commandChannel)
	handler := commander.New(";;", "826378911444500501")

	// register command (name, description, executor, "permission")
	handler.Register("example", "Example command", ExampleCommand, 100)

	// add command handler to discord session
	// to use the handler in your bot
	session.AddHandler(handler.Execute)

	// Configures the bot to run infinitely
	// and stops if you press CTRL+C
	session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
	err = session.Open()
	if err != nil {
		fmt.Println("Cannot connect to discord websocket")
		return
	}
	fmt.Println("The bot is running now. Terminate by using CTRL-C")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	session.Close()
}

// Example command sending a default message
func ExampleCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, _ = s.ChannelMessageSend(m.ChannelID, "Example message")
}
