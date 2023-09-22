package dira

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

func prepareInquirePayload(e ExecutionAbstraction) (string, map[string]interface{}, string) {
	url := fmt.Sprintf(InquireUrl, e.GetUrl(), e.GetContainer())

	inquireOptions := map[string]interface{}{
		"AttachStdin":  false,
		"AttachStdout": true,
		"AttachStderr": true,
		"Tty":          false,
		"Cmd":          []interface{}{"sh", "-c", e.GetCommand()}, // bash is acceptable instead of "sh"
	}

	inquireOptJson, _ := json.Marshal(inquireOptions)
	inquireOptionJsonLen := strconv.Itoa(len(inquireOptJson))

	return url, inquireOptions, inquireOptionJsonLen
}

func prepareExecPayload(e ExecutionAbstraction, id string) (string, map[string]interface{}) {
	url := fmt.Sprintf(ExecUrl, e.GetUrl(), id)

	execOptions := map[string]interface{}{
		"Detach": false,
		"Tty":    false,
	}
	return url, execOptions
}

func removeMatching(containerStdout []byte) string {
	stdout := string(containerStdout)
	regex := regexp.MustCompile(RegexPattern)
	return regex.ReplaceAllString(stdout, "")
}

func handleStdoutResult(matching bool, result io.ReadCloser) (res string, err error) {
	containerStdOut, ioErr := io.ReadAll(result)

	if err = ioErr; err != nil {
		return "", ioErr
	}

	res = string(containerStdOut)

	if matching {
		res = removeMatching(containerStdOut)
	}

	return res, err
}
