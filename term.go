package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func drawTitle() {
	divider := lipgloss.NewStyle().Padding(0, 1).Foreground(borderStyle).SetString("•").String()

	info := lipgloss.NewStyle().Foreground(specialStyle).Render
	welcome := strings.Builder{}
	
	onlineStatus := "[Status: Online]"
	if offline { onlineStatus = "[Status: Offline]" }

	title := lipgloss.NewStyle().
		Padding(0, 0, 0, 0).
		Width(80).
		BorderStyle(lipgloss.DoubleBorder()).
		BorderTop(true).
		BorderBottom(true).
		BorderForeground(borderStyle).
		Render("Enter a Command" + divider + info("[ Use 'help' ] v1.3.2 Duckmode  ") + info(onlineStatus))

	welcome.WriteString(title)

	fmt.Println(appStyle.Render(welcome.String()))
}

func drawStatus() {
	width := 14
	statusBarStyle := lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#343433", Dark: "#C1C6B2"}).
		Background(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#353533"})

	statusStyle := lipgloss.NewStyle().
		Inherit(statusBarStyle).
		Foreground(lipgloss.Color("#FFFDF5")).
		Background(lipgloss.Color("#FF5F87")).
		Padding(0, 1).
		MarginRight(1)

	statusText := lipgloss.NewStyle().Inherit(statusBarStyle).Align(lipgloss.Center)
	docStyle := lipgloss.NewStyle().Blink(true)
	doc := strings.Builder{}
	{
		w := lipgloss.Width

		statusKey := statusStyle.Render("COMMAND")
		statusVal := statusText.Copy().
			Width(width - w(statusKey)).
			Render("> ")

		bar := lipgloss.JoinHorizontal(lipgloss.Top,
			statusKey,
			statusVal,
		)

		doc.WriteString(statusBarStyle.Width(width).Render(bar))
	}
	
	fmt.Print(docStyle.Render(doc.String()))
}