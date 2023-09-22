package dira

import (
	"github.com/mindwingx/gocally"
	"net/http"
)

func (ce *CommandEntity) ExecuteContainerCommand() (res string, err error) {
	url, execOptions := prepareExecPayload(ce.GetRemoteAccess(), ce.GetContainerId())
	exec := gocally.SetRequest().WithUrl(url)

	if ce.GetRemoteAccess().GetTimeout() > 0 {
		exec.SetRequestTimeout(ce.GetRemoteAccess().GetTimeout())
	}

	stdout := exec.SetBody(execOptions).Post()
	result, err := stdout.Response()

	defer result.Body.Close()

	if result.StatusCode != http.StatusOK {
		return "", err
	}

	return handleStdoutResult(ce.GetRemoteAccess().GetRemoveMatching(), result.Body)
}
