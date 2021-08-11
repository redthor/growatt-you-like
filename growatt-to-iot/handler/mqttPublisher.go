package handler

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTPublisher struct {
	topic   string
	options *mqtt.ClientOptions
	client  mqtt.Client
}

func NewMQTTPublisher(tlsCert string, tlsPrivateKey string, tlsCA string, mqttEndpoint string) *MQTTPublisher {
	tlsConfig, certHash, err := getTLSConfig(tlsCert, tlsPrivateKey, tlsCA)
	if err != nil {
		log.Fatalf("Received an error extracting certificate information: %v", err.Error())
	}
	topic := fmt.Sprintf("%s/inbound-raw", certHash)
	log.Printf("Topic: %s", topic)

	opts := mqtt.NewClientOptions()
	port := 8883
	opts.AddBroker(fmt.Sprintf("tls://%s:%d", mqttEndpoint, port))
	// ClientId is following Terraform thing name
	opts.SetClientID("growatt-to-iot").SetTLSConfig(tlsConfig)

	return &MQTTPublisher{
		topic:   topic,
		options: opts,
	}
}

func (m *MQTTPublisher) Register() func(message []byte) {
	m.client = mqtt.NewClient(m.options)
	log.Println("Start MQTT Connect")

	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("failed to create connection: %v", token.Error())
	}

	log.Println("Done MQTT Connect")

	return m.Publish
}

func (m *MQTTPublisher) Publish(message []byte) {
	log.Printf("Publish %d bytes\n", len(message))

	if token := m.client.Publish(m.topic, 0, false, message); token.Wait() && token.Error() != nil {
		log.Fatalf("failed to send update: %v", token.Error())
	}
}

func getTLSConfig(tlsCert string, tlsPrivateKey string, tlsCA string) (*tls.Config, string, error) {
	log.Println("Loading certificates.")
	certPool := x509.NewCertPool()
	ca, err := os.ReadFile(tlsCA)
	if err != nil {
		return &tls.Config{}, "", err
	}
	certPool.AppendCertsFromPEM(ca)

	cert, err := tls.LoadX509KeyPair(tlsCert, tlsPrivateKey)
	if err != nil {
		return &tls.Config{}, "", err
	}

	x509Cert, err := x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		return &tls.Config{}, "", err
	}

	tlsConfig := tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certPool,
		// ClientAuth = whether to request cert from server.
		// Since the server is set up for SSL, this happens
		// anyways.
		ClientAuth: tls.NoClientCert,
		// ClientCAs = certs used to validate client cert.
		ClientCAs: nil,
		// Certificates = list of certs client sends to server.
		Certificates: []tls.Certificate{cert},
	}

	return &tlsConfig, hex.EncodeToString(x509Cert.SubjectKeyId), nil
}
