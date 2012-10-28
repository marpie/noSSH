package main

import (
	"fmt"
	"os"
  "time"
)

const HOST = "0.0.0.0:2222"
const LOGFILE = "nossh.log"
const SEPERATOR = "|||"

func main() {
	credsChan := make(chan *Credentials, 10)

	// start credentials processor
	go logCredentials(credsChan)

	// start SSH Server
	fmt.Printf("[*] Starting server on %s\n", HOST)
	if err := StartServer(HOST, "id_rsa", credsChan); err != nil {
		fmt.Printf("[E] Error while starting Server: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("[*] Waiting for connections...")
  for {
    time.Sleep(2*time.Second)
  }
	return
}

func logCredentials(credentialsChan CredentialsChan) {
  fmtStr := "%s: %s" + SEPERATOR + "%s\n"
	for {
		creds := <-credentialsChan
		fmt.Printf(fmtStr, time.Now().String(), creds.username, creds.password)
	}
	return
}
