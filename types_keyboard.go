// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package telegram

import "encoding/json"

// ReplyMarkup is additional interface options for sending messages
type ReplyMarkup interface {
	Bytes() ([]byte, error)
}

// ReplyKeyboardMarkup represents a custom keyboard with reply options.
type ReplyKeyboardMarkup struct {
	Keyboard  [][]KeyboardButton `json:"keyboard"`
	Resize    bool               `json:"resize_keyboard,omitempty"`
	Once      bool               `json:"one_time_keyboard,omitempty"`
	Selective bool               `json:"selective,omitempty"`
}

// Bytes serializes the structure into JSON format
func (m *ReplyKeyboardMarkup) Bytes() ([]byte, error) {
	return json.Marshal(m)
}

// KeyboardButton represents one button of the reply keyboard.
// For simple text buttons String can be used instead of this object to specify text of the button.
// Optional fields are mutually exclusive.
type KeyboardButton struct {
	Text     string `json:"text"`
	Contact  bool   `json:"request_contact,omitempty"`
	Location bool   `json:"request_location,omitempty"`
}

// FormatKeyboard reformats keyboard buttons
//
// Say you have a serial of buttons [1, 2, 3, 4, 5], by calling this method with ncol=3, it
// returns [[1, 2, 3], [4, 5]]
func FormatKeyboard(buttons []KeyboardButton, ncol int) [][]KeyboardButton {
	var ret [][]KeyboardButton
	if len(buttons) < 1 {
		return ret
	}

	x := 0  // current column
	y := -1 // current row

	for _, b := range buttons {
		if x == 0 {
			y++
			ret = append(ret, []KeyboardButton{})
		}

		ret[y] = append(ret[y], b)

		x++
		if x >= ncol {
			x = 0
		}
	}

	return ret
}

// ReplyKeyboardHide will hide the current custom keyboard and display the default letter-keyboard.
type ReplyKeyboardHide struct {
	Hide      bool `json:"hide_keyboard,omitempty"`
	Selective bool `json:"selective,omitempty"`
}

// Bytes serializes the structure into JSON format
func (h *ReplyKeyboardHide) Bytes() ([]byte, error) {
	return json.Marshal(h)
}

// InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	Keyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// Bytes serializes the structure into JSON format
func (m *InlineKeyboardMarkup) Bytes() ([]byte, error) {
	return json.Marshal(m)
}

// InlineKeyboardButton represents one button of an inline keyboard.
// You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	Text   string `json:"text"`
	URL    string `json:"url,omitempty"`
	Data   string `json:"callback_data,omitempty"`
	Switch string `json:"switch_inline_query,omitempty"`
}

// FormatInlineKeyboard reformats keyboard buttons
//
// Say you have a serial of buttons [1, 2, 3, 4, 5], by calling this method with ncol=3, it
// returns [[1, 2, 3], [4, 5]]
func FormatInlineKeyboard(buttons []InlineKeyboardButton, ncol int) [][]InlineKeyboardButton {
	var ret [][]InlineKeyboardButton
	if len(buttons) < 1 {
		return ret
	}

	x := 0  // current column
	y := -1 // current row

	for _, b := range buttons {
		if x == 0 {
			y++
			ret = append(ret, []InlineKeyboardButton{})
		}

		ret[y] = append(ret[y], b)

		x++
		if x >= ncol {
			x = 0
		}
	}

	return ret
}

// ForceReply will display a reply interface to the user (act as if the user has selected the bot‘s message and tapped ’Reply').
type ForceReply struct {
	Reply     bool `json:"force_reply,omitempty"`
	Selective bool `json:"selective,omitempty"`
}

// Bytes serializes the structure into JSON format
func (r *ForceReply) Bytes() ([]byte, error) {
	return json.Marshal(r)
}

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard.
type CallbackQuery struct {
	ID       string   `json:"id"`
	From     *Victim  `json:"from"`
	Message  *Message `json:"message,omitempty"`
	InlineID string   `json:"inline_message_id,omitempty"`
	Data     string   `json:"data,omitempty"`
}
