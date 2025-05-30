package disbot

import (
	"github.com/bwmarrin/discordgo"
)

func (b *Bot) onInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		interactionId := i.ApplicationCommandData().Name
		if c, ok := b.searchCommand(interactionId); ok {
			err := c.Handler(s, i)
			if err != nil {
				b.onError(s, i, err)
			}
		} else {
			b.onError(s, i, ErrUnknownCommandKey)
		}
		return
	case discordgo.InteractionMessageComponent:
		interactionId := i.MessageComponentData().CustomID
		if h, ok := b.componentReactionHandlers[interactionId]; ok {
			err := h(s, i)
			if err != nil {
				b.onError(s, i, err)
			}
		} else {
			b.onError(s, i, ErrUnknownCommandKey)
		}
		return
	case discordgo.InteractionModalSubmit:
		interactionId := i.ModalSubmitData().CustomID
		if h, ok := b.componentReactionHandlers[interactionId]; ok {
			err := h(s, i)
			if err != nil {
				b.onError(s, i, err)
			}
		} else {
			b.onError(s, i, ErrUnknownCommandKey)
		}
		return
	default:
		b.onError(s, i, ErrUndefinedCommandType)
		return
	}
}

func (b *Bot) searchCommand(name string) (command *ApplicationCommand, ok bool) {
	for _, c := range b.commands {
		if c.Name == name {
			return c, true
		}
	}
	return nil, false
}
