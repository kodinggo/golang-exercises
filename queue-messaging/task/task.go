package task

// define event type
const (
	TaskSendEmail  = "SEND_EMAIL"
	TaskUploadFile = "UPLOAD_FILE"
)

// define event payload
type (
	SendEmail struct {
		From    string
		To      string
		Subject string
		Content string
	}

	UplaodFile struct {
		Origin      string
		Destination string
	}
)
