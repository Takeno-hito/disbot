package disbot

import "github.com/bwmarrin/discordgo"

type LocaleFunc func(l discordgo.Locale) string

func Mono(s string) LocaleFunc {
	return func(l discordgo.Locale) string {
		return s
	}
}

func Bilingual(ja string, en string, l discordgo.Locale) string {
	if ja == "" {
		return en
	}
	if en == "" {
		return ja
	}
	if l == discordgo.Japanese {
		return ja
	}
	return en
}
