// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package api

// Message represents a message
type Message struct {
	ID                    int             `json:"message_id"`
	From                  *Victim         `json:"from,omitempty"`
	Timestamp             int64           `json:"date"`
	Chat                  *Victim         `json:"chat"`
	ForwardFrom           *Victim         `json:"forward_from,omitempty"`
	ForwardChat           *Victim         `json:"forward_from_chat,omitempty"`
	ForwardTimestamp      int64           `json:"forward_date,omitempty"`
	ReplyTo               *Message        `json:"reply_to_message,omitempty"`
	EditTimestamp         int64           `json:"edit_date,omitempty"`
	Text                  string          `json:"text,omitempty"`
	Entities              []MessageEntity `json:"entities,omitempty"`
	Audio                 *Audio          `json:"audio,omitempty"`
	Document              *Document       `json:"document,omitempty"`
	Photo                 []PhotoSize     `json:"photo,omitempty"`
	Sticker               *Sticker        `json:"sticker,omitempty"`
	Video                 *Video          `json:"video,omitempty"`
	Voice                 *Voice          `json:"voice,omitempty"`
	Caption               string          `json:"caption,omitempty"`
	Contact               *Contact        `json:"contact,omitempty"`
	Location              *Location       `json:"location,omitempty"`
	Venue                 *Venue          `json:"venue,omitempty"`
	NewChatMember         *Victim         `json:"new_chat_member,omitempty"`
	LeftChatMember        *Victim         `json:"left_chat_member,omitempty"`
	NewChatTitle          string          `json:"new_chat_title,omitempty"`
	NewChatPhoto          []PhotoSize     `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto       bool            `json:"delete_chat_photo,omitempty"`
	GroupChatCreated      bool            `json:"group_chat_created,omitempty"`
	SuperGroupChatCreated bool            `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated    bool            `json:"channel_chat_created,omitempty"`
	MigrateTo             int64           `json:"migrate_to_chat_id,omitempty"`
	MigrateFrom           int64           `json:"migrate_from_chat_id,omitempty"`
	Pinned                *Message        `json:"pinned_message,omitempty"`
}

// valid message entity types
const (
	MentionEntity     = "mention"
	HashTagEntity     = "hashtag"
	BotCommandEntity  = "bot_command"
	URLEntity         = "url"
	EmailEntity       = "email"
	BoldEntity        = "bold"
	ItalicEntity      = "italic"
	CodeEntity        = "code"
	PreEntity         = "pre"
	TextLinkEntity    = "text_link"
	TextMentionEntity = "text_mention"
)

// MessageEntity represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	Type   string  `json:"type"`
	Offset int     `json:"offset"`
	Length int     `json:"length"`
	URL    string  `json:"url,omitempty"`
	User   *Victim `json:"user,omitempty"`
}

// PhotoSize represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size,omitempty"`
}

// Audio represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	FileID    string `json:"file_id"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer,omitempty"`
	Title     string `json:"title,omitempty"`
	MimeType  string `json:"mime_type,omitempty"`
	FileSize  int    `json:"file_size,omitempty"`
}

// Document represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	FileID   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileName string     `json:"file_name,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// Sticker represents a sticker.
type Sticker struct {
	FileID   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	Emoji    string     `json:"emoji,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// Video represents a video file
type Video struct {
	FileID   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	MimeType string     `json:"mime_type,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// Voice represents a voice note.
type Voice struct {
	FileID   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type,omitempty"`
	FileSize int    `json:"file_size,omitempty"`
}

// Contact represents a phone contact.
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserID      int64  `json:"user_id,omitempty"`
}

// Location represents a point on the map.
type Location struct {
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`
}

// Venue represents a venue.
type Venue struct {
	Location   *Location `json:"location"`
	Title      string    `json:"title"`
	Address    string    `json:"address"`
	Foursquare string    `json:"fourscuare_id,omitempty"`
}
