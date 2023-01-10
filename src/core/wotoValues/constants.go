package wotoValues

const (
	MaximumHistory           = 3
	MultiDbLength            = 10
	MultiDbLastIndex         = 9
	MultiDbFirstIndex        = 0
	MinPostContentTextLength = 1
)

const (
	ResultIdentifier     = "::"
	AdvancedInlinePrefix = "-wh::"
	AdvancedInlineSuffix = "::"
	HelpDataInline       = "help-inline"
	StartDataCreate      = "create"
)

const (
	PermissionNormalUser UserPermission = iota
	PermissionSpecialUser
	PermissionAdmin
	PermissionDev
	PermissionOwner
)
