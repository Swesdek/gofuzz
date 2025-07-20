package dirfuzz

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type dirConfig struct {
	url    string
	client *http.Client
	ctx    context.Context
}

func NewDirConfig(url string, timeout time.Duration) dirConfig {
	client := http.DefaultClient
	client.Timeout = timeout

	return dirConfig{
		url:    url,
		client: client,
	}
}

func (dc dirConfig) ProcessWord(word string) (string, error) {
	completeUrl := strings.Replace(dc.url, "FUZZ", word, 1)

	req, err := http.NewRequest("GET", completeUrl, nil)
	if err != nil {
		return "", err
	}

	res, err := dc.client.Do(req)
	if err != nil {
		return "", err
	}

	if res.StatusCode == 404 {
		return "", err
	}

	sb := strings.Builder{}
	sb.WriteString(completeUrl)
	sb.WriteString(" ")
	sb.WriteString(strconv.Itoa(res.StatusCode))

	return sb.String(), nil
}
