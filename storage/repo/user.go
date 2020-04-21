package repo

import (
	"errors"

	"github.com/golang/protobuf/ptypes/wrappers"
)

//User ...
type User struct {
	ID           string                `json:"id,omitempty"`
	CreatedAt    string                `json:"created_at,omitempty"`
	UpdatedAt    string                `json:"updated_at,omitempty"`
	DeletedAt    string                `json:"deleted_at,omitempty"`
	Mail         *wrappers.StringValue `json:"mail,omitempty"`
	Password     string                `json:"password,omitempty"`
	AccessToken  string                `json:"access_token,omitempty"`
	RefreshToken string                `json:"refresh_token,omitempty"`
	UserTypeID   int32                 `json:"user_type_id,omitempty"`
	IsVerified   bool                  `json:"is_verified,omitempty"`
}

var (
	//ErrAlreadyExists ...
	ErrAlreadyExists = errors.New("Already exists")
)

//UserStorageI ...
type UserStorageI interface {
	Create(*User) error
	Update(*User) error
	GetUser(id string) (*User, error)
	GetAllUsers() ([]*User, error)
	Delete(id string) error
	CheckMail(mail string) (bool, error)
	GetUserByMail(mail string) (*User, error)
	UpdateUserTokens(id, accessToken, refreshToken string) error
	CreateConnection(userID, socialID, accessToken, connType string) error
	IsConnectionExists(socialID, connType string) (bool, error)
	IsConnectionExistsFromUserID(userID, connType string) (bool, error)
	GetUserIDFromExistingConnection(socialID, accessToken, connType string) (string, error)
	DeleteUserConnections(userID string) error
	SetMail(userID, mail string) error
	SetPassword(id, password string) error
}
