package github_api

import (
	"net/http"
	"time"
	"fmt"
	"github.com/joho/godotenv"
)

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	err := godotenv.Load(".env")
	if err != nil {
	  return nil, fmt.Errorf("error loading .env file")
	}
	return &Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}, nil
}

const baseURL = "https://api.github.com/"
