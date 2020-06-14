package utils

import (
	constant "infoqerja-line/app/utils/constant"
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
