package main

import "github.com/sirupsen/logrus"
import "fmt"

func main() {
	//logrus.Info("hello logrus")
	//logrus.WithFields(logrus.Fields{
	//	"event": "1",
	//	"topic": "2",
	//	"key":   "3",
	//}).Fatal("Failed to send event")
	level, _ := logrus.ParseLevel("info")
	logrus.SetLevel(level)
	logrus.Debug("hello logrus")
	fmt.Println(logrus.ErrorLevel)
}
