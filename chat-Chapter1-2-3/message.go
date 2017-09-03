package main

import (
	"time"
)

// message shows a message
type message struct {
	Name    string
	Message string
	When    time.Time
}
