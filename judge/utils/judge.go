package utils

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"xyz.xeonds/nano-oj/database/model"
)

func ParseResult(out io.Reader) (model.Status, string, error) {
	var result model.Status
	var info string
	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Result: ") {
			res, _ := strconv.Atoi(line[8:])
			result = model.Status(res)
		} else if strings.HasPrefix(line, "Info: ") {
			info += line[6:] + "\n"
		}
	}
	if err := scanner.Err(); err != nil {
		return model.CompilationError, "Failed to read container logs", err
	}
	return result, info, nil
}
