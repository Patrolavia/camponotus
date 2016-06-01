// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package raw

import (
	"io"
	"strconv"
)

// VictimType represents 5 kinds of valid receiver
type VictimType string

// these are valid receiver types
const (
	VTypeUser       VictimType = ""
	VTypePrivate    VictimType = "private"
	VTypeGroup      VictimType = "group"
	VTypeSuperGroup VictimType = "supergroup"
	VTypeChannel    VictimType = "channel"
)

// Victim represents a valid receiver, which can be a user or a chat
type Victim struct {
	ID        int64      `json:"id"`             // use int64 to compatible with chat
	Type      VictimType `json:"type,omitempty"` // omit if is user
	Title     string     `json:"title,omitempty"`
	FirstName string     `json:"first_name,omitempty"`
	LastName  string     `json:"last_name,omitempty"`
	Username  string     `json:"username,omitempty"`
}

// Identifier returns string representation of identifier.
//
// For user, private chat and group chat, it will be Victim.ID.
// For others, it will be Victim.Username prefixes with "@"
func (v Victim) Identifier() string {
	if v.Type == VTypeSuperGroup || v.Type == VTypeChannel {
		return "@" + v.Username
	}
	return strconv.FormatInt(v.ID, 10)
}

// UserProfilePhotos represents a user's profile picture.
type UserProfilePhotos struct {
	Length int           `json:"total_count"`
	Photos [][]PhotoSize `json:"photos"`
}

/*
File represents a file ready to be downloaded.
The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>.
It is guaranteed that the link will be valid for at least 1 hour.
When the link expires, a new one can be requested by calling getFile.

Maximum file size to download is 20 MB
*/
type File struct {
	ID   string `json:"file_id"`
	Size string `json:"file_size,omitempty"`
	Path string `json:"file_path,omitempty"`
}

// ChatMember contains information about pne member of the chat.
type ChatMember struct {
	User   *Victim `json:"user"`
	Status string  `json:"status"`
}

// these are valid member status
const (
	MStatusCreator = "creator"
	MStatusAdmin   = "administrator"
	MStatusMember  = "member"
	MStatusLeft    = "left"
	MStatusKicked  = "kicked"
)

// InputFile represents the contents of a file to be uploaded.
// you should use existing file if InputFile.FileID exists.
type InputFile struct {
	FileID string
	Reader io.Reader
}
