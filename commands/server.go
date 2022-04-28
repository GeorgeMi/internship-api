package commands

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/GeorgeMi/internship-api/rest"
	"github.com/mitchellh/cli"
)

type ServerCommand struct {
	r *rest.Service
}

func NewServerCommand() (cli.Command, error) {
	var err error
	c := new(ServerCommand)
	if c.r, err = rest.NewService(); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *ServerCommand) Synopsis() string {
	return "serve Internship API for front-end consumption"
}

func (c *ServerCommand) Help() string {
	return "try using ./[buildName] run"
}

func (c *ServerCommand) doRun(errc chan<- error) {
	log.Println("Listening on :3005")
	errc <- http.ListenAndServe(":3005", c.r.Container())
}

func (c *ServerCommand) Run(_ []string) int {
	errc := make(chan error, 1)
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, syscall.SIGTERM)

	// start server
	go c.doRun(errc)

	// wait for signal to exit
	select {
	case err := <-errc:
		log.Printf("Shutting down due to an unexpected error: %s", err)
		return 1
	case sig := <-sigc:
		log.Printf("Shutting down because of signal: %s", sig)
		return 0
	}
}
