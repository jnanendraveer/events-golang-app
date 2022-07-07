package CommonFunction

import (
	"fmt"
	"time"

	"github.com/jnanendraveer/events-golang-app/api/utils/StatusCode"
)

func Attachments(statusCode int64, err error, message string) map[string]interface{} {
	StatusText := StatusCode.StatusText
	attachment := map[string]interface{}{}
	attachment["code"] = statusCode
	attachment["status"] = "success"
	if err != nil {
		attachment["status"] = "Failed"
		attachment["message"] = fmt.Sprint(err)
		return attachment
	}
	if message == "" {
		attachment["message"] = StatusText(StatusCode.StatusOK)
		return attachment
	}
	attachment["message"] = message
	return attachment
}

func CurrentTime() time.Time {
	return time.Now()
}
