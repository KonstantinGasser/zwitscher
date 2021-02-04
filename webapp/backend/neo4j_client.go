package main

import (
	"log"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type neoClient struct {
	driver neo4j.Driver
}

func (c *neoClient) getFollows(userID string) (*ResponseFollowFollowers, error) {

	followers, err := c.Follower(userID)
	if err != nil {
		return &ResponseFollowFollowers{
			Msg:    "Failed to get Followers of yours",
			Status: http.StatusInternalServerError,
			Data:   make(map[string]interface{}),
		}, err
	}
	following, err := c.Following(userID)
	if err != nil {
		return &ResponseFollowFollowers{
			Msg:    "Failed to get Followers of yours",
			Status: http.StatusInternalServerError,
			Data:   make(map[string]interface{}),
		}, err
	}

	return &ResponseFollowFollowers{
		Msg:    "here you go",
		Status: http.StatusOK,
		Data: map[string]interface{}{
			"followers": followers,
			"following": following,
		},
	}, nil
}

func (c *neoClient) Follower(userID string) ([]string, error) {
	session := c.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	follower, err := session.Run("MATCH(x:User)-[:FOLLOWS]->(:User {id: $id}) RETURN x.id", map[string]interface{}{
		"id": userID,
	})
	if err != nil {
		return []string{}, err
	}

	var record *neo4j.Record
	var followers []string
	for follower.NextRecord(&record) {
		followers = append(followers, record.Values[0].(string))
	}
	log.Println(followers)
	return followers, nil
}

func (c *neoClient) Following(userID string) ([]string, error) {
	session := c.driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	follower, err := session.Run("MATCH(:User {id: $id})-[:FOLLOWS]->(x:User) RETURN x.id", map[string]interface{}{
		"id": userID,
	})
	if err != nil {
		return []string{}, err
	}

	var record *neo4j.Record
	var followers []string
	for follower.NextRecord(&record) {
		log.Println(record.Values)
		followers = append(followers, record.Values[0].(string))
	}
	log.Println(followers)
	return followers, nil
}

func newDriver() neo4j.Driver {
	d, err := neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "sink", ""))
	// I am sorry but it is 3am so this justifys the code following
	if err != nil {
		panic(err)
	}
	return d
}
