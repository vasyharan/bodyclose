package util

import (
	"io"
	"log"
)

func Close(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Printf("error closing io: %w", err)
	}
}
