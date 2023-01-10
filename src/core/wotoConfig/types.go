package wotoConfig

type AnonimasuRobotConfig struct {
	BotToken            string `section:"main" key:"bot_token"`
	TargetChat          int64  `section:"main" key:"target_chat"`
	BlockExpirationDays int    `section:"main" key:"block_expiration_days"`
}
