package modules
import(
	"github.com/ArtemOlchev/goRestApi/models"
	// "time"
)

// Получить сущность
func GetEntity(id string) (e *models.Entity, err error){
	
	return
}

// Получить сущности
func GetEntities() (es []models.Entity, err error){
	
	return
}

// Создать сущность
func CreateEntity(e models.Entity) (outE *models.Entity, err error){
	
	return
}

// Сохранить сущность
func SaveEntity(e models.Entity) (outE *models.Entity, err error){

	return
}

// Удалить сущность
func DeleteEntity(id string) (e *models.Entity, err error){

	return
}

// Получить свойства сущности
func GetEntityProperties(id string) (ps []models.Property, err error){

	return
}

// Получить объекты сущности
/*
TODO: так как сущность имеет непредсказуемую структуру, получение объектов
через GO является неосуществимым на данный момент. Поэтому получение объектов
сущности будет реализовано через сервис написанный на JavaScript.
*/
func GetEntityObjects(id string) (){

	return
}

// Получить объект сущности
/*
TODO: так как сущность имеет непредсказуемую структуру, получение объектов
через GO является неосуществимым на данный момент. Поэтому получение объектов
сущности будет реализовано через сервис написанный на JavaScript.
*/
func GetEntityObjectById(entityId string, objId string) (){

	return
}

// Получить объект сущности по наименованию
/*
TODO: так как сущность имеет непредсказуемую структуру, получение объектов
через GO является неосуществимым на данный момент. Поэтому получение объектов
сущности будет реализовано через сервис написанный на JavaScript.
*/
func GetEntityObjectByName(entityId string, objName string) (){

	return
}