package users

type Order struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
	BookID int64 `json:"book_id"`
}
type Orders []*Order
