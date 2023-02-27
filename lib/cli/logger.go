package cli

import (
	"github.com/fatih/color"
	"log"
)

type Logger struct {
}

func (c Logger) errorColor() *color.Color {
	errorCli := color.New()
	errorCli.Add(color.FgWhite)
	errorCli.Add(color.BgRed)
	return errorCli
}

func (c Logger) PrintErrorf(format string, args ...interface{}) {
	errorCli := c.errorColor()
	_, err := errorCli.Printf(format, args)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
}

func (c Logger) PrintError(format string) {
	errorCli := c.errorColor()
	_, err := errorCli.Print(format)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
}

func (c Logger) PrintSuccessF(format string, args ...interface{}) {
	color.Green(format, args)
}

func (c Logger) PrintSuccess(str string) {
	color.Green(str)
}
