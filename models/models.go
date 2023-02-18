package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserUUID        uuid.UUID `json:"user_uuid"`
	UserEmail       string    `json:"user_email"`
	UserPhoneNumber string    `json:"user_phone_number"`
	UserPassword    string    `json:"user_password"`
	UserBirthday    string    `json:"user_birthday"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type AccessDetails struct {
	AccessUUID string `json:"access_uuid"`
	UserUUID   string `json:"user_uuid"`
}

type TokenDetails struct {
	UserUUID     uuid.UUID `json:"user_uuid"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	AccessUUID   uuid.UUID `json:"access_uuid"`
	RefreshUUID  uuid.UUID `json:"refresh_uuid"`
	AtExpires    int64     `json:"at_expires"`
	RtExpires    int64     `json:"rt_expires"`
}

// === Request ===
type RequestJoinUser struct {
	UserUUID        uuid.UUID `json:"user_uuid"`
	UserEmail       string    `json:"user_email"`
	UserPhoneNumber string    `json:"user_phone_number"`
	UserPassword    string    `json:"user_password"`
	UserBirthday    string    `json:"user_birthday"`
}

type RequestLoginUser struct {
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
}

// === Resoponse ===
type MetaData struct {
	StatusCode int    `json:"status_code"`
	Status     string `json:"status"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
}

type UserData struct {
	Token string `json:"token"`
}

type SuccessResponse struct {
	MetaData    `json:"meta"`
	ResponsedAt time.Time `json:"responsed_at"`
}

type SuccessLoginResponse struct {
	MetaData    `json:"meat"`
	UserData    `json:"data"`
	ResponsedAt time.Time `json:"responsed_at"`
}

type ErrorResponse struct {
	MetaData    `json:"meta"`
	ResponsedAt time.Time `json:"responsed_at"`
}
