package models

// Object - модель объекта сущности
type Object struct{
	id string
	name string
	entity Ref // ссылка на сущность
	createDate string
	createUser string
	updateDate string
	updateUser string
}