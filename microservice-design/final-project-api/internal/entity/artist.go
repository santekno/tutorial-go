package entity

// The entity will be used for artist definition
type Artist struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Rating int    `json:"rating"`
}
