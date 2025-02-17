package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const APIKEY = "193ef3a"

type MovieInfo struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	ImdbRating string `json:"imdbRating"`
	ImdbID     string `json:"imdbID"`
}

func sendGetRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	// 使用http模块中的GET函数运行时实际的GET请求
	// 将响应存储在resp和err中

	defer resp.Body.Close() // 在函数返回之前，确保响应中关闭Body输入流
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return string(body), errors.New(resp.Status)
	}
	return string(body), nil
}

// 200是HTTP成功状态码，表示请求成功

func SearchByName(name string) (*MovieInfo, error) { // 接收参数string，返回的值是json响应
	parms := url.Values{}
	parms.Set("apikey", APIKEY)
	parms.Set("t", name)
	siteURL := "http://www.omdbapi.com/?" + parms.Encode()
	body, err := sendGetRequest(siteURL)
	if err != nil {
		return nil, errors.New(err.Error() + "\nBody:" + body)
	}
	mi := &MovieInfo{}
	return mi, json.Unmarshal([]byte(body), mi)
}

// mi:=&movieInfo{}
// return mi,json.Unmarshal([]byte(body),mi)

func SearchById(id string) (*MovieInfo, error) {
	parms := url.Values{}
	parms.Set("apikey", APIKEY)
	parms.Set("i", id)
	siteURL := "http://www.omdbapi.com/?" + parms.Encode()
	body, err := sendGetRequest(siteURL)
	if err != nil {
		return nil, errors.New(err.Error() + "\nBody:" + body)
	}
	mi := &MovieInfo{}
	return mi, json.Unmarshal([]byte(body), mi)
}

func main() {
	body, _ := SearchById("tt3896298")
	fmt.Println(body.Title)
	body, _ = SearchByName("Game of")
	fmt.Println(body.Title)
}
