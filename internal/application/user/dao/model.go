package dao

type User struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`
}
