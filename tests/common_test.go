package tests

import (
	"log"
	"testing"
)

func TestUserHistorySlices(t *testing.T) {
	history := []string{
		"h1", // remove
		"h2", // remove
		"h3", // remove
		"h4", // keep
		"h5", // keep
		"h6", // keep
	}
	counter := len(history) - 3
	correct := history[counter:]
	removing := history[:counter]
	log.Println(correct, removing)
}
