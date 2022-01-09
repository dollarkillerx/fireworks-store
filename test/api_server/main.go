package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dollarkillerx/urllib"
)

type Login struct {
	ErrMsg string `json:"errMsg"`
	Code   string `json:"code"`
}

func main() {
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("login in...")
		defer request.Body.Close()
		all, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var login Login
		err = json.Unmarshal(all, &login)
		if err != nil {
			log.Fatalln(err)
		}

		appid := "appid"
		secret := "secret"
		code, resp, err := urllib.Get(fmt.Sprintf("https://applets.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appid, secret, login.Code)).Debug().
			ByteOriginal()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(code)
		fmt.Println(string(resp))
	})

	http.ListenAndServe("0.0.0.0:8087", nil)
}
