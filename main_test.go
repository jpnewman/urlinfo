package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	logging "github.com/jpnewman/urlinfo/logging"
	"github.com/sirupsen/logrus"
)

func MockLogger() {
	logging.RootLogger = logrus.New()
	logging.RootLogger.Out = ioutil.Discard

	logging.Logger = logging.RootLogger.WithFields(logrus.Fields{
		"type": "test",
	})
}

func setup() {
	fmt.Println("Test Setup...")

	MockLogger()
}

func teardown() {
	fmt.Println("Test Teardown...")

	logging.LogFileClose()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
