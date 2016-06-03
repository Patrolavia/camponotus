// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package telegram

import (
	"io"
)

// Fake is mocked telegram API, provides dumb response for you.
//
// You should not use this directly, embbed it and overwrite the methods you need instead.
type Fake struct {
	Me *Victim
}

func (f *Fake) AnswerCallbackQuery(query, text string, alert bool) error {
	return nil
}

func (f *Fake) AnswerInlineQuery(query string, results []InlineQueryResult, cache int, personal bool, next, pm, pmParam string) error {
	return nil
}

func (f *Fake) EditCaption(chat string, msg int, caption, mode string, noPreview bool, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) EditInlineCaption(msg, caption, mode string, noPreview bool, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) EditInlineMarkup(msg string, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) EditInlineText(msg, text, mode string, noPreview bool, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) EditMarkup(chat string, msg int, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) EditText(chat string, msg int, text, mode string, noPreview bool, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) ForwardMessage(to, from string, silent bool, message int) (*Message, error) {
	return nil, nil
}

func (f *Fake) GetChat(chat string) (*Victim, error) {
	return nil, nil
}

func (f *Fake) GetChatAdministrators(chat string) ([]ChatMember, error) {
	return []ChatMember{}, nil
}

func (f *Fake) GetChatMember(chat string) (*ChatMember, error) {
	return nil, nil
}

func (f *Fake) GetChatMembersCount(chat string) (int, error) {
	return 0, nil
}

func (f *Fake) GetFile(file string) (*File, error) {
	return nil, nil
}

func (f *Fake) GetMe() (*Victim, error) {
	return f.Me, nil
}

func (f *Fake) GetUpdates(offset, limit, timeout int) ([]Update, error) {
	return []Update{}, nil
}

func (f *Fake) GetUserProfilePhotos(user, offset, limit int) (*UserProfilePhotos, error) {
	return nil, nil
}

func (f *Fake) KickChatMember(chat string, user int) error {
	return nil
}

func (f *Fake) LeaveChat(chat string) error {
	return nil
}

func (f *Fake) SendAudio(chat, audio string, duration int, performer, title string, silent bool, reply int, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) SendChatAction(chat, action string) error {
	return nil
}

func (f *Fake) SendContact(chat, phone, firstName, lastName string, silent bool, reply int, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) SendDocument(chat, document, caption string, silent bool, reply int, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) SendLocation(chat string, lat, lng float64, silent bool, reply int, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) SendMessage(chat, text, mode string, noPreview, silent bool, reply int, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) SendPhoto(chat, photo, caption string, silent bool, reply int, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) SendSticker(chat, sticker, caption string, silent bool, reply int, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) SendVenue(chat string, lat, lng float64, title, addr, foursq string, silent bool, reply int, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) SendVideo(chat, video string, duration, width, height int, caption string, silent bool, reply int, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) SendVoice(chat, voice string, duration int, silent bool, reply int, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) SetWebhook(cb string, certificate io.Reader) error {
	return nil
}

func (f *Fake) UnbanChatMember(chat string, user int) error {
	return nil
}

func (f *Fake) UploadAudio(chat string, audio io.Reader, duration int, performer, title string, silent bool, reply int, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) UploadDocument(chat string, document io.Reader, caption string, silent bool, reply int, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) UploadPhoto(chat string, photo io.Reader, caption string, silent bool, reply int, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) UploadSticker(chat string, sticker io.Reader, caption string, silent bool, reply int, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) UploadVideo(chat string, video io.Reader, duration, width, height int, caption string, silent bool, reply int, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}

func (f *Fake) UploadVoice(chat string, voice io.Reader, duration int, silent bool, reply int, markup ReplyMarkup) (*Message, error) {
	return nil, nil
}
