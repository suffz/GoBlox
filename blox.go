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

func (k *App) ServerDiscordToRoblox(DiscordID, GuildID string) (data Roblox, err Error) {
	if resp, e := http.DefaultClient.Do(genHttp(must(http.NewRequest("GET", fmt.Sprintf(ServerDiscordToRoblox, GuildID, DiscordID), nil)), []headers{{Name: "Authorization", Value: k.Key}})); e == nil {
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
