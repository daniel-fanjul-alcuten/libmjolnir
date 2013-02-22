package mjölnir

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func parseMFile(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(file)

	// skip the target
	_, err = reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	files := []string{}

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return files, nil
			}
			return nil, err
		}
		files = append(files, strings.Trim(line, " \t\n\\"))
	}
	return files, nil
}
