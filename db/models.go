package db

import "time"

type User struct {
	ID                 int
	AliexpressLogin    string
	AliexpressPassword string
}

type Index struct {
	ID               int
	Index            string
	URL              string
	Location         string
	LastModification time.Time
}
