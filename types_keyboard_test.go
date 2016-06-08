// This file is part of Camponotus
// Camponotus is free software: see LICENSE.txt for more details.

package telegram

import (
	"reflect"
	"testing"
)

func TestFormatKeyboard(t *testing.T) {
	tbl := []struct {
		ncol   int
		src    []KeyboardButton
		expect [][]KeyboardButton
	}{
		{
			3,
			[]KeyboardButton{
				KeyboardButton{Text: "1"},
				KeyboardButton{Text: "2"},
				KeyboardButton{Text: "3"},
				KeyboardButton{Text: "4"},
				KeyboardButton{Text: "5"},
				KeyboardButton{Text: "6"},
			},
			[][]KeyboardButton{
				[]KeyboardButton{
					KeyboardButton{Text: "1"},
					KeyboardButton{Text: "2"},
					KeyboardButton{Text: "3"},
				},
				[]KeyboardButton{
					KeyboardButton{Text: "4"},
					KeyboardButton{Text: "5"},
					KeyboardButton{Text: "6"},
				},
			},
		},
		{
			3,
			[]KeyboardButton{
				KeyboardButton{Text: "1"},
				KeyboardButton{Text: "2"},
				KeyboardButton{Text: "3"},
				KeyboardButton{Text: "4"},
				KeyboardButton{Text: "5"},
			},
			[][]KeyboardButton{
				[]KeyboardButton{
					KeyboardButton{Text: "1"},
					KeyboardButton{Text: "2"},
					KeyboardButton{Text: "3"},
				},
				[]KeyboardButton{
					KeyboardButton{Text: "4"},
					KeyboardButton{Text: "5"},
				},
			},
		},
		{
			1,
			[]KeyboardButton{
				KeyboardButton{Text: "1"},
				KeyboardButton{Text: "2"},
				KeyboardButton{Text: "3"},
			},
			[][]KeyboardButton{
				[]KeyboardButton{
					KeyboardButton{Text: "1"},
				},
				[]KeyboardButton{
					KeyboardButton{Text: "2"},
				},
				[]KeyboardButton{
					KeyboardButton{Text: "3"},
				},
			},
		},
	}

	for _, data := range tbl {
		actual := FormatKeyboard(data.src, data.ncol)
		if !reflect.DeepEqual(actual, data.expect) {
			t.Errorf(
				"FormatKeyboard(%d buttons, %d in a row) should have %d rows, got %d x %d (row x col)",
				len(data.src), data.ncol, len(data.expect), len(actual), len(actual[0]))

		}
	}
}

func TestFormatInlineKeyboard(t *testing.T) {
	tbl := []struct {
		ncol   int
		src    []InlineKeyboardButton
		expect [][]InlineKeyboardButton
	}{
		{
			3,
			[]InlineKeyboardButton{
				InlineKeyboardButton{Text: "1"},
				InlineKeyboardButton{Text: "2"},
				InlineKeyboardButton{Text: "3"},
				InlineKeyboardButton{Text: "4"},
				InlineKeyboardButton{Text: "5"},
				InlineKeyboardButton{Text: "6"},
			},
			[][]InlineKeyboardButton{
				[]InlineKeyboardButton{
					InlineKeyboardButton{Text: "1"},
					InlineKeyboardButton{Text: "2"},
					InlineKeyboardButton{Text: "3"},
				},
				[]InlineKeyboardButton{
					InlineKeyboardButton{Text: "4"},
					InlineKeyboardButton{Text: "5"},
					InlineKeyboardButton{Text: "6"},
				},
			},
		},
		{
			3,
			[]InlineKeyboardButton{
				InlineKeyboardButton{Text: "1"},
				InlineKeyboardButton{Text: "2"},
				InlineKeyboardButton{Text: "3"},
				InlineKeyboardButton{Text: "4"},
				InlineKeyboardButton{Text: "5"},
			},
			[][]InlineKeyboardButton{
				[]InlineKeyboardButton{
					InlineKeyboardButton{Text: "1"},
					InlineKeyboardButton{Text: "2"},
					InlineKeyboardButton{Text: "3"},
				},
				[]InlineKeyboardButton{
					InlineKeyboardButton{Text: "4"},
					InlineKeyboardButton{Text: "5"},
				},
			},
		},
		{
			1,
			[]InlineKeyboardButton{
				InlineKeyboardButton{Text: "1"},
				InlineKeyboardButton{Text: "2"},
				InlineKeyboardButton{Text: "3"},
			},
			[][]InlineKeyboardButton{
				[]InlineKeyboardButton{
					InlineKeyboardButton{Text: "1"},
				},
				[]InlineKeyboardButton{
					InlineKeyboardButton{Text: "2"},
				},
				[]InlineKeyboardButton{
					InlineKeyboardButton{Text: "3"},
				},
			},
		},
	}

	for _, data := range tbl {
		actual := FormatInlineKeyboard(data.src, data.ncol)
		if !reflect.DeepEqual(actual, data.expect) {
			t.Errorf(
				"FormatInlineKeyboard(%d buttons, %d in a row) should have %d rows, got %d x %d (row x col)",
				len(data.src), data.ncol, len(data.expect), len(actual), len(actual[0]))

		}
	}
}
