package GoBlox

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func Config(key string) App {
	return App{Key: key}
}

func CheckForPremium(Data map[string]any) (Return []PremiumGroups) {
	if info, ok := Data["resolved"].(map[string]any)["roblox"]; ok {
		for _, value := range info.(map[string]interface{})["groupsv2"].(map[string]interface{}) {
			G := value.(map[string]interface{})["group"].(map[string]interface{})
			R := value.(map[string]interface{})["role"].(map[string]interface{})
			Data := PremiumGroups{
				Group: Group{
					ID:               int(G["id"].(float64)),
					Name:             G["name"].(string),
					MemberCount:      int(G["memberCount"].(float64)),
					HasVerifiedBadge: G["hasVerifiedBadge"].(bool),
				},
				Role: Role{
					ID:   int(R["id"].(float64)),
					Name: R["name"].(string),
					Rank: int(R["rank"].(float64)),
				},
			}
			Return = append(Return, Data)
		}
	}
	return
}

func (k *App) ServerDiscordToRoblox(DiscordID, GuildID string) (data Roblox, err Error) {
	if resp, e := http.DefaultClient.Do(genHttp(must(http.NewRequest("GET", fmt.Sprintf(ServerDiscordToRoblox, GuildID, DiscordID), nil)), []headers{{Name: "Authorization", Value: k.Key}})); e == nil {
		res, _ := io.ReadAll(resp.Body)
		switch resp.StatusCode {
		case 200:
			json.Unmarshal(res, &data)

			var Check map[string]any = make(map[string]any)
			json.Unmarshal(res, &Check)
			data.Resolved.Roblox.Groupsv2 = CheckForPremium(Check)

			return
		case 404:
			json.Unmarshal(res, &err)
			return
		default:
			return data, Error{Error: fmt.Sprintf("Unknown status code: %v | %v", resp.StatusCode, string(res))}
		}
	} else {
		return data, Error{Error: e.Error()}
	}
}

func (k *App) ServerRobloxToDiscord(DiscordID, GuildID string) (data Discord, err Error) {
	if resp, e := http.DefaultClient.Do(genHttp(must(http.NewRequest("GET", fmt.Sprintf(ServerRobloxToDiscord, GuildID, DiscordID), nil)), []headers{{Name: "Authorization", Value: k.Key}})); e == nil {
		res, _ := io.ReadAll(resp.Body)
		switch resp.StatusCode {
		case 200:
			json.Unmarshal(res, &data)
			return
		case 404:
			json.Unmarshal(res, &err)
			return
		default:
			return data, Error{Error: fmt.Sprintf("Unknown status code: %v | %v", resp.StatusCode, string(res))}
		}
	} else {
		return data, Error{Error: e.Error()}
	}
}

func (k *App) ServerGroups(DiscordID, GuildID string) (Data Groups, err Error) {
	if resp, e := http.DefaultClient.Do(genHttp(must(http.NewRequest("GET", fmt.Sprintf(ServerGroups, DiscordID, GuildID), nil)), []headers{{Name: "Authorization", Value: k.Key}})); e == nil {
		res, _ := io.ReadAll(resp.Body)
		switch resp.StatusCode {
		case 200:
			json.Unmarshal(res, &Data)
			return
		case 404:
			json.Unmarshal(res, &err)
			return
		default:
			return Data, Error{Error: fmt.Sprintf("Unknown status code: %v | %v", resp.StatusCode, string(res))}
		}
	} else {
		return Data, Error{Error: e.Error()}
	}
}

func (k *App) GlobalDiscordToRoblox(DiscordID string) (data Roblox, err Error) {
	if resp, e := http.DefaultClient.Do(genHttp(must(http.NewRequest("GET", fmt.Sprintf(GlobalDiscordToRoblox, DiscordID), nil)), []headers{{Name: "Authorization", Value: k.Key}})); e == nil {
		res, _ := io.ReadAll(resp.Body)
		switch resp.StatusCode {
		case 200:
			json.Unmarshal(res, &data)

			var Check map[string]any = make(map[string]any)
			json.Unmarshal(res, &Check)
			data.Resolved.Roblox.Groupsv2 = CheckForPremium(Check)

			return
		case 404:
			json.Unmarshal(res, &err)
			return
		default:
			return data, Error{Error: fmt.Sprintf("Unknown status code: %v | %v", resp.StatusCode, string(res))}
		}
	} else {
		return data, Error{Error: e.Error()}
	}
}

func (k *App) GlobalRobloxToDiscord(RobloxID string) (data Discord, err error) {
	if resp, err := http.DefaultClient.Do(genHttp(must(http.NewRequest("GET", fmt.Sprintf(GlobalRobloxToDiscord, RobloxID), nil)), []headers{{Name: "Authorization", Value: k.Key}})); err == nil {
		res, _ := io.ReadAll(resp.Body)
		switch resp.StatusCode {
		case 200:
			json.Unmarshal(res, &data)
			return data, err
		default:
			return data, errors.New(fmt.Sprintf("Unknown status code: %v | %v", resp.StatusCode, string(res)))
		}
	} else {
		return data, err
	}
}

func genHttp(req *http.Request, headers []headers) *http.Request {
	for _, h := range headers {
		req.Header.Add(h.Name, h.Value)
	}
	return req
}

func must(req *http.Request, err error) *http.Request {
	return req
}
