package cli

import (
	"github.com/fatih/color"
	"log"
)

type Color struct {
}

func (c Color) errorColor() *color.Color {
	errorCli := color.New()
	errorCli.Add(color.FgWhite)
	errorCli.Add(color.BgRed)
	return errorCli
}

func (c Color) PrintErrorf(format string, args ...interface{}) {
	errorCli := c.errorColor()
	_, err := errorCli.Printf(format, args)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
}

func (c Color) PrintError(format string) {
	errorCli := c.errorColor()
	_, err := errorCli.Print(format)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
}

func (c Color) PrintSuccessF(format string, args ...interface{}) {
	color.Green(format, args)
}

func (c Color) PrintSuccess(str string) {
	color.Green(str)
}
