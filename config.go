package main

type config struct {
	endpoint string
	limit    int
	page     int // starts at page -1 (no page)
}
