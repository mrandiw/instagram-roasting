package generate

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MrAndiw/instagram-roasting/app/helper"

	"github.com/MrAndiw/instagram-go-scraper/instagram"
)

type (
	ModuleGenerate interface {
		GetInstagramProfile(username string) instagram.InstagramProfile
		GenerateRoast(profile instagram.InstagramProfile) (string, error)
		WrapText(text string, width int) []string
	}

	moduleGenerate struct {
		Config helper.Config
	}
)

func NewGenerate(config helper.Config) ModuleGenerate {
	return &moduleGenerate{
		Config: config,
	}
}

func (module *moduleGenerate) GetInstagramProfile(username string) instagram.InstagramProfile {
	// Instantiate default collector
	ig := instagram.Init()

	Instagram := instagram.NewInstagram(ig)

	User := Instagram.GetFullBio(username)

	return User
}

func (module *moduleGenerate) GenerateRoast(profile instagram.InstagramProfile) (string, error) {

	Followers := strconv.Itoa(profile.Followers)
	igDesc := fmt.Sprintf("Name : %s , Bio : %s , Followers : %s", profile.Title, profile.Bio, Followers)

	joinStr := fmt.Sprintf("%s = %s", module.Config.TextRoast, igDesc)
	Response, err := GetGemini(joinStr, module.Config)
	if err != nil {
		return "", fmt.Errorf("[GenerateRoast] Error GetGemini : %s", err)
	}

	return Response, nil
}

// wrapText wraps the input text into lines based on the given width.
func (module *moduleGenerate) WrapText(text string, width int) []string {
	var lines []string
	paragraphs := strings.Split(text, "\n")

	for _, paragraph := range paragraphs {
		words := strings.Fields(paragraph)
		var currentLine strings.Builder

		for _, word := range words {
			if currentLine.Len()+len(word)+1 > width {
				lines = append(lines, currentLine.String())
				currentLine.Reset()
			}
			if currentLine.Len() > 0 {
				currentLine.WriteString(" ")
			}
			currentLine.WriteString(word)
		}
		if currentLine.Len() > 0 {
			lines = append(lines, currentLine.String())
		}
		lines = append(lines, "")
	}
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}
