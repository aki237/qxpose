package client

import (
	"errors"

	"github.com/urfave/cli"
)

// Init function initializes the client command
// commandline functionality and retursn the cli.Command
func Init() cli.Command {
	return cli.Command{
		Name:   "client",
		Usage:  "Run a client instance",
		Action: createClient,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "tunnel",
				Usage:    "Remote public tunnel address to connect to",
				Required: true,
			},
			cli.StringFlag{
				Name:     "local",
				Usage:    "Local TCP server to proxy the connections to",
				Required: true,
			},
			cli.UintFlag{
				Name:  "i,idle-timeout",
				Usage: "Idle timeout for the quic sessions (in seconds)",
				Value: 1800,
			},
		},
	}
}

func createClient(ctx *cli.Context) error {
	tunnel := ctx.String("tunnel")
	if tunnel == "" {
		return errors.New("Tunnel address cannot be empty")
	}

	local := ctx.String("local")
	if local == "" {
		return errors.New("Local address cannot be empty")
	}

	return NewClient(tunnel, local, ctx.Uint("i")).Start()
}
