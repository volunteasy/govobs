package user

import (
	"govobs/core/types"
)

type (
	User struct {
		ID       types.ID
		Name     string
		Nickname types.Nickname
		Document types.Document
		Contact  types.Contact
	}

	Filter struct {
		Role           string
		OrganizationID types.ID
		Identification string

		types.Filter
	}
)