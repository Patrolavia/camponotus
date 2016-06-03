// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package telegram

import "io"

// Fake is mocked telegram API, provides dumb response for you.
//
// You should not use this directly, embbed it and overwrite the methods you need instead.
func Fake(me *Victim) API {
	return &fake{me}
}

type fake struct {
	Me *Victim
}

func (f *fake) AnswerCallbackQuery(query, text string, alert bool) error {
	return nil
}

func (f *fake) AnswerInlineQuery(query string, results []InlineQueryResult, cache int, personal bool, next, pm, pmParam string) error {
	return nil
}

func (f *fake) EditCaption(chat string, msg int, caption, mode string, noPreview bool, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *fake) EditInlineCaption(msg, caption, mode string, noPreview bool, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *fake) EditInlineMarkup(msg string, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *fake) EditInlineText(msg, text, mode string, noPreview bool, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *fake) EditMarkup(chat string, msg int, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *fake) EditText(chat string, msg int, text, mode string, noPreview bool, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *fake) ForwardMessage(to, from string, silent bool, message int) (*Message, error) {
	return nil, nil
}

func (f *fake) GetChat(chat string) (*Victim, error) {
	return nil, nil
}

func (f *fake) GetChatAdministrators(chat string) ([]ChatMember, error) {
	return []ChatMember{}, nil
}

func (f *fake) GetChatMember(chat string) (*ChatMember, error) {
	return nil, nil
}

func (f *fake) GetChatMembersCount(chat string) (int, error) {
	return 0, nil
}

func (f *fake) GetFile(file string) (*File, error) {
	return nil, nil
}

func (f *fake) GetMe() (*Victim, error) {
	return f.Me, nil
}

func (f *fake) GetUpdates(offset, limit, timeout int) ([]Update, error) {
	return []Update{}, nil
}

func (f *fake) GetUserProfilePhotos(user, offset, limit int) (*UserProfilePhotos, error) {
	return nil, nil
}

func (f *fake) KickChatMember(chat string, user int) error {
	return nil
}

func (f *fake) LeaveChat(chat string) error {
	return nil
}

func (f *fake) SendAudio(chat, audio string, duration int, performer, title string, opts *Options) (*Message, error) {
	return nil, nil
}

func (f *fake) SendChatAction(chat, action string) error {
	return nil
}

func (f *fake) SendContact(chat, phone, firstName, lastName string, opts *Options) (*Message, error) {
	return nil, nil
}

func (f *fake) SendDocument(chat, document, caption string, opts *Options) (*Message, error) {
	return nil, nil
}

func (f *fake) SendLocation(chat string, lat, lng float64, opts *Options) (*Message, error) {
	return nil, nil
}

func (f *fake) SendMessage(chat, text string, opts *Options) (*Message, error) {
	return nil, nil
}

func (f *fake) SendPhoto(chat, photo, caption string, opts *Options) (*Message, error) {
	return nil, nil
}

func (f *fake) SendSticker(chat, sticker, caption string, opts *Options) (*Message, error) {
	return nil, nil
}

func (f *fake) SendVenue(chat string, lat, lng float64, title, addr, foursq string, opts *Options) (*Message, error) {
	return nil, nil
}

func (f *fake) SendVideo(chat, video string, duration, width, height int, caption string, opts *Options) (*Message, error) {
	return nil, nil
}

func (f *fake) SendVoice(chat, voice string, duration int, opts *Options) (*Message, error) {
	return nil, nil
}

func (f *fake) SetWebhook(cb string, certificate io.Reader) error {
	return nil
}

func (f *fake) UnbanChatMember(chat string, user int) error {
	return nil
}

func (f *fake) UploadAudio(chat string, audio io.Reader, duration int, performer, title string, opts *Options) (*Message, error) {
	return nil, nil
}

func (f *fake) UploadDocument(chat string, document io.Reader, caption string, opts *Options) (*Message, error) {
	return nil, nil
}

func (f *fake) UploadPhoto(chat string, photo io.Reader, caption string, opts *Options) (*Message, error) {
	return nil, nil
}

func (f *fake) UploadSticker(chat string, sticker io.Reader, caption string, opts *Options) (*Message, error) {
	return nil, nil
}

func (f *fake) UploadVideo(chat string, video io.Reader, duration, width, height int, caption string, opts *Options) (*Message, error) {
	return nil, nil
}

func (f *fake) UploadVoice(chat string, voice io.Reader, duration int, opts *Options) (*Message, error) {
	return nil, nil
}
