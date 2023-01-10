package wotoConfig

import (
	"time"

	"github.com/AnimeKaizoku/ssg/ssg"
)

func ParseConfig(configFile string) (*AnonimasuRobotConfig, error) {
	if ConfigSettings != nil {
		return ConfigSettings, nil
	}

	s := &AnonimasuRobotConfig{}

	err := ssg.ParseConfig(s, configFile)
	if err != nil {
		return nil, err
	}

	ConfigSettings = s

	return ConfigSettings, nil
}

func LoadConfig() (*AnonimasuRobotConfig, error) {
	return ParseConfig("config.ini")
}

func GetCmdPrefixes() []rune {
	return []rune{'/', '!'}
}

func GetBotToken() string {
	if ConfigSettings != nil {
		return ConfigSettings.BotToken
	}
	return ""
}

func GetTargetChat() int64 {
	if ConfigSettings != nil {
		return ConfigSettings.TargetChat
	}
	return 0
}

func GetExpirationDays() time.Duration {
	if ConfigSettings != nil {
		return time.Duration(ConfigSettings.BlockExpirationDays) * 24 * time.Hour
	}

	// default is 60 days
	return 60 * 24 * time.Hour
}

func GetWorkingChatId() int64 {
	if ConfigSettings == nil {
		return 0
	}
	return ConfigSettings.TargetChat
}
