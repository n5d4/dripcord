package handlers

import (
	"dripcord/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
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
	s.Identify.Intents = discordgo.IntentsGuildMessages
	return bot, nil
}

func (b *Bot) SendMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.ChannelID != b.Config.DiscordChannel {
		return
	}

	translatedText := Translate(m.Content, b.Config.GeminiToken)
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(translatedText))
}

func (b *Bot) Run() {
	err := b.Session.Open()

	if err != nil {
		log.Fatal("Error opening connection:", err)
	}

	fmt.Println("Bot is running...")

	go func() {
		ticker := time.NewTicker(10 * time.Minute) // Change interval as needed
		defer ticker.Stop()

		for range ticker.C {
			_, err := b.Session.ChannelMessageSend("233418632388935683", "Scheduled message: I'm still alive!")
			if err != nil {
				log.Println("Error sending scheduled message:", err)
			}
		}
	}()

	select {} // runs the bot indefinitely
}
