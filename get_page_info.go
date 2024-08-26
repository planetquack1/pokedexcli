package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Page struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous string   `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func getPageInfo(cfg *config, url string) Page {

	// get bytes from cache, if in cache
	body, ok := cfg.cache.Get(url)

	// if not in cache, HTTP
	if !ok {
		// GET
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		body, err = io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	// Unmarshal
	page := Page{}
	err := json.Unmarshal(body, &page)
	if err != nil {
		fmt.Println(err)
	}

	return page
}
