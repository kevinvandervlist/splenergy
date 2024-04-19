package cmd

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

const APP_NAME = "splenergy"
const LOG_LEVEL_FLAG = "log-level"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   APP_NAME,
	Short: "todo",
	Long:  `multiline todo`,
}

func Execute() {
	var logLevel string
	// Any persistent flags
	rootCmd.PersistentFlags().StringVar(&logLevel, LOG_LEVEL_FLAG, "info", "Set the log level (panic, fatal, error, warn, info, debug, trace)")

	// Register our custom commands
	rootCmd.AddCommand(streamCmd())

	// Prerun hook for configuring the log level properly.
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		// Note: don't use the global log level as it _overrides_ the local log level settings.
		lvl := zerolog.GlobalLevel()
		switch logLevel {
		case "panic":
			lvl = zerolog.PanicLevel
		case "fatal":
			lvl = zerolog.FatalLevel
		case "error":
			lvl = zerolog.ErrorLevel
		case "warn":
			lvl = zerolog.WarnLevel
		case "info":
			lvl = zerolog.InfoLevel
		case "debug":
			lvl = zerolog.DebugLevel
		case "trace":
			lvl = zerolog.TraceLevel
		default:
			lvl = zerolog.InfoLevel
		}

		// Set de default logger to stdout by default
		log.Logger = zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339,
		}).With().Timestamp().Logger().Level(lvl)

		log.Debug().Msgf("Log level is set to: %s", lvl.String())
		return nil
	}

	// Finally execute
	err := rootCmd.Execute()
	if err != nil {
		log.Err(err).Msgf("Failed to execute mingvar")
		os.Exit(1)
	}
}
