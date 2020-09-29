package Model

import "LearnGoGin/app/Util"

func FindById(id int) (User, error) {
	con := Util.GetDatabaseConnection()

	user := User{ID: id}

	err := con.Model(&user).WherePK().Select()

	if err != nil {
		return user, err
	}

	return user, nil
}
