package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID       int64  `db:"id" json:"id"`
	Account  string `db:"account" json:"account"`
	Password  string `db:"password" json:"password"`
	Nickname  string `db:"nickname" json:"nickname"`
	Phone  string `db:"phone" json:"phone"`
	Expired  int64 `db:"expired" json:"expired"`
	IsDisabled  uint8 `db:"is_disabled" json:"is_disabled"`
	IsDeleted  uint8 `db:"is_deleted" json:"is_deleted"`
	Salt  string `db:"account" json:"account"`
	UpdateTime  *time.Time `db:"update_time" json:"update_time"`
	CreateTime  *time.Time `db:"create_time" json:"create_time"`
}

// TableName returns the database table name of a User.
func (c *User) TableName() string {
	return "user"
}

// PrimaryKey returns the primary key of a User.
func (c *User) PrimaryKey() string {
	return "id"
}

// SortBy returns the column name that
// should be used as a fallback for sorting a set of User.
func (c *User) SortBy() string {
	return "create_time"
}

// Scan binds mysql rows to this User.
func (c *User) Scan(rows *sql.Rows) error {
	c.CreateTime = new(time.Time)
	c.UpdateTime = new(time.Time)
	return rows.Scan(
		&c.ID,
		&c.Account,
		&c.Password,
		&c.Nickname,
		&c.Phone,
		&c.Expired,
		&c.IsDisabled,
		&c.IsDeleted,
		&c.Salt,
		&c.UpdateTime,
		&c.CreateTime,
	)
}

type Users []*User

// Scan binds mysql rows to this Users.
func (cs *Users) Scan(rows *sql.Rows) (err error) {
	cp := *cs
	for rows.Next() {
		c := new(User)
		if err = c.Scan(rows); err != nil {
			return
		}
		cp = append(cp, c)
	}

	if len(cp) == 0 {
		return sql.ErrNoRows
	}

	*cs = cp

	return rows.Err()
}
