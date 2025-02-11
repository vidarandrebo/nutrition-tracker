package utils

import (
	"bufio"
	"io"
	"strings"
)

func ReadEnv(fs io.Reader) map[string]string {
	keys := make(map[string]string)

	scanner := bufio.NewScanner(fs)

	for scanner.Scan() {
		line := scanner.Text()
		kvPair := strings.SplitN(line, "=", 2)
		keys[kvPair[0]] = kvPair[1]
	}
	return keys
}
