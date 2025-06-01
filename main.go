package main
import (
	"fmt"
	"os"
	"github.com/gocolly/colly"
)

func NewApp() *App {
	return &App{
		crawler: colly.NewCollector(),
	}
}



func main() {
	// Initialize the application
	app := NewApp()

	// Run the application
	if err := app.Run(); err != nil {
		panic(err)
	}
}