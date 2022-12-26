package entity

// The entity will be used for album definition
type Album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Price  float32 `json:"price"`
	Artist Artist  `json:"artist"`
}
