package managers

import "github.com/dev-jiemu/itops-assignment/backend/pkg/models"

type UserList struct {
	Users []*models.User
}

var UserManager *UserList

func NewUserManager() *UserList {
	return &UserList{
		Users: initUsers(),
	}
}

// initUsers : 시스템에 미리 정의된 사용자 (main func 돌때 설정)
func initUsers() []*models.User {
	var initModel = make([]*models.User, 0, 3)

	initModel = append(initModel, &models.User{ID: 1, Name: "김개발"})
	initModel = append(initModel, &models.User{ID: 2, Name: "이디자인"})
	initModel = append(initModel, &models.User{ID: 3, Name: "박기획"})

	return initModel
}

// TODO : user CRUD

func (v *UserList) GetUserInfo(userId uint) *models.User {
	for _, user := range v.Users {
		if user.ID == userId {
			return user
		}
	}

	return &models.User{}
}
