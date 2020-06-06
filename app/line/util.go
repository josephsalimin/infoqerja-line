package line

import (
	"io/ioutil"
)

// GetMessageFromFile : Get text message from text files
func GetMessageFromFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return "", err
	}

	return string(data), nil
}
