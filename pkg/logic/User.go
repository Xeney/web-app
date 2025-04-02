package logic

import (
	"database/sql"
	"errors"
)

type FavCourses struct {
	Name string
	Href string
}
type User struct {
	Id        int
	Image     string
	Login     string
	Password  string
	Skill     int
	Role1     string
	Role2     string
	Favorites map[string]FavCourses
}

func AddUser(avatar, login, password, role_1, role_2 string, favByte map[string]FavCourses, skill float64) (sql.Result, error) {
	result, err := Create_sql(avatar, login, password, role_1, role_2, favByte, skill)
	if err != nil {
		return nil, errors.New("error sql")
	}
	return result, nil
}

func GetUser(login string) (User, error) {
	result, err := GetToId_sql(login)
	if err != nil {
		return User{}, errors.New("error sql")
	}
	return result, nil
}

func AddFavoritesCourse(id int, login_user string) error {
	err := UpdateFavorites_sql(id, login_user)
	if err != nil {
		return err
	}
	return nil
}

func UpdateSkill(skill int, login_user string) error {
	err := UpdateSkill_sql(skill, login_user)
	if err != nil {
		return err
	}
	return nil
}

// func ProdAllRead(dd_name string) ([]Product_id, error) {
// 	result, err := Get_sql(dd_name)
// 	if err != nil {
// 		return nil, errors.New("error sql")
// 	}
// 	return result, nil
// }

// func ProdIDRead(dd_name string, id uint8) (Product_id, error) {
// 	result, err := GetToId_sql(dd_name, id)
// 	if err != nil {
// 		return Product_id{}, errors.New("error sql")
// 	}
// 	return result, nil
// }

// func ProdUpdateIngrs(dd_name string, id uint8, ingr Ingredient_Product) error {
// 	product, err := GetToId_sql(dd_name, id)
// 	if err != nil {
// 		return errors.New("error sql")
// 	}
// 	product.Ingredients = append(product.Ingredients, ingr)
// 	err = Update_sql(dd_name, id, product.Ingredients)
// 	if err != nil {
// 		return errors.New("error sql")
// 	}
// 	return nil
// }

// func ProdDelete(dd_name string, id uint8) error {
// 	err := Delete_sql(dd_name, id)
// 	if err != nil {
// 		return errors.New("error sql")
// 	}
// 	return nil
// }

// func ProdCalculation(p_id Product_id) (int, error) {
// 	var final_calc float32
// 	for _, v := range p_id.Ingredients {
// 		ing := float32(v.Gram) * (float32(v.Calories) / 100.0)
// 		final_calc += ing
// 	}
// 	return int(final_calc), nil
// }
