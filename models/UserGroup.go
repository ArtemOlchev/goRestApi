package models

// UserGroup - модель принадлежности пользователя группе
type UserGroup struct{
	id string
	name string
	group Ref
	user Ref
}