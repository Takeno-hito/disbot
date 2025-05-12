package disbot

import "github.com/bwmarrin/discordgo"

func (c *ApplicationCommand) Command() *discordgo.ApplicationCommand {
	if c.Detail.DescriptionEn == "" {
		return &discordgo.ApplicationCommand{
			Name:                     c.Name,
			Description:              c.Detail.DescriptionJa,
			Options:                  c.Detail.Options,
			DefaultMemberPermissions: c.DefaultMemberPermissions,
			DefaultPermission:        c.DefaultPermission,
		}
	}
	if c.Detail.DescriptionJa == "" {
		return &discordgo.ApplicationCommand{
			Name:                     c.Name,
			Description:              c.Detail.DescriptionEn,
			Options:                  c.Detail.Options,
			DefaultMemberPermissions: c.DefaultMemberPermissions,
			DefaultPermission:        c.DefaultPermission,
		}
	}

	return &discordgo.ApplicationCommand{
		Name:        c.Name,
		Description: c.Detail.DescriptionEn,
		DescriptionLocalizations: &map[discordgo.Locale]string{
			discordgo.Japanese: c.Detail.DescriptionJa,
		},
		Options:                  c.Detail.Options,
		DefaultMemberPermissions: c.DefaultMemberPermissions,
		DefaultPermission:        c.DefaultPermission,
	}
}
