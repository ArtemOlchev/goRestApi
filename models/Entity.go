/*
models - пакет содержащий модели системы.
Модели:
	Entity - сущность
	Property - свойство
	DataType - тип данных
	PropertyType - тип свойства
	Object - объект сущности
	User - пользователь
	Group - группа пользователей
	UserGroup - объект соотношения пользователь-группа
	PermitType - тип доступа
	Permit - доступ группы к сущности
	Ref - ссылка на модель
*/
package models

var dateTimeLayout string = "02.01.2006 15:04:05"

// Entity - модель сущности
type Entity struct{
	id string
	name string
	systemName string // системное наименование, name - на английском
	createDate string
	createUser string
	updateDate string
	updateUser string
}