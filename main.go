package main
import httpx

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

    // Serve r on port 80.
    server := httpx.NewServer(":80", r)
    err := server.Start()
    // Serve r on a systemd socket (FileDescriptorName=myapp-http).
    server := httpx.NewServer("systemd:myapp-http", r)
    err := server.Start()

    cert, err := tls.LoadX509KeyPair("/srv/cert.pem", "/srv/key.pem")
    // Serve r on port 443.
    server := httpx.NewServerTLS(":443", cert, r)
    err := server.Start()
    // Serve r on a systemd TLS socket (FileDescriptorName=myapp-https).
    server := httpx.NewServerTLS("systemd:myapp-https", cert, r)
    err := server.Start()

    // Serve up to 1000 simultaneous connections on port 8080.
    server := httpx.NewServer(":8080", r)
    server.MaxConnections = 1000
    err := server.Start()

