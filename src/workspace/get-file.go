package main

import (
	"net/http"
	"fmt"
	"log"
	"io"
	"os"
)

func do_fetch_config(branch string) {
	info_run("Get configfile")

	url := fmt.Sprintf("https://raw.githubusercontent.com/mamemomonga/docker-workspaces/%s/config.yaml",branch)
	filename := "config.yaml"

	log.Printf("Fetch: %s", url)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatal(fmt.Sprintf("Status Code: %d", res.StatusCode))
	}

	defer res.Body.Close()
	errChan := make(chan error)
	defer close(errChan)

	go func() {
		log.Printf("Write: %s\n", filename)
		outfile, err := os.Create(filename)
		if err != nil {
			errChan <- err
			return
		}
		defer outfile.Close()
		if _, err := io.Copy(outfile, res.Body); err != nil {
			errChan <- err
			return
		}
		errChan <- nil
		return
	}()
	err = <-errChan
	if err != nil {
		log.Fatal(err)
	}
}
