// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package data

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Employee struct {
	ID        pgtype.UUID `json:"id"`
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
}
