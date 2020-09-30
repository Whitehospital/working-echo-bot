package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	version      = "5.124"
	reqURL       = "https://api.vk.com/method/wall.get?"
	vkServiceKey = "f1a1e5f7f1a1e5f7f1a1e5f73bf1d5cce4ff1a1f1a1e5f7aed2e257e7c31c0116dbe3e1"
)

type wallResponse struct {
	Body body `json:"response"`
}

type body struct {
	Items []Items `json:"items"`
}

type Items struct {
	Text string `json:"text"`
}

func getPosts(groupId string) ([]Items, error) {
	u := url.Values{}
	u.Set("count", "200")
	u.Set("offset", "0")
	u.Set("filter", "owner")
	u.Set("owner_id", groupId)
	u.Set("access_token", vkServiceKey)
	u.Set("v", version)

	req := reqURL + u.Encode()
	resp, err := http.Get(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := new(wallResponse)
	if err := json.Unmarshal(b, response); err != nil {
		return nil, err
	}

	return response.Body.Items, nil

}
