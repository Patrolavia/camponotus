// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package raw

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
)

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
