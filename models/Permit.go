package models

// Permit - модель разрешения
type Permit struct{
	id string
	entity Ref // ссылка на сущность
	permitType Ref // ссылка на тип разрешения
	group Ref // ссылка на группу
}