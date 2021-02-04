package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type APIserver struct {
	memcache *memClient
	neo4j    *neoClient
	kafka    *kafkaClient
}

func (srv *APIserver) setUp() {

	http.HandleFunc("/getfollow", srv.getFollowFollow)
	http.HandleFunc("/dash", srv.getDashboard)
	http.HandleFunc("/zwitscherlos", srv.postTweet)
}
func (srv *APIserver) run() error {
	return http.ListenAndServe(":7080", nil)
}

// get follow list from user x
// get followed by list from user x
func (srv *APIserver) getFollowFollow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	userID := r.URL.Query().Get("user_id")
	log.Printf("User id :%s\n", userID)
	if userID == "" {
		http.Error(w, `{"msg": "user_id not found in query", "status": 404}`, http.StatusBadRequest)
		return
	}
	follows, err := srv.neo4j.getFollows(userID)
	if err != nil {
		log.Printf("neo4j: %v\n", err)
	}
	data, err := json.MarshalIndent(follows, "", "  ")
	if err != nil {
		http.Error(w, `{"msg": "failed to marshal response", "status": 500}`, http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// get personalized dashboard tweets from memcache
func (srv *APIserver) getDashboard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	userID := r.URL.Query().Get("user_id")
	log.Printf("User id :%s\n", userID)
	if userID == "" {
		http.Error(w, `{"msg": "user_id not found in query", "status": 404}`, http.StatusBadRequest)
		return
	}

	tweets, err := srv.memcache.dashboard(userID)
	if err != nil {
		log.Printf("memcache: %v\n", err)
	}
	data, err := json.MarshalIndent(tweets, "", "  ")
	if err != nil {
		http.Error(w, `{"msg": "failed to marshal response", "status": 500}`, http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (srv *APIserver) postTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if r.Method == "OPTIONS" {
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil || len(payload) <= 0 {
		http.Error(w, "ups", http.StatusBadRequest)
		return
	}

	srv.kafka.msgChan <- payload
	w.WriteHeader(http.StatusOK)
}

type ResponseDashboard struct {
	Msg    string      `json:"msg"`
	Status int         `json:"status"`
	Data   []itemTweet `json:"content"`
}

type ResponseFollowFollowers struct {
	Msg    string                 `json:"msg"`
	Status int                    `json:"status"`
	Data   map[string]interface{} `json:"content"`
}