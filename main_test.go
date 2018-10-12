package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func MockLogger() {
	RootLogger = logrus.New()
	RootLogger.Out = ioutil.Discard

	Logger = RootLogger.WithFields(logrus.Fields{
		"type": "test",
	})
}

func setup() {
	fmt.Println("Test Setup...")

	MockLogger()
}

func teardown() {
	fmt.Println("Test Teardown...")

	LogFileClose()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
