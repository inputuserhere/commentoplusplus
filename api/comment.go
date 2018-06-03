package main

import (
	"time"
)

type comment struct {
	CommentHex    string    `json:"commentHex"`
	Domain        string    `json:"domain,omitempty"`
	Path          string    `json:"url,omitempty"`
	CommenterHex  string    `json:"commenterHex"`
	Markdown      string    `json:"markdown"`
	Html          string    `json:"html"`
	ParentHex     string    `json:"parentHex"`
	Score         int       `json:"score"`
	State         string    `json:"-"`
	CreationDate  time.Time `json:"creationDate"`
	VoteDirection int       `json:"voteDirection"`
}