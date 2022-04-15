package gtfo_overlay_sdk

import (
	"fmt"

	"github.com/imroc/req"
)

type baseClient struct {
	useDev bool
}

func (c baseClient) getUrlPrefix() string {
	if c.useDev {
		return "https://gtfo-dev.lan.r35.net"
	}
	return "https://gtfo.lan.r35.net"
}

func (c baseClient) getCommonHeaders() req.Header {
	return req.Header{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
}

func (c baseClient) makePutRequest(path string) (*req.Resp, error) {
	return req.Put(fmt.Sprintf("%s/%s", c.getUrlPrefix(), path), c.getCommonHeaders())
}

func (c baseClient) makeGetRequest(path string) (*req.Resp, error) {
	return req.Get(fmt.Sprintf("%s/%s", c.getUrlPrefix(), path), c.getCommonHeaders())
}
