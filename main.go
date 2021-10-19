package main

func main() {
  // Initialize dependencies, pass them to the Application.
	logger := NewLogger()
	app := myapp.New(logger)

	// Wait for shut down in a separate goroutine.
	errCh := make(chan error)
	go func() {
		shutdownCh := make(chan os.Signal)
		signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)
		<-shutdownCh

		errCh <- app.Shutdown()
	}()

	// Start the server and handle any errors.
	if err := app.Start(); err != nil {
		logger.Fatal().Msg(err.Error())
	}
	// Handle shutdown errors.
	if err := <-errCh; err != nil {
		logger.Warn().Msg(err.Error())
	}
}
