// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package telegram

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
)

// API maps all Telegram Bot api methods
type API interface {
	AnswerCallbackQuery(query, text string, alert bool) error
	AnswerInlineQuery(query string, results []InlineQueryResult, cache int, personal bool, next, pm, pmParam string) error
	EditCaption(chat string, msg int, caption, mode string, noPreview bool, markup ReplyMarkup) (*Message, error)
	EditInlineCaption(msg, caption, mode string, noPreview bool, markup ReplyMarkup) (*Message, error)
	EditInlineMarkup(msg string, markup ReplyMarkup) (*Message, error)
	EditInlineText(msg, text, mode string, noPreview bool, markup ReplyMarkup) (*Message, error)
	EditMarkup(chat string, msg int, markup ReplyMarkup) (*Message, error)
	EditText(chat string, msg int, text, mode string, noPreview bool, markup ReplyMarkup) (*Message, error)
	ForwardMessage(to, from string, silent bool, message int) (*Message, error)
	GetChat(chat string) (*Victim, error)
	GetChatAdministrators(chat string) ([]ChatMember, error)
	GetChatMember(chat string) (*ChatMember, error)
	GetChatMembersCount(chat string) (int, error)
	GetFile(file string) (*File, error)
	GetMe() (*Victim, error)
	GetUpdates(offset, limit, timeout int) ([]Update, error)
	GetUserProfilePhotos(user, offset, limit int) (*UserProfilePhotos, error)
	KickChatMember(chat string, user int) error
	LeaveChat(chat string) error
	SendAudio(chat, audio string, duration int, performer, title string, opts *Options) (*Message, error)
	SendChatAction(chat, action string) error
	SendContact(chat, phone, firstName, lastName string, opts *Options) (*Message, error)
	SendDocument(chat, document, caption string, opts *Options) (*Message, error)
	SendLocation(chat string, lat, lng float64, opts *Options) (*Message, error)
	SendMessage(chat, text string, opts *Options) (*Message, error)
	SendPhoto(chat, photo, caption string, opts *Options) (*Message, error)
	SendSticker(chat, sticker, caption string, opts *Options) (*Message, error)
	SendVenue(chat string, lat, lng float64, title, addr, foursq string, opts *Options) (*Message, error)
	SendVideo(chat, video string, duration, width, height int, caption string, opts *Options) (*Message, error)
	SendVoice(chat, voice string, duration int, opts *Options) (*Message, error)
	SetWebhook(cb string, certificate io.Reader) error
	UnbanChatMember(chat string, user int) error
	UploadAudio(chat string, audio io.Reader, duration int, performer, title string, opts *Options) (*Message, error)
	UploadDocument(chat string, document io.Reader, caption string, opts *Options) (*Message, error)
	UploadPhoto(chat string, photo io.Reader, caption string, opts *Options) (*Message, error)
	UploadSticker(chat string, sticker io.Reader, caption string, opts *Options) (*Message, error)
	UploadVideo(chat string, video io.Reader, duration, width, height int, caption string, opts *Options) (*Message, error)
	UploadVoice(chat string, voice io.Reader, duration int, opts *Options) (*Message, error)
}

type api struct {
	Client *http.Client // leave empty to use http.DefaultClient
	Token  string       // bot token, see https://core.telegram.org/bots#botfather
}

// New creates an API instance.
//
// You can pass nil to use http.DefaultClient
func New(token string, c *http.Client) API {
	return &api{c, token}
}

func (a *api) dest(method string) string {
	return "https://api.telegram.org/bot" + a.Token + "/" + method
}

func (a *api) client() (ret *http.Client) {
	ret = a.Client
	if ret == nil {
		ret = http.DefaultClient
	}

	return
}

func (a *api) call(method string, params url.Values) (ret []byte, err error) {
	c := a.client()

	res, err := c.PostForm(a.dest(method), params)
	if err != nil {
		return
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}

func (a *api) callAndSet(method string, params url.Values, res result) error {
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

func (a *api) callAndSetMsg(method string, params url.Values) (ret *Message, err error) {
	var m msgResult
	err = a.callAndSet(method, params, &m)
	return m.Message, err
}

func (a *api) upload(method string, params url.Values, field string, data io.Reader) (ret []byte, err error) {
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

func (a *api) uploadAndSet(method string, params url.Values, field string, data io.Reader, res result) error {
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

func (a *api) uploadAndSetMsg(method string, params url.Values, field string, data io.Reader) (*Message, error) {
	var m msgResult
	err := a.uploadAndSet(method, params, field, data, &m)
	return m.Message, err
}
