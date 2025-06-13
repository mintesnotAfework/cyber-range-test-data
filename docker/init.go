package docker

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/docker/docker/client"
	"github.com/joho/godotenv"
)

type DockerManager struct {
	cli *client.Client
}

func NewDockerManager() (*DockerManager, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	// cli, err := client.NewClientWithOpts(
	// 	client.WithHost(os.Getenv("DOCKER_HOSTING_URL")),
	// 	client.WithAPIVersionNegotiation())
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to create Docker client: %v", err)
	// }

	cacertPath := "cert/root.pem"
	certPath := "cert/crt.pem"
	keyPath := "cert/key.pem"

	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, errors.New("failed to load client certificate")
	}

	caCert, err := os.ReadFile(cacertPath)
	if err != nil {
		return nil, errors.New("can not load the root CA")
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return nil, errors.New("can not create trusted CA lists")
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
		MinVersion:   tls.VersionTLS12,
		ServerName:   "docker.ctf.me",
	}

	cli, err := client.NewClientWithOpts(
		client.WithHost(os.Getenv("DOCKER_HOSTING_URL")),
		client.WithAPIVersionNegotiation(),
		client.WithHTTPClient(&http.Client{
			Transport: &http.Transport{
				TLSClientConfig: tlsConfig,
			},
		}),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create Docker client: %v", err)
	}

	return &DockerManager{cli: cli}, nil
}
