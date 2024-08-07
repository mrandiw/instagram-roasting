package main

import (
	"strings"

	"github.com/MrAndiw/instagram-roasting/app/cmd"
	"github.com/MrAndiw/instagram-roasting/app/generate"
	"github.com/MrAndiw/instagram-roasting/app/helper"

	"github.com/gdamore/tcell/v2"
)

func main() {
	// Initialize the screen
	screen := cmd.Init()
	defer screen.Fini()

	// Get screen size
	width, height := screen.Size()

	// Define areas
	outputAreaHeight := height - 3 // Reserve 3 lines for input
	inputAreaY := outputAreaHeight + 1
	inputAreaWidth := width

	var inputBuffer strings.Builder
	var outputText []string

	// get config
	config := helper.GetConfig()

	// load module
	moduleScreen := cmd.NewScreen(screen)
	moduleScreen.Drawmodul(width, outputAreaHeight, inputAreaWidth, inputAreaY, outputText, inputBuffer)

	// load module generate
	moduleGenerate := generate.NewGenerate(config)

	for {
		event := screen.PollEvent()
		switch e := event.(type) {
		case *tcell.EventKey:
			switch e.Key() {
			case tcell.KeyEnter:
				// When Enter is pressed, move input text to output box
				outputTextTemp := inputBuffer.String()
				inputBuffer.Reset()

				// get instagram profile
				IGProfile := moduleGenerate.GetInstagramProfile(outputTextTemp)

				// generate roasting
				roast := moduleGenerate.GenerateRoast(IGProfile)
				finalOutput := moduleGenerate.WrapText(roast, width-2) // Wrap text with width limit

				moduleScreen.Drawmodul(width, outputAreaHeight, inputAreaWidth, inputAreaY, finalOutput, inputBuffer)
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				// Handle backspace
				if inputBuffer.Len() > 0 {
					inputBuffer.Reset()
					moduleScreen.Drawmodul(width, outputAreaHeight, inputAreaWidth, inputAreaY, outputText, inputBuffer)
				}
			case tcell.KeyCtrlC:
				// Exit the program
				return
			default:
				// Handle other keys
				if e.Rune() != 0 {
					inputBuffer.WriteRune(e.Rune())
					moduleScreen.Drawmodul(width, outputAreaHeight, inputAreaWidth, inputAreaY, outputText, inputBuffer)
				}
			}
		}
	}
}
