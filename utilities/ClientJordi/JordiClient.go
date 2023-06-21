package ClientJordi

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ClientJordi struct {
	BaseUrl string
	client  *http.Client
}

func New(url string) ClientJordi {
	return ClientJordi{
		BaseUrl: url,
		client:  &http.Client{},
	}
}

func (c ClientJordi) CreateNewUrlWithParams(param []string) string {
	c.BaseUrl = fmt.Sprintf("%s%s", c.BaseUrl, param[0])
	NewUrl, err := url.Parse(c.BaseUrl)
	if err != nil {
		fmt.Println("Errore durante il parsing dell'URL di base:", err)
		panic(err)
	}
	params := url.Values{}
	for i := 1; i < len(param); i += 2 {
		params.Set(param[i], param[i+1])
	}

	NewUrl.RawQuery = params.Encode()

	fullUrl := NewUrl.String()
	return fullUrl
}

func (c ClientJordi) SetUrl(NewUrl string) ClientJordi {
	c.BaseUrl = NewUrl
	return c
}

func (c ClientJordi) Get(path string) *http.Response {
	NewUrl := fmt.Sprintf("%s%s", c.BaseUrl, path)
	resp, err := http.Get(NewUrl)
	if err != nil {
		panic(err)
	}
	return resp
}

func (c ClientJordi) GetWithVariables(path string, variables []string) *http.Response {
	NewUrl := fmt.Sprintf("%s%s", c.BaseUrl, path)
	for i := range variables {
		NewUrl = fmt.Sprintf("%s/%s", NewUrl, variables[i])
	}
	resp, err := http.Get(NewUrl)

	if err != nil {
		panic(err)
	}
	return resp
}

func (c ClientJordi) Add(path string, body string) *http.Response {
	NewUrl := fmt.Sprintf("%s%s", c.BaseUrl, path)
	resp, err := http.Post(NewUrl, "application/json", strings.NewReader(body))
	if err != nil {
		panic(err)
	}
	return resp
}

func (c ClientJordi) Delete(path string, name string, surname string) *http.Response {
	NewUrl := fmt.Sprintf("%s%s/%s/%s", c.BaseUrl, path, name, surname)

	req, err := http.NewRequest("DELETE", NewUrl, nil)
	if err != nil {
		panic(err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		panic(err)
	}
	return resp
}

func (c ClientJordi) GetBodyResp(resp *http.Response) []byte {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Print response
	return body
}
