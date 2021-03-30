package commander

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	emb "github.com/clinet/discordgo-embed"
	"log"
)

func sendNotFound(s *discordgo.Session, m *discordgo.MessageCreate, cmd string, c *Commander) {
	embed := emb.NewEmbed()
	embed.SetTitle(fmt.Sprintf("Command `%s` not found", cmd))
	embed.AddField("Tips:", fmt.Sprintf("Type `%shelp`", c.Prefix))
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed.MessageEmbed)
	if err != nil {
		log.Fatal("Error while sending not found embed")
	}
}
