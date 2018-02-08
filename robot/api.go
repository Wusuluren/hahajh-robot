package robot

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"os"
)

type HahajhItem struct {
	Text    string
	Picture string
}

const (
	loginUrl   = "login"
	logoutUrl  = "logout"
	publishUrl = "publish"
)

type Account struct {
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Urls      map[string]string
	cookieJar *cookiejar.Jar
	cookie    *http.Cookie
}

func InitAccount(account *Account) error {
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}
	account.cookieJar = cookieJar
	return nil
}

func (ac *Account) Login() error {
	formData := map[string]string{
		"username": ac.Username,
		"password": ac.Password,
	}
	byte, err := json.Marshal(formData)
	if err != nil {
		return err
	}
	client := &http.Client{
		CheckRedirect: nil,
		Jar:           ac.cookieJar,
	}
	req, err := http.NewRequest("POST", ac.Urls[loginUrl], bytes.NewReader(byte))
	if err != nil {
		return err
	}
	return checkHttpRespError(client.Do(req))
}

func (ac *Account) Logout() error {
	client := &http.Client{
		CheckRedirect: nil,
		Jar:           ac.cookieJar,
	}
	req, err := http.NewRequest("GET", ac.Urls[logoutUrl], nil)
	if err != nil {
		return err
	}
	return checkHttpRespError(client.Do(req))
}

func (ac *Account) Publish(item HahajhItem) error {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	defer w.Close()
	f, err := os.Open(item.Picture)
	if err != nil {
		return err
	}
	defer f.Close()
	fw, err := w.CreateFormFile("pic", item.Picture)
	if err != nil {
		return err
	}
	if _, err = io.Copy(fw, f); err != nil {
		return err
	}
	fw, err = w.CreateFormField("text_area")
	if err != nil {
		return err
	}
	if _, err = fw.Write([]byte(item.Text)); err != nil {
		return err
	}

	req, err := http.NewRequest("POST", ac.Urls[publishUrl], &b)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())
	client := &http.Client{
		CheckRedirect: nil,
		Jar:           ac.cookieJar,
	}
	return checkHttpRespError(client.Do(req))
}
