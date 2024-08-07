package cmd

import (
	"log"
	"strings"

	"github.com/gdamore/tcell/v2"
)

type (
	ModuleScreen interface {
		Drawmodul(width, outputAreaHeight, inputAreaWidth, inputAreaY int, outputText []string, inputBuffer strings.Builder)
	}

	moduleScreen struct {
		Screen tcell.Screen
	}
)

func NewScreen(screen tcell.Screen) ModuleScreen {
	return &moduleScreen{
		Screen: screen,
	}
}

func Init() tcell.Screen {
	// Initialize the screen
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Failed to create screen: %v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("Failed to initialize screen: %v", err)
	}

	return screen
}

func (modul *moduleScreen) Drawmodul(width, outputAreaHeight, inputAreaWidth, inputAreaY int, outputText []string, inputBuffer strings.Builder) {

	// Define styles
	defaultStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	borderStyle := tcell.StyleDefault.Foreground(tcell.ColorGreen).Background(tcell.ColorBlack)

	modul.Screen.Clear()

	// Draw the output area border
	for x := 0; x < width; x++ {
		modul.Screen.SetContent(x, 0, '─', nil, borderStyle)                // Top border
		modul.Screen.SetContent(x, outputAreaHeight, '─', nil, borderStyle) // Bottom border
	}
	for y := 0; y <= outputAreaHeight; y++ {
		modul.Screen.SetContent(0, y, '│', nil, borderStyle)       // Left border
		modul.Screen.SetContent(width-1, y, '│', nil, borderStyle) // Right border
	}
	modul.Screen.SetContent(0, 0, '┌', nil, borderStyle)                      // Top-left corner
	modul.Screen.SetContent(width-1, 0, '┐', nil, borderStyle)                // Top-right corner
	modul.Screen.SetContent(0, outputAreaHeight, '└', nil, borderStyle)       // Bottom-left corner
	modul.Screen.SetContent(width-1, outputAreaHeight, '┘', nil, borderStyle) // Bottom-right corner

	// Draw the output text inside the output area
	for i, line := range outputText {
		for j, ch := range line {
			index := j + 2
			modul.Screen.SetContent(index, i+1, ch, nil, tcell.StyleDefault.Foreground(tcell.ColorGreenYellow).Background(tcell.ColorBlack))
		}
	}

	// Draw the input area border
	for x := 0; x < inputAreaWidth; x++ {
		modul.Screen.SetContent(x, inputAreaY+1, '─', nil, borderStyle) // Bottom border
	}

	// Draw the input prompt and current input text
	inputPrompt := "Input Instagram Username Here : "
	for i, ch := range inputPrompt {
		modul.Screen.SetContent(i, inputAreaY, ch, nil, defaultStyle)
	}
	inputText := inputBuffer.String()
	for i, ch := range inputText {
		modul.Screen.SetContent(i+len(inputPrompt), inputAreaY, ch, nil, defaultStyle)
	}

	modul.Screen.Show()
}
