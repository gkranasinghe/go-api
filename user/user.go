package user

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
	Rank int64  `json:"rank"`
}
