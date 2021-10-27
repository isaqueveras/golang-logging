package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	file, err := os.OpenFile("logrus/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.SetOutput(file)
	logrus.SetReportCaller(true)
	logrus.StandardLogger()

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.WithFields(logrus.Fields{"Day": "foo", "Time": "bar"}).Info("This is a json message")

	logrus.Info("This is some info")
	logrus.Warn("This is some warning")
	logrus.Error("This is some error")
}
