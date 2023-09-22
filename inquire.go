package dira

import (
	"errors"
	"github.com/mindwingx/gocally"
	"net/http"
)

type InquireResponse struct {
	Id     string `json:"id,omitempty"`
	ErrMsg string `json:"message,omitempty"`
}

func (ce *CommandEntity) InquireContainer() CommandAbstraction {
	var response InquireResponse
	url, inquireOptions, inquireOptionsJsonLen := prepareInquirePayload(ce.rae)

	inquire := gocally.SetRequest().
		WithUrl(url).
		SetHeader("Content-Length", inquireOptionsJsonLen).
		SetBody(inquireOptions).
		Post()

	status, err := inquire.Entity(&response)

	if status["status_code"].(int) != http.StatusCreated && err == nil {
		err = errors.New(response.ErrMsg)
	}

	ce.SetContainerId(response.Id).
		SetError(err)

	return ce
}
