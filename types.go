package GoBlox

import "time"

var (
	ServerDiscordToRoblox = "https://api.blox.link/v4/public/guilds/%v/discord-to-roblox/%v"
	ServerRobloxToDiscord = "https://api.blox.link/v4/public/guilds/%v/roblox-to-discord/%v"
	ServerGroups          = "https://api.blox.link/v4/public/guilds/%v/update-user/%v"
)

var (
	GlobalDiscordToRoblox = "https://api.blox.link/v4/public/discord-to-roblox/%v"
	GlobalRobloxToDiscord = "https://api.blox.link/v4/public/roblox-to-discord/%v"
)

type App struct {
	Key string
}

type Discord struct {
	DiscordIDs []string `json:"discordIDs"`
}

type Roblox struct {
	RobloxID string   `json:"robloxID"`
	Resolved Resolved `json:"resolved"`
}
type Resolved struct {
}

type Error struct {
	Error string `json:"error"`
}

type Bloxlink struct {
	Success bool `json:"success"`
	User    User `json:"user"`
}

type User struct {
	RobloxID       string `json:"robloxId"`
	PrimaryAccount string `json:"primaryAccount"`
}

type RobloxData struct {
	Description            string      `json:"description"`
	Created                time.Time   `json:"created"`
	IsBanned               bool        `json:"isBanned"`
	ExternalAppDisplayName interface{} `json:"externalAppDisplayName"`
	HasVerifiedBadge       bool        `json:"hasVerifiedBadge"`
	ID                     int         `json:"id"`
	Name                   string      `json:"name"`
	DisplayName            string      `json:"displayName"`
}

type headers struct {
	Name, Value string
}

type Groups struct {
	Added    []string `json:"added"`
	Removed  []string `json:"removed"`
	Nickname string   `json:"nickname"`
}
