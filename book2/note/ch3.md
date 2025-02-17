第三章 Go模块

```go
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
```

逻辑为
1. 构建URL参数，包含API和我们想要搜索的电影名
2. 组合我们想要查询的REST API URL，以及步骤（1）的参数
3. 向站点发出请求，如果有错误，返回nil和error
4. 如果没有错误，将响应的body传递给json.Unmarshal，将其解码为MovieInfo结构体

