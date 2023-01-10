package wotoValues

import (
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

var Core *WotoCore = &WotoCore{}

// handlers that need to be shared globally between all plugins
var (
	// CreatePostHandler is the handler responsible for creating a post,
	// set in `path://src/plugins/postsPlugin/helpers.go`.
	CreatePostHandler handlers.Response
	// HelpHandler is the handler responsible for sending the help message
	// to the user.
	// set in `path://src/plugins/helpPlugin/helpers.go`.
	HelpHandler handlers.Response
)
