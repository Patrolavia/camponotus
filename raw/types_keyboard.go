// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package raw

// ReplyKeyboardMarkup represents a custom keyboard with reply options.
type ReplyKeyboardMarkup struct {
	Keyboard  [][]KeyboardButton `json:"keyboard"`
	Resize    bool               `json:"resize_keyboard,omitempty"`
	Once      bool               `json:"one_time_keyboard,omitempty"`
	Selective bool               `json:"selective,omitempty"`
}

// KeyboardButton represents one button of the reply keyboard.
// For simple text buttons String can be used instead of this object to specify text of the button.
// Optional fields are mutually exclusive.
type KeyboardButton struct {
	Text     string `json:"text"`
	Contact  bool   `json:"request_contact,omitempty"`
	Location bool   `json:"request_location,omitempty"`
}

// ReplyKeyboardHide will hide the current custom keyboard and display the default letter-keyboard.
type ReplyKeyboardHide struct {
	Hide      bool `json:"hide_keyboard,omitempty"`
	Selective bool `json:"selective,omitempty"`
}

// InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	Keyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton represents one button of an inline keyboard.
// You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	Text   string `json:"text"`
	URL    string `json:"url,omitempty"`
	Data   string `json:"callback_data,omitempty"`
	Switch string `json:"switch_inline_query,omitempty"`
}

// CallbackQuery represents an incoming callback query from a callback button in an inline keyboard.
type CallbackQuery struct {
	ID       string   `json:"id"`
	From     *Victim  `json:"from"`
	Message  *Message `json:"message,omitempty"`
	InlineID string   `json:"inline_message_id,omitempty"`
	Data     string   `json:"data,omitempty"`
}
