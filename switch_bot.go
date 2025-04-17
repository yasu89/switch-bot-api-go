package switchbot

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	DefaultBaseApiURL = "https://api.switch-bot.com/v1.1"
)

type Client struct {
	secret     string
	token      string
	httpClient http.Client
	debug      bool
	baseApiURL string
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
		baseApiURL: DefaultBaseApiURL,
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

func OptionBaseApiURL(baseApiURL string) func(*Client) {
	return func(client *Client) {
		client.baseApiURL = baseApiURL
	}
}

type ResponseParser func(client *Client, bodyBytes []byte) error

func (client *Client) setHeader(req *http.Request) error {
	nonce := uuid.NewString()
	timestamp := time.Now().UnixMilli()

	data := fmt.Sprintf("%s%d%s", client.token, timestamp, nonce)

	mac := hmac.New(sha256.New, []byte(client.secret))
	if _, err := mac.Write([]byte(data)); err != nil {
		return err
	}

	signature := mac.Sum(nil)
	signatureB64 := strings.ToUpper(base64.StdEncoding.EncodeToString(signature))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", client.token)
	req.Header.Set("sign", signatureB64)
	req.Header.Set("nonce", nonce)
	req.Header.Set("t", fmt.Sprintf("%d", timestamp))

	return nil
}

func (client *Client) GetRequest(path string, parser ResponseParser) error {
	url := fmt.Sprintf("%s%s", client.baseApiURL, path)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	err = client.setHeader(req)
	if err != nil {
		return err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if client.debug {
		log.Printf("Response: %s", string(responseBodyBytes))
	}

	err = parser(client, responseBodyBytes)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) PostRequest(path string, request *ControlRequest) (*CommonResponse, error) {
	url := fmt.Sprintf("%s%s", client.baseApiURL, path)
	requestBodyJson, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyJson))
	if err != nil {
		return nil, err
	}

	err = client.setHeader(req)
	if err != nil {
		return nil, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if client.debug {
		log.Printf("Response: %s", string(responseBodyBytes))
	}

	response := &CommonResponse{}
	err = json.Unmarshal(responseBodyBytes, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
