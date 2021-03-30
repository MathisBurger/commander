package commander

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

/////////////////////////////////////////////////////////////
// Main command handler struct. The whole library          //
// is based on it. It takes all required params            //
// needed to execute a command and handle it               //
/////////////////////////////////////////////////////////////
type Commander struct {
	Prefix         string
	CommandChannel string
	Commands       map[string]Command
}

// creates a new instance of the command handler
// returns the instance which can be used for
// registration and execution
func New(prefix string, cmdChannel string) *Commander {
	return &Commander{
		prefix,
		cmdChannel,
		map[string]Command{},
	}
}

// register a new command
func (c *Commander) Register(
	command string,
	description string,
	executor func(s *discordgo.Session, m *discordgo.MessageCreate),
	permission int,
) {
	c.Commands[command] = Command{command, description, executor, permission}
}

// executes a command with given
// parameters and sends not found,
// if the command was not found
func (c *Commander) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, c.Prefix) {
		if c.CommandChannel == "" || c.CommandChannel == m.ChannelID {
			command := strings.Split(strings.Split(m.Content, ";;")[1], " ")
			if command[0] == "help" {
				helpCommand(s, m, c)
			}
			if val, ok := c.Commands[command[0]]; ok {
				val.Executor(s, m)
			} else {
				sendNotFound(s, m, command[0], c)
			}
		}
	}
}
