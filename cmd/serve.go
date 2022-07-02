package cmd

import (
	"fmt"
	"strings"

	"github.com/adrg/xdg"
	zlog "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/bow/courier/internal"
)

const (
	addrKey   = "addr"
	dbNameKey = "db-path"
)

var (
	relDBName     = "courier/courier.db"
	defaultDBName = fmt.Sprintf("$XDG_DATA_HOME/%s", relDBName)

	relUDS      = "courier/courier.socket"
	defaultAddr = fmt.Sprintf("$XDG_RUNTIME_DIR/%s", relUDS)
)

var serveCmdName = "serve"

var serveViper = newViper(serveCmdName)

var serveCmd = cobra.Command{
	Use:   serveCmdName,
	Short: "Start the server",
	RunE: func(cmd *cobra.Command, args []string) error {

		dbPath, err := resolveDBPath(serveViper)
		if err != nil {
			return err
		}

		addr, err := resolveUDSAddr(serveViper)
		if err != nil {
			return err
		}

		server, err := internal.NewServerBuilder().
			Address(addr).
			StorePath(dbPath).
			Logger(zlog.Logger.With().Logger()).
			Build()

		if err != nil {
			return err
		}

		return server.Serve()
	},
}

func init() {
	rootCmd.AddCommand(&serveCmd)

	flags := serveCmd.Flags()

	flags.StringP(addrKey, "a", defaultAddr, "listening address")
	_ = serveViper.BindPFlag(addrKey, flags.Lookup(addrKey))

	flags.StringP(dbNameKey, "d", defaultDBName, "data store location")
	_ = serveViper.BindPFlag(dbNameKey, flags.Lookup(dbNameKey))
}

func resolveDBPath(v *viper.Viper) (dbPath string, err error) {
	dbPath = v.GetString(dbNameKey)
	if dbPath == defaultDBName {
		dbPath, err = xdg.DataFile(relDBName)
		if err != nil {
			return "", err
		}
	}
	return dbPath, nil
}

func resolveUDSAddr(v *viper.Viper) (addr string, err error) {
	addr = v.GetString(addrKey)
	// return string unchanged if tcp is requested.
	if strings.HasPrefix(strings.ToLower(addr), "tcp") {
		return addr, nil
	}
	if addr == defaultAddr {
		addr, err = xdg.RuntimeFile(relUDS)
		if err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("file://%s", addr), nil
}
