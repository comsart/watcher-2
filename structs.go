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

type Instructions struct {
	ShouldBeForced  bool     `json:"shouldBeForced"`
	AppsToTerminate []string `json:"appsToTerminate"`
}

type ResponseDTO[T any] struct {
	Success bool
	Message string
	Data    T
}
