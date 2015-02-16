package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Client struct {
	Id string
	Context *gin.Context
	Time int32
}

func (client *Client) New(context *gin.Context) Client {
	client.Id = random()
	client.Context = context
	client.Time = int32(time.Now().Unix())
	return *Client
}

func (client *Client) {
}