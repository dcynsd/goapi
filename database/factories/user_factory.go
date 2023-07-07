package factories

import (
	"goapi/app/models"

	"github.com/bxcodec/faker/v3"
)

func MakeUsers(times int) []models.User {

	var objs []models.User

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := models.User{
			Username: faker.Username(),
			Password: "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
			Name:     faker.Name(),
			Avatar:   faker.URL(),
		}
		objs = append(objs, model)
	}

	return objs
}
