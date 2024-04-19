package cmd

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog/log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

const MQTT_HOST string = "mqtt-host"
const MQTT_PORT string = "mqtt-port"

// streamCmd represents the stream command
func streamCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stream",
		Short: "short",
		Long:  `long`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var host string
			err := ReadCLIArg(cmd, MQTT_HOST, &host)
			if err != nil {
				log.Err(err).Msgf("Failed to obtain host")
				return err
			}

			var port int
			err = ReadCLIArgI(cmd, MQTT_PORT, &port)
			if err != nil {
				log.Err(err).Msgf("Failed to obtain port")
				return err
			}

			log.Info().Msgf("Connecting to mqtt://%s:%d", host, port)

			var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
				fmt.Printf("TOPIC: %s\n", msg.Topic())
				fmt.Printf("MSG: %s\n", msg.Payload())
			}

			mqtt.DEBUG = &log.Logger
			mqtt.ERROR = &log.Logger

			opts := mqtt.NewClientOptions().
				AddBroker(fmt.Sprintf("tcp://%s:%d", host, port)).
				SetClientID(APP_NAME)

			opts.SetKeepAlive(2 * time.Second)
			opts.SetDefaultPublishHandler(f)
			opts.SetPingTimeout(1 * time.Second)

			c := mqtt.NewClient(opts)
			if token := c.Connect(); token.Wait() && token.Error() != nil {
				panic(token.Error())
			}

			if token := c.Subscribe("94:05:BB:80:40:69", 0, f); token.Wait() && token.Error() != nil {
				fmt.Println(token.Error())
				os.Exit(1)
			}

			return nil
		},
	}

	var mqttHost string
	var mqttPort int

	// 192.168.178.29
	cmd.PersistentFlags().StringVar(&mqttHost, MQTT_HOST, "", "The MQTT host to connect to")
	cmd.PersistentFlags().IntVar(&mqttPort, MQTT_PORT, 1883, "The MQTT port to connect to")

	return cmd
}
