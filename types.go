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
	Added    []string `json:"addedRoles"`
	Removed  []string `json:"removedRoles"`
	Nickname string   `json:"nickname"`
}

// Premium

type Resolved struct {
	Roblox  PremiumRoblox  `json:"roblox"`
	Discord PremiumDiscord `json:"discord"`
}
type Badges struct {
	ImageURI string `json:"ImageUri"`
	Name     string `json:"Name"`
}
type Group struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	MemberCount      int    `json:"memberCount"`
	HasVerifiedBadge bool   `json:"hasVerifiedBadge"`
}
type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Rank int    `json:"rank"`
}
type PremiumGroups struct {
	Group Group `json:"group"`
	Role  Role  `json:"role"`
}
type Avatar struct {
	BustThumbnail     string `json:"bustThumbnail"`
	HeadshotThumbnail string `json:"headshotThumbnail"`
	FullBody          string `json:"fullBody"`
}

type PremiumRoblox struct {
	Name                   string          `json:"name"`
	ID                     int             `json:"id"`
	DisplayName            string          `json:"displayName"`
	Description            string          `json:"description"`
	IsBanned               bool            `json:"isBanned"`
	Created                time.Time       `json:"created"`
	Badges                 []Badges        `json:"badges"`
	ProfileLink            string          `json:"profileLink"`
	Presence               interface{}     `json:"presence"`
	Groups                 []PremiumGroups `json:"groups"`
	Avatar                 Avatar          `json:"avatar"`
	Rap                    interface{}     `json:"rap"`
	Value                  interface{}     `json:"value"`
	PlaceVisits            interface{}     `json:"placeVisits"`
	HasDisplayName         bool            `json:"hasDisplayName"`
	ExternalAppDisplayName interface{}     `json:"externalAppDisplayName"`
	HasVerifiedBadge       bool            `json:"hasVerifiedBadge"`
	Groupsv2               []PremiumGroups `json:"groupsv2"`
}

// discord

type DiscordUser struct {
	ID               string      `json:"id"`
	Username         string      `json:"username"`
	Avatar           string      `json:"avatar"`
	Discriminator    string      `json:"discriminator"`
	PublicFlags      int         `json:"public_flags"`
	Flags            int         `json:"flags"`
	Banner           string      `json:"banner"`
	AccentColor      interface{} `json:"accent_color"`
	GlobalName       string      `json:"global_name"`
	AvatarDecoration interface{} `json:"avatar_decoration"`
	DisplayName      string      `json:"display_name"`
	BannerColor      interface{} `json:"banner_color"`
}
type PremiumDiscord struct {
	Avatar                     interface{} `json:"avatar"`
	CommunicationDisabledUntil interface{} `json:"communication_disabled_until"`
	Flags                      int         `json:"flags"`
	JoinedAt                   time.Time   `json:"joined_at"`
	Nick                       interface{} `json:"nick"`
	Pending                    bool        `json:"pending"`
	PremiumSince               interface{} `json:"premium_since"`
	Roles                      []string    `json:"roles"`
	User                       DiscordUser `json:"user"`
	Mute                       bool        `json:"mute"`
	Deaf                       bool        `json:"deaf"`
}
