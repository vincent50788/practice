package fileserver

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"

	"gitlab.geax.io/zqb/apis/service/internal/app/logger"
)

var (
	errFileNotFound = errors.New("file not found")

	client *SquirrelClient
)

const (
	loginURI string = "/v1/user/login"
	fileURI  string = "/v1/fs"
)

type loginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type response struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type FileInfo struct {
	Name    string    `json:"name"`
	Path    string    `json:"path"`
	Type    string    `json:"type"`
	Size    int64     `json:"size"`
	ModTime time.Time `json:"mtime"`
}

type SquirrelClient struct {
	Host   string
	user   string
	passwd string
	token  string
}

func New(cfg Config) (FileServer, error) {
	client = &SquirrelClient{
		Host:   cfg.Host,
		user:   cfg.Username,
		passwd: cfg.Password,
	}
	err := client.Login()
	if err != nil {
		logger.ApLog().Error(err)
		return nil, err
	}

	return client, nil
}

// Login ...
func (c *SquirrelClient) Login() (err error) {

	httpClient := &http.Client{}

	loginReq := loginRequest{
		Name:     c.user,
		Password: c.passwd,
	}

	loginReqByte, err := json.Marshal(loginReq)
	if err != nil {
		logger.ApLog().Error(err)
		return
	}

	url := c.Host + loginURI
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(loginReqByte))
	if err != nil {
		logger.ApLog().Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		logger.ApLog().Error(err)
		return
	}

	var body []byte

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.ApLog().Error(err)
		return
	}

	data := &token{}

	err = json.Unmarshal(body, data)
	if err != nil {
		logger.ApLog().Error(err)
		return
	}

	c.token = data.AccessToken

	return
}

// Upload ...
func (c *SquirrelClient) Upload(path string, fileName string, f io.Reader) (err error) {
	httpClient := &http.Client{}

	url := c.Host + fileURI + path

	fields := map[string]string{
		"filename": fileName,
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", fields["filename"])
	if err != nil {
		logger.ApLog().Error(err)
		return
	}

	_, err = io.Copy(fw, f)
	if err != nil {
		logger.ApLog().Error(err)
		return
	}

	for k, v := range fields {
		_ = writer.WriteField(k, v)
	}

	err = writer.Close()
	if err != nil {
		logger.ApLog().Error(err)
		return
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		logger.ApLog().Error(err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := httpClient.Do(req)
	if err != nil {
		logger.ApLog().Error(err)
		return
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.ApLog().Error(err)
		return
	}

	data := &response{}

	err = json.Unmarshal(respBody, data)
	if err != nil {
		logger.ApLog().Error(err)
		return
	}

	return
}

// Delete ...
func (c *SquirrelClient) Delete(path string) (ok bool, err error) {
	httpClient := &http.Client{}

	url := c.Host + fileURI + path
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		logger.ApLog().Error(err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := httpClient.Do(req)
	if err != nil {
		logger.ApLog().Error(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.ApLog().Error(err)
		return
	}

	data := &response{}
	err = json.Unmarshal(body, data)
	if err != nil {
		logger.ApLog().Error(err)
		return
	}
	ok = true

	return
}

// GetFileList ...
func (c *SquirrelClient) GetFileList(path string) (file FileInfo, err error) {

	httpClient := &http.Client{}

	url := c.Host + fileURI + path + "?op=info"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		logger.ApLog().Error(err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+c.token)

	resp, err := httpClient.Do(req)
	if err != nil {
		logger.ApLog().Error(err)
		return
	}

	var body []byte

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.ApLog().Error(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("error code=%v", resp.StatusCode))
		return
	}

	data := &response{}
	err = json.Unmarshal(body, data)
	if err != nil {
		logger.ApLog().Error(err)
		return
	}

	if data.Status == 0 {
		err = json.Unmarshal(data.Data, &file)
		if err != nil {
			logger.ApLog().Error(err)
			return
		}
		return
	}
	err = errFileNotFound

	return
}
