package switchbot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	ApiUrl = "https://api.switch-bot.com/v1.1"
)

type Client struct {
	secret     string
	token      string
	httpClient http.Client
	debug      bool
}

type CommonResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func NewClient(secret string, token string, options ...Option) *Client {
	client := &Client{
		secret:     secret,
		token:      token,
		httpClient: http.Client{},
	}

	for _, opt := range options {
		opt(client)
	}

	return client
}

type Option func(*Client)

func OptionDebug(debugFlag bool) func(*Client) {
	return func(client *Client) {
		client.debug = debugFlag
	}
}

type ResponseParser func(bodyBytes []byte) error

func (client *Client) GetRequest(path string, parser ResponseParser) error {
	nonce := uuid.NewString()
	timestamp := time.Now().UnixMilli()

	data := fmt.Sprintf("%s%d%s", client.token, timestamp, nonce)

	mac := hmac.New(sha256.New, []byte(client.secret))
	if _, err := mac.Write([]byte(data)); err != nil {
		return err
	}

	signature := mac.Sum(nil)
	signatureB64 := strings.ToUpper(base64.StdEncoding.EncodeToString(signature))

	url := fmt.Sprintf("%s%s", ApiUrl, path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", client.token)
	req.Header.Set("sign", signatureB64)
	req.Header.Set("nonce", nonce)
	req.Header.Set("t", fmt.Sprintf("%d", timestamp))

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if client.debug {
		log.Printf("Response: %s", string(bodyBytes))
	}

	err = parser(bodyBytes)
	if err != nil {
		return err
	}

	return nil
}
