// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlmodels

import (
	"fmt"
	"io"
	"strconv"
)

// Input to define permissions of a role
type PermissionInput struct {
	Code  string `json:"code"`
	Level string `json:"level"`
}

// Input to create or update a quote
type QuoteInput struct {
	ID      *int   `json:"id"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

// Input to create or update a role
type RoleInput struct {
	ID          *int               `json:"id"`
	Name        string             `json:"name"`
	Permissions []*PermissionInput `json:"permissions"`
	Users       []int              `json:"users"`
}

// Input to update the sort order of an entity
type SortOrderInput struct {
	ID       int `json:"id"`
	Position int `json:"position"`
}

// UploadResult is returned when a file upload succeeded
type UploadResult struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
}

// Input to create or update a user
type UserInput struct {
	ID             *int   `json:"id"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"password_repeat"`
	IsSuperuser    bool   `json:"is_superuser"`
	Roles          []int  `json:"roles"`
}

// Possible sort directions
type SortDirection string

const (
	SortDirectionAsc  SortDirection = "ASC"
	SortDirectionDesc SortDirection = "DESC"
)

var AllSortDirection = []SortDirection{
	SortDirectionAsc,
	SortDirectionDesc,
}

func (e SortDirection) IsValid() bool {
	switch e {
	case SortDirectionAsc, SortDirectionDesc:
		return true
	}
	return false
}

func (e SortDirection) String() string {
	return string(e)
}

func (e *SortDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortDirection", str)
	}
	return nil
}

func (e SortDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
