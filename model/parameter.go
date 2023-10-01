package model

type Parameter struct {
	Id            string `gorm:"primaryKey;column:id"`
	ParameterName string `gorm:"not null;column:parameter_name"`
	Value         int    `gorm:"not null;column:value"`
}
