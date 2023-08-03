package main

import "github.com/jaevor/go-nanoid"

func generateShortId() string {
	shortIdGenerator, err := nanoid.Standard(10)

	if err != nil {
		panic(err)
	}

	shortId := shortIdGenerator()

	return shortId
}
