package disbot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

type InteractionHandler func(s *discordgo.Session, i *discordgo.InteractionCreate) error

type Bot struct {
	session  *discordgo.Session
	token    string
	commands []*ApplicationCommand
	// onError ハンドラーでエラーが発生したときに呼ばれる。Bot ユーザーにエラーを通知するために i を使用可能
	onError                   func(s *discordgo.Session, i *discordgo.InteractionCreate, err error)
	componentReactionHandlers map[string]InteractionHandler
	registeredCommands        []*discordgo.ApplicationCommand
}

type CommandDetail struct {
	DescriptionEn string
	DescriptionJa string
	Options       []*discordgo.ApplicationCommandOption
}

type ApplicationCommand struct {
	Name                     string
	Detail                   *CommandDetail
	GuildId                  string
	Handler                  InteractionHandler
	DefaultMemberPermissions *int64
	DefaultPermission        *bool
}

func New(token string, commands []*ApplicationCommand, messageActionMap map[string]InteractionHandler, onError func(s *discordgo.Session, i *discordgo.InteractionCreate, err error), onReady func(s *discordgo.Session, r *discordgo.Ready)) (*Bot, error) {
	_s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}

	b := &Bot{
		token:                     token,
		commands:                  commands,
		componentReactionHandlers: messageActionMap,
		onError:                   onError,
		session:                   _s,
	}

	b.session.AddHandler(b.onInteractionCreate)
	b.session.AddHandler(onReady)

	return b, nil
}

func (b *Bot) Start() error {
	if b.session == nil {
		return fmt.Errorf("session is not initialized")
	}

	err := b.session.Open()
	if err != nil {
		return fmt.Errorf("failed to open session: %w", err)
	}

	b.registeredCommands = make([]*discordgo.ApplicationCommand, len(b.commands))
	for i, v := range b.commands {
		b.registeredCommands[i] = v.Command()
	}
	b.registeredCommands, err = b.session.ApplicationCommandBulkOverwrite(b.session.State.User.ID, "", b.registeredCommands)

	return err
}

func (b *Bot) Close() error {
	_commands, err := b.session.ApplicationCommandBulkOverwrite(b.session.State.User.ID, "", []*discordgo.ApplicationCommand{})
	if err != nil {
		return fmt.Errorf("cannot delete all commands: %w", err)
	}

	b.registeredCommands = _commands
	b.commands = nil

	if err := b.session.Close(); err != nil {
		return fmt.Errorf("cannot close Discord connection: %w", err)
	}

	return nil
}

func (b *Bot) Session() *discordgo.Session {
	return b.session
}

func (b *Bot) appId() string {
	if b.session == nil {
		return ""
	}
	return b.session.State.User.ID
}

func (b *Bot) RegisterCommand(guildId string, c ...*ApplicationCommand) error {
	for _, cmd := range c {
		cmd.GuildId = guildId
	}
	b.commands = append(b.commands, c...)

	if b.session == nil {
		return fmt.Errorf("session is not initialized")
	}

	for _, cmd := range c {
		r, err := b.session.ApplicationCommandCreate(b.appId(), guildId, cmd.Command())
		if err != nil {
			return fmt.Errorf("failed to register command: %w", err)
		}

		b.registeredCommands = append(b.registeredCommands, r)
	}

	return nil
}
