package main

import (
	ssh "code.google.com/p/go.crypto/ssh"
	"io/ioutil"
)

type Credentials struct {
	username string
	password string
}
type CredentialsChan chan *Credentials

func StartServer(ip string, keyFile string, credentialsChan CredentialsChan) error {
	// The PasswordCallback always returns false. The supplied credentials get
	// passed to the specified callback for further processing.
	config := &ssh.ServerConfig{
		PasswordCallback: func(conn *ssh.ServerConn, user, pass string) bool {
			credentialsChan <- &Credentials{user, pass}
			return false
		},
	}

	// Read the server's private key
	pemBytes, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return err
	}
	if err = config.SetRSAPrivateKey(pemBytes); err != nil {
		return err
	}

	// set up listener
	listener, err := ssh.Listen("tcp", ip, config)
	if err != nil {
		return err
	}

	go processConnections(listener)

	return nil
}

func processConnections(listener *ssh.Listener) {
	for {
		// accept new connection
		sConn, err := listener.Accept()
		if err != nil {
			continue
		}

		// dispatch handshake and discard result (will always be false)
		go sConn.Handshake()
	}
}
