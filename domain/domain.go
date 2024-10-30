package domain

import "time"

type User struct {
	Id       string `bson:"_id"`
	UserName string `bson:"userName"`
	GuildID  string `bson:"guildId"`
}

type Goal struct {
	Id          string    `bson:"_id"`
	CreatedAt   time.Time `bson:"createdAt"`
	Goal        string    `bson:"goal"`
	GoalType    string    `bson:"goalType"`
	SetSchedule bool      `bson:"setSchedule"`
}
