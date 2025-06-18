package main

type RequestDTO[T any] struct {
	Command  string `json:"command"`
	UserName string `json:"userName"`
	Data     *T     `json:"data"`
}

type LogRecord struct {
	Level          string `json:"level"`
	Message        string `json:"message"`
	LoggerName     string `json:"loggerName"`
	ScreenshotData string `json:"screenshotData"`
}

type ServerState struct {
	LessonIsDone    bool     `json:"shouldBeForced"`
	AppsToTerminate []string `json:"appsToTerminate"` // todo nie musza byc juz wysylane z serwera bo ta apke tez gituje
}

type ResponseDTO[T any] struct {
	Success bool
	Message string
	Data    T
}
