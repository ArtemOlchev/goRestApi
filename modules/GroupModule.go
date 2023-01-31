package modules
import(
	"github.com/ArtemOlchev/goRestApi/models"
	// "time"
)

// Получить группы
func GetGroups() (gs []models.Group, err error){

	return
}

// Получить группу
func GetGroup(id string) (g *models.Group, err error){

	return
}

// Создать группу
func CreateGroup(g models.Group) (outG *models.Group, err error){

	return
}

// Сохранить группу
func SaveGroup(g models.Group) (outG *models.Group, err error){

	return
}

// Удалить группу
func DeleteGroup(id string) (g *models.Group, err error){

	return
}

func AddUserToGroup(gId, uId string) (err error){

	return
}