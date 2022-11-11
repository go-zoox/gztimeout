package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-zoox/cli"
	"github.com/go-zoox/command"
	"github.com/go-zoox/timeout"
)

func main() {
	app := cli.NewSingleProgram(&cli.SingleProgramConfig{
		Name:    "timeout",
		Usage:   "timeout exec shell script",
		Version: Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "age",
				Usage:    "Timeout Age",
				Aliases:  []string{"a"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "exec",
				Usage:    "Exec Commands",
				Aliases:  []string{"e"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "message",
				Usage:   "Timeout Message",
				Aliases: []string{"m"},
			},
		},
	})

	app.Command(func(ctx *cli.Context) error {
		maxAgeRaw := ctx.String("age")
		script := ctx.String("exec")
		message := ctx.String("message")

		if message == "" {
			message = fmt.Sprintf("timeout to exec: %s", script)
		}

		maxAge, err := strconv.Atoi(maxAgeRaw)
		if err != nil {
			return err
		}

		cmd := &command.Command{
			Script: script,
			Environment: map[string]string{
				"POWER_BY": "GoZoox",
			},
		}

		if err := timeout.Timeout(cmd.Run, time.Duration(maxAge)*time.Second, message); err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}

		return nil
	})

	app.Run()
}
