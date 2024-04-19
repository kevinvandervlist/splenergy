package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func ReadCLIArg(cmd *cobra.Command, name string, value *string) error {
	var err error
	*value, err = cmd.Flags().GetString(name)
	if err != nil {
		log.Err(err).Msgf("Unable to acquire the %s", *value)
	} else {
		log.Debug().Msgf("User provided %s: %s", name, *value)
	}
	return err
}

func ReadCLIArgI(cmd *cobra.Command, name string, value *int) error {
	var err error
	*value, err = cmd.Flags().GetInt(name)
	if err != nil {
		log.Err(err).Msgf("Unable to acquire the %d", *value)
	} else {
		log.Debug().Msgf("User provided %s: %d", name, *value)
	}
	return err
}
