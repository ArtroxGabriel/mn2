package ui

import "github.com/charmbracelet/lipgloss"

var (
	DocStyle          = lipgloss.NewStyle().Margin(1, 2)
	TitleStyle        = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("62")).MarginBottom(1)
	ListItemStyle     = lipgloss.NewStyle().PaddingLeft(2)
	SelectedItemStyle = lipgloss.NewStyle().PaddingLeft(0).Foreground(lipgloss.Color("75")).SetString("> ")
	InputValueStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("220"))
	InputLabelStyle   = lipgloss.NewStyle()
	FocusedInputStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("208")).Bold(true)
	ErrorStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
	ResultStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("46")).Bold(true)
	HelpStyle         = lipgloss.NewStyle().Faint(true).Italic(true).MarginTop(1)

	CursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("208"))
)

func RenderListItem(content string, selected bool) string {
	if selected {
		return SelectedItemStyle.Render(content)
	}
	return ListItemStyle.Render(content)
}
