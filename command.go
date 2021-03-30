package commander

import "github.com/bwmarrin/discordgo"

// The command struct
// it defines the data to register a
// command
type Command struct {
	Command     string
	Description string
	Executor    func(s *discordgo.Session, m *discordgo.MessageCreate)
	Permission  int
}
