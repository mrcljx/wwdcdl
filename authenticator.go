package main

import (
	"net/http"
	"encoding/json"
	"io"
	"io/ioutil"
	"fmt"
	"net/url"
	"net/http/cookiejar"
	"os/exec"
	"os"
	"time"
	"flag"
	"github.com/howeyc/gopass"
)

var username, password, teamId string
var askPassword bool

func init() {
  flag.StringVar(&username, "username", "", "AppleID username")
  flag.StringVar(&teamId, "team", "", "Apple Developer Team ID")
  flag.BoolVar(&askPassword, "password", false, "Ask for AppleID password")
}

type Authenticator struct {
	cookies []*http.Cookie
	authenticated bool
}

func NewAuthenticator() *Authenticator {
	if askPassword {
		fmt.Printf("Password: ")
		password = string(gopass.GetPasswd())
		askPassword = false
	}

	return &Authenticator{
		cookies: make([]*http.Cookie, 0),
		authenticated: false,
	}
}

func (a *Authenticator) IsAuthenticated() bool {
	return a.authenticated 
}

func (a *Authenticator) Authenticate() (err error) {
	if a.authenticated {
		return
	}
	
	a.authenticated = true
	
	err = a.loadCookiesViaCasper()
	return
}

func (a *Authenticator) loadCookiesViaCasper() (err error) {
	casper, err := exec.LookPath("casperjs")
	
	if err != nil {
		return
	}
	
	file, err := ioutil.TempFile("", "wwdcdl")
	
	if err != nil {
		return
	}
	
	cmd := exec.Command(casper, "login.coffee", file.Name(), username, password, teamId)
	
	stdout, err := cmd.StdoutPipe()
	
	if err != nil {
		return
	}
	
	go io.Copy(os.Stdout, stdout)
	
	stderr, err := cmd.StderrPipe()
	
	if err != nil {
		return
	}
	
	go io.Copy(os.Stderr, stderr)

	stdin, err := cmd.StdinPipe()
	
	if err != nil {
		return
	}
	
	go io.Copy(stdin, os.Stdin)
	
	err = cmd.Run()
	
	if err != nil {
		return
	}
	
	fmt.Printf(output)
	
	a.loadCookiesFromFile(file.Name())
	
	return
}

func (a *Authenticator) loadCookiesFromFile(fileName string) (err error) {
	data, err := ioutil.ReadFile(fileName)
	a.loadCookies(data)
	return
}

func (a *Authenticator) loadCookies(data []byte) {
	var rawCookies []map[string]interface{}
	
	json.Unmarshal(data, &rawCookies)
	
	for _, rawCookie := range rawCookies {
		cookie := &http.Cookie{
			Name: rawCookie["name"].(string),
			Value: rawCookie["value"].(string),
			Path: rawCookie["path"].(string),
			Domain: rawCookie["domain"].(string),
			Expires: time.Unix(int64(rawCookie["expiry"].(float64)), 0),
			RawExpires: rawCookie["expires"].(string),
			MaxAge: 0,
			Secure: rawCookie["secure"].(bool),
			HttpOnly: rawCookie["httponly"].(bool),
		}
		
		a.cookies = append(a.cookies, cookie) 
	}
	
	cookieUrl, _ := url.Parse("https://apple.com")
	http.DefaultClient.Jar, _ = cookiejar.New(nil)
	http.DefaultClient.Jar.SetCookies(cookieUrl, a.cookies)
}

func (a *Authenticator) Extend(req *http.Request) (err error) {
	err = a.Authenticate()
	return
}
