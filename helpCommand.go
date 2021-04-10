package commander

import (
	"github.com/bwmarrin/discordgo"
	embed "github.com/clinet/discordgo-embed"
	"strconv"
)

// The internal help command
// of the command handler
func helpCommand(s *discordgo.Session, m *discordgo.MessageCreate, c *Commander) {
	emb := embed.NewEmbed()
	emb.Title = "Help"
	emb.Color = 0xADE81C
	emb.Description = "Help command"
	channel, _ := s.Channel(c.CommandChannel)
	emb.AddField("Prefix:", c.Prefix)
	emb.AddField("Command channel:", channel.Name)
	for _, v := range c.Commands {
		emb.AddField(v.Command, "description: "+v.Description+"\npermission: "+strconv.Itoa(v.Permission))
	}
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, emb.MessageEmbed)
	if err != nil {
		panic(err.Error())
	}
}
