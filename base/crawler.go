// Copyright (c) 2015, The Tony Authors.
// All rights reserved.
//
// Author: Rentong Zhang <rentongzh@gmail.com>

package base

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"time"
)

type Crawler struct {
	utf8Converter *Utf8Converter
}

func NewCrawler() *Crawler {
	return &Crawler{
		utf8Converter: NewUtf8Converter(),
	}
}

func (c *Crawler) httpGet(url string) (error, []byte) {
	r, err := http.Get(url)
	if err != nil {
		return err, []byte("")
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err, []byte("")
	}
	return nil, c.utf8Converter.ToUTF8(b)
}

func (c *Crawler) GetRawHtml(url string, repeatTimes int) (error, string) {
	if len(url) == 0 {
		return errors.New("invalid url"), ""
	}
	var retErr error
	for i := 0; i < repeatTimes; i++ {
		err, html := c.httpGet(url)
		if err != nil {
			retErr = err
			time.Sleep(1 * time.Second)
			continue
		}
		return nil, string(html)
	}
	return retErr, ""
}

func (c *Crawler) GetDomHtml(url string, repeatTimes int) (error,
	*goquery.Document) {
	if len(url) == 0 {
		return errors.New("invalid url"), nil
	}

	var retErr error
	for i := 0; i < repeatTimes; i++ {
		err, html := c.httpGet(url)
		if err != nil {
			retErr = err
			time.Sleep(1 * time.Second)
			continue
		}
		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
		if err != nil {
			retErr = err
			continue
		}
		return nil, doc
	}
	return retErr, nil
}
