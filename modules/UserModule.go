package modules
import(
	"github.com/ArtemOlchev/goRestApi/models"
	// "time"
)

// Получить пользователей
func GetUsers() (us[]models.User, err error){

	return
}

// Получить пользователя
func GetUser(id string) (u *models.User, err error){

	return
}

// Создать пользователя
func CreateUser(u models.User) (outU *models.User, err error){

	return
}

// Сохранить пользователя
func SaveUser(u models.User) (outU *models.User, err error){

	return
}

// Удалить пользователя
func DeleteUser(id string) (u *models.User, err error){

	return
}

// Изменить пароль пользователя
func ChangeUserPassword(oldPsw, newPsw string) (err error){
	
	return
}
