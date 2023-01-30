package models

// Property - модель свойства сущности
type Property struct{
	id string
	name string
	systemName string // системное наименование, name - на английском
	entity Ref // ссылка на сущность
	dataType Ref // ссылка на тип данных
	propType Ref // ссылка на тип свойства
	createDate string
	createUser string
	updateDate string
	updateUser string
}