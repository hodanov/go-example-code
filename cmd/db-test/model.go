package main

import "time"

// User is a struct that is mapped to users table.
type User struct {
	ID      int
	Name    string
	Age     int
	Created time.Time
}

// Snippet is a struct that is mapped to snippets table.
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
}
