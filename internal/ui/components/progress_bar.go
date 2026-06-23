package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type ProgressBar struct {
	width      int
	fillColor  string
	emptyColor string
	labelStyle lipgloss.Style
}

func NewProgressBar(width int, fillColor string, emptyColor string, labelColor string) ProgressBar {
	return ProgressBar{
		width:      width,
		fillColor:  fillColor,
		emptyColor: emptyColor,
		labelStyle: lipgloss.NewStyle().Foreground(lipgloss.Color(labelColor)),
	}
}

func (pb ProgressBar) Render(value float64, label string) string {
	if value < 0 {
		value = 0
	}
	if value > 1 {
		value = 1
	}

	filled := int(float64(pb.width) * value)
	empty := pb.width - filled

	fill := lipgloss.NewStyle().Foreground(lipgloss.Color(pb.fillColor)).Render(strings.Repeat("█", filled))
	void := lipgloss.NewStyle().Foreground(lipgloss.Color(pb.emptyColor)).Render(strings.Repeat("░", empty))

	pct := lipgloss.NewStyle().Foreground(lipgloss.Color(pb.fillColor)).Render(fmt.Sprintf(" %3.0f%%", value*100))

	lbl := pb.labelStyle.Render(label)

	return lbl + " " + fill + void + pct
}

type StatsPanel struct {
	heartFull  string
	heartEmpty string
	streakCol  string
	scoreCol   string
	labelCol   string
}

func NewStatsPanel(heartFull, heartEmpty, streakCol, scoreCol, labelCol string) StatsPanel {
	return StatsPanel{
		heartFull:  heartFull,
		heartEmpty: heartEmpty,
		streakCol:  streakCol,
		scoreCol:   scoreCol,
		labelCol:   labelCol,
	}
}

func (sp StatsPanel) Render(livesRemaining, maxLives, streak, score int) string {
	fullHeart := lipgloss.NewStyle().Foreground(lipgloss.Color(sp.heartFull)).Render("♥")
	emptyHeart := lipgloss.NewStyle().Foreground(lipgloss.Color(sp.heartEmpty)).Render("♡")

	hearts := ""
	if maxLives > 0 {
		for i := 0; i < maxLives; i++ {
			if i < livesRemaining {
				hearts += fullHeart + " "
			} else {
				hearts += emptyHeart + " "
			}
		}
	} else {
		hearts = lipgloss.NewStyle().Foreground(lipgloss.Color(sp.heartFull)).Render("∞")
	}
	label := lipgloss.NewStyle().Foreground(lipgloss.Color(sp.labelCol))

	streakStr := label.Render("streak ") + lipgloss.NewStyle().Foreground(lipgloss.Color(sp.streakCol)).Bold(true).Render(fmt.Sprintf("%d", streak))
	scoreStr := label.Render("score ") + lipgloss.NewStyle().Foreground(lipgloss.Color(sp.scoreCol)).Bold(true).Render(fmt.Sprintf("%d", score))

	return hearts + " " + streakStr + " " + scoreStr
}
