package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/ArtroxGabriel/numeric-methods-cli/internal/ui"
)

func main() {

	p := tea.NewProgram(ui.NewMainModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Ocorreu um erro: %v\n", err)
		log.Fatal(err)
		os.Exit(1)
	}
}
