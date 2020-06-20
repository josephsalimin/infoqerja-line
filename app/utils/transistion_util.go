package utils

import (
	constant "infoqerja-line/app/utils/constant"

	"github.com/line/line-bot-sdk-go/linebot"
)

// IsStateValid : Due to the usage of magic string and no enum implementation in golang, for a while i will use this validation to check availability of the state of certain UserData instance
func IsStateValid(state string) bool {
	switch state {
	case constant.WaitTitleInput:
		return true
	case constant.WaitDescInput:
		return true
	case constant.WaitDateInput:
		return true
	default:
		return false
	}
}

// GetSource : Get source for any event happening to bot
func GetSource(event linebot.Event) string {
	switch event.Source.Type {
	case linebot.EventSourceTypeUser:
		return event.Source.UserID
	case linebot.EventSourceTypeGroup:
		return event.Source.UserID
	case linebot.EventSourceTypeRoom:
		return event.Source.RoomID
	}
	return event.Source.UserID
}
