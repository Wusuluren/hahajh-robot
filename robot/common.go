package robot

import (
	"errors"
	"fmt"
	"net/http"
)

func httpStatusCodeError(statusCode int) error {
	return errors.New(fmt.Sprintf("http status code: %d", statusCode))
}

func checkHttpRespError(resp *http.Response, err error) error {
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return httpStatusCodeError(resp.StatusCode)
	}
	return nil
}
