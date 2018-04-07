package models

type Connection struct {
	Server string `bson:"Server" json:"Server"`

	User string `bson:"User" json:"User"`

	Database string `bson:"Database" json:"Database"`

	Password string `bson:"Password" json:"Password"`
}
