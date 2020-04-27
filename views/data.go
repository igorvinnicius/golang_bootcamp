package views

import "github.com/igorvinnicius/lenslocked-go-web/models"

const(
	AlertLevelError = "danger"
	AlertLevelWarning = "warning"
	AlertLevelInfo = "info"
	AlertLevelSuccess = "success"

	AlertGenericMessage = "Something went wrong. Please try again!"
)

type Alert struct {
	Level string
	Message string
}

type Data struct {
	Alert *Alert
	User *models.User
	Yield interface {}
}

func (d *Data) SetAlert(err error) {
	
	if pErr, ok := err.(PublicError); ok {

		d.Alert = &Alert {
			Level: AlertLevelError,
			Message: pErr.Public(),
		}
	} else {
		d.Alert = &Alert {
			Level: AlertLevelError,
			Message: AlertGenericMessage,
		}
	}
}

func (d *Data) AlertError(message string) {
	d.Alert = &Alert {
		Level: AlertLevelError,
		Message: message,
	}
}

type PublicError interface {
	error
	Public() string
}