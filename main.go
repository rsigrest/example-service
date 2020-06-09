package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"

	"github.com/bestateless/example-service/proto"
)

func main() {

	name := "example-service"

	app := &cli.App{
		Name: name,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: fmt.Sprintf("/var/stateless/%s/config.toml", name),
				Usage: "Path to a configuration file.",
			},
		},
		Action: func(c *cli.Context) error {

			// configure logger
			log.Logger = log.With().Caller().Logger()
			log.Logger = log.With().Str("service", name).Logger()

			log.Info().Msgf("Service %s started.", name)

			// default configuration values
			viper.SetDefault("ServiceAddress", ":9000")
			// viper.SetDefault("MetricsAddress", ":9001")

			// attempt to read configuration file
			viper.SetConfigFile(c.String("config"))
			if err := viper.ReadInConfig(); err != nil {
				log.Debug().Err(err)
			}

			// log final configuration
			if c, err := json.Marshal(viper.AllSettings()); err == nil {
				log.Debug().RawJSON("configuration", c).Msg("")
			}

			// create TCP listener for gRPC server
			l, err := net.Listen("tcp", c.String("ServiceAddress"))
			if err != nil {
				log.Fatal().
					Err(err).
					Msgf("Unable to create a new TCP listener on '%s'. Port may already be in use.", c.String("ServiceAddress"))
			}
			defer l.Close()

			// start gRPC service
			grpcServer := grpc.NewServer()
			proto.RegisterExampleServiceServer(grpcServer, &Service{})
			log.Fatal().
				Err(grpcServer.Serve(l)).
				Msg("gRPC server exited.")

			return nil
		},
	}
	log.Fatal().Err(app.Run(os.Args))
}
