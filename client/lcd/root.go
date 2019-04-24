package lcd

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/log"
	rpcserver "github.com/tendermint/tendermint/rpc/lib/server"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	keybase "github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/server"

	// Import statik for light client stuff
	_ "github.com/hashgard/hashgard/client/lcd/statik"
)

// RestServer represents the Light Client Rest server
type RestServer struct {
	Mux     *mux.Router
	CliCtx  context.CLIContext
	KeyBase keybase.Keybase
	Cdc     *codec.Codec

	log         log.Logger
	listener    net.Listener
	fingerprint string
}

// NewRestServer creates a new rest server instance
func NewRestServer(cdc *codec.Codec) *RestServer {
	r := mux.NewRouter()
	cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout)).With("module", "LCD")

	return &RestServer{
		Mux:    r,
		CliCtx: cliCtx,
		Cdc:    cdc,

		log: logger,
	}
}

// Start starts the rest server
func (rs *RestServer) Start(listenAddr string, maxOpen int) (err error) {

	server.TrapSignal(func() {
		err := rs.listener.Close()
		rs.log.Error("error closing listener", "err", err)
	})

	rs.listener, err = rpcserver.Listen(
		listenAddr,
		rpcserver.Config{MaxOpenConnections: maxOpen},
	)
	if err != nil {
		return
	}
	rs.log.Info(fmt.Sprintf("Starting Hashgard Lite REST service (chain-id: %q)...",
		viper.GetString(client.FlagChainID)))


	return rpcserver.StartHTTPServer(rs.listener, rs.Mux, rs.log)
}

// ServeCommand will start a Hashgard Lite REST service as a blocking process. It
// takes a codec to create a RestServer object and a function to register all
// necessary routes.
func ServeCommand(cdc *codec.Codec, registerRoutesFn func(*RestServer)) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start LCD (light-client daemon), a local REST server",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			rs := NewRestServer(cdc)

			registerRoutesFn(rs)

			// Start the rest server and return error if one exists
			err = rs.Start(viper.GetString(client.FlagListenAddr),
				viper.GetInt(client.FlagMaxOpenConnections))

			return err
		},
	}

	return client.RegisterRestServerFlags(cmd)
}

// nolint
func (rs *RestServer) registerSwaggerUI() {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}
	staticServer := http.FileServer(statikFS)
	rs.Mux.PathPrefix("/swagger-ui/").Handler(http.StripPrefix("/swagger-ui/", staticServer))
}

func validateCertKeyFiles(certFile, keyFile string) error {
	if keyFile == "" {
		return errors.New("a key file is required")
	}
	if _, err := os.Stat(certFile); err != nil {
		return err
	}
	if _, err := os.Stat(keyFile); err != nil {
		return err
	}
	return nil
}
