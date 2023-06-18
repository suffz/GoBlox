package GoBlox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetRoblox(RobloxID string) (data RobloxData) {
	if resp, err := http.DefaultClient.Do(genHttp(must(http.NewRequest("GET", "https://users.roblox.com/v1/users/"+RobloxID, nil)), []headers{})); err == nil {
		if resp.StatusCode == 200 {
			body, _ := io.ReadAll(resp.Body)
			json.Unmarshal(body, &data)
		}
	}
	return
}

func FindRoblox(RobloxName string) (Data, error) {
	if resp, err := http.Post("https://users.roblox.com/v1/usernames/users", "application/json", bytes.NewBuffer(Struct2Bytes(RobloxInfo{
		Usernames:          []string{RobloxName},
		ExcludeBannedUsers: true,
	}))); err == nil {
		bo, _ := io.ReadAll(resp.Body)
		switch resp.StatusCode {
		case 200:
			var D RI
			json.Unmarshal(bo, &D)
			for _, names := range D.Data {
				if strings.EqualFold(names.Name, RobloxName) {
					return names, nil
				}
			}
		default:
			return Data{}, errors.New(fmt.Sprintf("Unknown status code: %v | %v", resp.StatusCode, string(bo)))
		}
	} else {
		return Data{}, err
	}
	return Data{}, errors.New("Unable to get profile")
}

func Struct2Bytes(L any) []byte {
	Body, _ := json.Marshal(L)
	return Body
}

type RobloxInfo struct {
	Usernames          []string `json:"usernames"`
	ExcludeBannedUsers bool     `json:"excludeBannedUsers"`
}

type RI struct {
	Data []Data `json:"data"`
}
type Data struct {
	RequestedUsername string `json:"requestedUsername"`
	HasVerifiedBadge  bool   `json:"hasVerifiedBadge"`
	ID                int    `json:"id"`
	Name              string `json:"name"`
	DisplayName       string `json:"displayName"`
}
