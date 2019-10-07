package backend

// Environment is where the gui will run.
type Environment interface {
	// CreateHandler creates a new handler.
	CreateHandler(Title string, X, Y, Width, Height int) (Handler, error)

	// Run the application.
	// this function must handler all events and must draw the window.
	Run() error
}
