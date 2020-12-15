package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"net/url"

// 	"github.com/tarb/util/www"
// )

// func main() {

// 	username := ""
// 	password := ""
// 	appkey := ""

// 	client := www.NewDefault()
// 	client.SetDefaultHeaders(func(h http.Header) {
// 		h.Set("X-Application", appkey)
// 		h.Set("Accept", "application/json")
// 		h.Set("Connection", "keep-alive")
// 	})

// 	bs, err := client.Post("https://identitysso.betfair.com/api/login").
// 		WithFormBody(func(form url.Values) {
// 			form.Set("username", username)
// 			form.Set("password", password)
// 		}).
// 		DoWithRetry(5, www.LinearJitterDelay).
// 		CollectBytes()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(string(bs))
// }
