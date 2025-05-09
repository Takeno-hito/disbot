package disbot

import (
	"github.com/bwmarrin/discordgo"
)

func ReplyEphemeral(s *discordgo.Session, i *discordgo.InteractionCreate, m string) error {
	return s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: m,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func SendMessage(s *discordgo.Session, channelID string, msg string) error {
	_, err := s.ChannelMessageSend(channelID, msg)
	return err
}

func SendMessageComplex(s *discordgo.Session, channelID string, msg *discordgo.MessageSend) error {
	_, err := s.ChannelMessageSendComplex(channelID, msg)
	return err
}
