package utils

import "os"

func GetExecutableName() string {
	path := os.Args[0]

	i := len(path) - 1

	for ; i != 0; i-- {
		if path[i] == '/' {
			break
		}
	}

	return path[i+1:]
}
