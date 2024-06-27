package voice

import (
	"github.com/bwmarrin/discordgo"
)

// Connect the bot to the specified voice channel.
func ConnectToVoiceChannel(s *discordgo.Session, guildID, channelID string) (*discordgo.VoiceConnection, error) {
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return nil, err
	}
	return vc, nil
}

// Disconnect the bot from the voice channel.
func DisconnectFromVoiceChannel(vc *discordgo.VoiceConnection) error {
	err := vc.Disconnect()
	if err != nil {
		return err
	}
	return nil
}
