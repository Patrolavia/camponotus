// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
)

// types for api method result

type result interface {
	OK() bool
}

type boolResult struct {
	Ok bool `json:"ok"`
}

func (r boolResult) OK() bool {
	return r.Ok
}

type intResult struct {
	boolResult
	Result int64 `json:"result"`
}

type updateResult struct {
	boolResult
	Updates []Update `json:"result"`
}

type userResult struct {
	boolResult
	User *Victim `json:"result"`
}

type msgResult struct {
	boolResult
	Message *Message `json:"result"`
}

type profilePhotoResult struct {
	boolResult
	UserProfilePhotos *UserProfilePhotos `json:"result"`
}

type fileResult struct {
	boolResult
	File *File `json:"result"`
}

type memberResult struct {
	boolResult
	Member *ChatMember `json:"result"`
}

type membersResult struct {
	boolResult
	Members []ChatMember `json:"result"`
}

// ErrNotOK means Telegram server returns fail on your method call
type ErrNotOK struct {
	Method string
	Params url.Values
	Field  string
	Data   io.Reader
}

func (e ErrNotOK) Error() string {
	msg := "is"
	if e.Data != nil {
		msg = "is not"
	}
	return fmt.Sprintf(
		"Telegram server returns fail on method %s, param %v, field %s, reader %s nil",
		e.Method,
		e.Params,
		e.Field,
		msg,
	)
}

// API maps all Telegram Bot API methods
type API struct {
	Client *http.Client // leave empty to use http.DefaultClient
	Token  string       // bot token, see https://core.telegram.org/bots#botfather
}

func (a *API) dest(method string) string {
	return "https://api.telegram.org/bot" + a.Token + "/" + method
}

func (a *API) client() (ret *http.Client) {
	ret = a.Client
	if ret == nil {
		ret = http.DefaultClient
	}

	return
}

func (a *API) call(method string, params url.Values) (ret []byte, err error) {
	c := a.client()

	res, err := c.PostForm(a.dest(method), params)
	if err != nil {
		return
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

func (a *API) callAndSet(method string, params url.Values, res result) error {
	if res == nil {
		res = &boolResult{}
	}

	buf, err := a.call(method, params)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, res)
	if err == nil && !res.OK() {
		err = &ErrNotOK{
			Method: method,
			Params: params,
		}
	}
	return err
}

func (a *API) callAndSetMsg(method string, params url.Values) (ret *Message, err error) {
	var m msgResult
	err = a.callAndSet(method, params, &m)
	return m.Message, err
}

func (a *API) upload(method string, params url.Values, field string, data io.Reader) (ret []byte, err error) {
	c := a.client()
	pr, pw := io.Pipe()
	defer pw.Close()
	w := multipart.NewWriter(pw)
	defer w.Close()
	ch := make(chan error)

	// preparing form data in background, hope this can reduce memory footprint
	go func(w *multipart.Writer, ch chan error) {
		for key, val := range params {
			w.WriteField(key, val[0])
		}
		fw, err := w.CreateFormFile(field, field)
		if err != nil {
			ch <- err
			return
		}
		_, err = io.Copy(fw, data)
		ch <- err
	}(w, ch)

	res, err := c.Post(a.dest(method), w.FormDataContentType(), pr)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if err := <-ch; err != nil {
		// upload failed
		return ret, err
	}
	close(ch)

	return ioutil.ReadAll(res.Body)
}

func (a *API) uploadAndSet(method string, params url.Values, field string, data io.Reader, res result) error {
	if res == nil {
		res = &boolResult{}
	}

	buf, err := a.upload(method, params, field, data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, res)
	if err == nil && !res.OK() {
		err = &ErrNotOK{method, params, field, data}
	}
	return err
}

func (a *API) uploadAndSetMsg(method string, params url.Values, field string, data io.Reader) (*Message, error) {
	var m msgResult
	err := a.uploadAndSet(method, params, field, data, &m)
	return m.Message, err
}

// helpers

func optStr(params url.Values, key, val string) {
	if val != "" {
		params.Set(key, val)
	}
}

func optInt(params url.Values, key string, val int) {
	if val != 0 {
		params.Set(key, strconv.Itoa(val))
	}
}

func optBool(params url.Values, key string, val bool) {
	if val {
		params.Set(key, "true")
	}
}

func optFloat(params url.Values, key string, val float64) {
	if val != 0.0 {
		params.Set(key, strconv.FormatFloat(val, 'f', -1, 64))
	}
}

func optJSON(params url.Values, key string, val []byte) {
	if val != nil && len(val) > 0 {
		params.Set(key, string(val))
	}
}
