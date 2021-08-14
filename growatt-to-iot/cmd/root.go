package cmd

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"os"

	"github.com/redthor/growatt-you-like/growatt-to-iot/handler"
	"github.com/redthor/growatt-you-like/growatt-to-iot/sumolog"
	"github.com/redthor/growatt-you-like/growatt-to-iot/util"
	"github.com/spf13/cobra"
)

var (
	ip                 net.IP
	port               int
	growattServerProxy bool

	tlsCert       string
	tlsPrivateKey string
	tlsCA         string

	mqttEndpoint string

	sumoLogicEndpoint string

	rootCmd = &cobra.Command{
		Use:   "growatt-to-iot",
		Short: "Sends messages to AWS IOT",
		Run:   growattToIOT,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func growattToIOT(cmd *cobra.Command, args []string) {
	configLog()

	log.Println("Starting growatt-to-iot.")
	printOtions()
	mqttPublisher := handler.NewMQTTPublisher(tlsCert, tlsPrivateKey, tlsCA, mqttEndpoint)

	chain := util.NewChain()
	onMessage := chain.AddFn(mqttPublisher.Register())

	// If we are proxying calls. This could be done from the cloud?
	if growattServerProxy {
		onMessage = chain.AddFn(handler.NewGrowattProxy().Register())
	}

	packetLength := 4096
	listenHandler := handler.NewListenHandler(ip, port, packetLength, onMessage)
	listenHandler.Start()
}

func configLog() {
	// Send logs to stdOut and SumoLogic if available
	slURL, err := url.Parse(sumoLogicEndpoint)
	if err == nil && len(sumoLogicEndpoint) > 0 {
		sumo, err := sumolog.NewSumoLogic(slURL)

		if err == nil {
			multi := io.MultiWriter(os.Stdout, sumo)
			log.SetOutput(multi)
		}
	}
}

func printOtions() {
	log.Printf("Listen on IP = %v, Port = %v", ip, port)

	proxyMessage := "Will proxy to Growatt Server"
	if !growattServerProxy {
		proxyMessage = "Will not proxy to Growatt Server"
	}
	log.Println(proxyMessage)

	log.Printf("TLS Cert = %v, TLS Private Key = %v, TLS CA = %v", tlsCert, tlsPrivateKey, tlsCA)
	log.Printf("mqtt endpoint = %v", mqttEndpoint)
}

func init() {
	// Default is promiscuous
	defaultIP := net.IPv4(0, 0, 0, 0)
	rootCmd.Flags().IPVar(&ip, "ip", defaultIP, "IP Address to listen to Growatt messages.")

	// Default to Growatt port
	defaultPort := 5279
	rootCmd.Flags().IntVar(&port, "port", defaultPort, "The port to listen to Growatt messages.")

	defaultGrowattServerProxy := true
	helpMsgGrowattServerProxy := "Forward messages to Growatt Server. --growattServerProxy=0 to turn off"
	rootCmd.Flags().BoolVar(&growattServerProxy, "growattServerProxy", defaultGrowattServerProxy, helpMsgGrowattServerProxy)

	defaultTLSCert := "growatt-to-iot.cert.pem"
	rootCmd.Flags().StringVar(&tlsCert, "tlsCert", defaultTLSCert, "TLS certificate.")

	defaultTLSPrivateKey := "growatt-to-iot.private.key"
	rootCmd.Flags().StringVar(&tlsPrivateKey, "tlsPrivateKey", defaultTLSPrivateKey, "TLS private key.")

	defaultTLSCA := "AmazonRootCA1.pem"
	rootCmd.Flags().StringVar(&tlsCA, "tlsCA", defaultTLSCA, "TLS CA.")

	rootCmd.Flags().StringVar(&mqttEndpoint, "mqttEndpoint", "", "mqtt endpoint. (required)")
	err := rootCmd.MarkFlagRequired("mqttEndpoint")
	if err != nil {
		log.Fatalf("Error marking flag required: %s", err.Error())
	}

	rootCmd.Flags().StringVar(&sumoLogicEndpoint, "sumoLogicEndpoint", "", "sumo logic http endpoint.")
}
