package handlers

import (
	"dripcord/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

type Bot struct {
	Config  *config.Config
	Session *discordgo.Session
}

func NewBot(cfg *config.Config) (*Bot, error) {
	s, err := discordgo.New("Bot " + cfg.DiscordBotToken)
	if err != nil {
		return nil, err
	}

	bot := &Bot{Session: s, Config: cfg}
	s.AddHandler(bot.SendMessage)
	s.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsDirectMessages

	return bot, nil
}

func (b *Bot) SendMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		fmt.Println(m.Author.Username + ": " + m.Content)
		return
	}

	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Error getting channel info:", err)
		return
	}

	if channel.Type == discordgo.ChannelTypeDM || m.ChannelID == b.Config.DiscordChannel {

		fmt.Println(m.Author.Username + ": " + m.Content)

		translatedText, err := Drip(m.Content, b.Config.GeminiToken)
		if err != nil {
			fmt.Println("failed to process message:", err)
		} else {
			s.ChannelMessageSend(m.ChannelID, translatedText)
		}
	}
}

func (b *Bot) SendDM(userID, message string) error {
	channel, err := b.Session.UserChannelCreate(userID)
	if err != nil {
		return fmt.Errorf("failed to create DM channel: %v", err)
	}

	_, err = b.Session.ChannelMessageSend(channel.ID, message)
	if err != nil {
		return fmt.Errorf("failed to send DM: %v", err)
	}

	return nil
}

func (b *Bot) Run() {
	err := b.Session.Open()

	if err != nil {
		log.Fatal("Error opening connection:", err)
	}

	fmt.Println("Bot is running...")

	select {} // runs the bot indefinitely
}
