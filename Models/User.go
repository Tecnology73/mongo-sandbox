package Models

type User struct {
	BaseModel `bson:",inline"`

	Name          string `json:"name" bson:"name"`
	Email         string `json:"email" bson:"email"`
	LoginAttempts int16  `json:"loginAttempts" bson:"loginAttempts"`
	IsLocked      bool   `json:"isLocked" bson:"isLocked"`
}
