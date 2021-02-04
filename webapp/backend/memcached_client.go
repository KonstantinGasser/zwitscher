package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
)

type memClient struct {
	client *memcache.Client
}

func (c *memClient) dashboard(userID string) (*ResponseDashboard, error) {
	item, err := c.client.Get(userID)
	if err != nil {
		log.Printf("MEM: %v\n", err)
		return &ResponseDashboard{
			Msg:    "memcache client failed to get(id) or had a get_miss...",
			Status: 500,
			Data:   nil,
		}, err
	}
	return c.tweetList(item)
}

func (c *memClient) tweetList(item *memcache.Item) (*ResponseDashboard, error) {

	// get data from item: in string form sep=#@#@#

	stringTweets := strings.Split(string(item.Value), "#@#@#")
	stringTweets = stringTweets[1:]
	resp := &ResponseDashboard{
		Msg:    "here you go",
		Status: 200,
		Data:   []itemTweet{},
	}
	for _, tweet := range stringTweets {
		var iTweet itemTweet
		err := json.Unmarshal([]byte(tweet), &iTweet)
		if err != nil {
			continue
		}
		resp.Data = append(resp.Data, iTweet)
	}
	return resp, nil
}

// type itemTweet struct {
// 	Author      string `json:"author"`
// 	NumOfLikes  string `json:"number_of_likes"`
// 	NumOfShares string `json:"number_of_shares"`
// 	Content     string `json:"content"`
// 	UserID      string `json:"user_id"`
// }
