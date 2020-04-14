package views

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
	Yeld interface {}
}