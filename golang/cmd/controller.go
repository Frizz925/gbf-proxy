package cmd

import (
	"errors"
	"log"

	"github.com/Frizz925/gbf-proxy/golang/controller"

	"github.com/spf13/cobra"
)

var controllerCmd = &cobra.Command{
	Use:   "controller <listen-address> <cache-address> <web-address> <web-hostname>",
	Short: "Start the Granblue Proxy controller service",
	Long: `Arguments:
  listen-address  The address this service should listen at
  cache-address   The address for the cache service
  web-address     The address for the web server serving static files
  web-hostname    The hostname for the web server
`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		nargs := len(args)
		if nargs < 1 {
			return errors.New("Missing listen-address argument")
		} else if nargs < 2 {
			return errors.New("Missing cache-address argument")
		} else if nargs < 3 {
			return errors.New("Missing web-address argument")
		} else if nargs < 4 {
			return errors.New("Missing web-hostname argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		s := controller.New(&controller.ServerConfig{
			CacheAddr: args[1],
			WebAddr:   args[2],
			WebHost:   args[3],
		})
		_, err := s.Open(args[0])
		if err != nil {
			panic(err)
		}
		log.Printf("Controller at %s -> Web server at %s", args[0], args[1])
		s.WaitGroup().Wait()
	},
}

func init() {
	rootCmd.AddCommand(controllerCmd)
}
