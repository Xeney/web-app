package logic

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	// _ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

type Course struct {
	Id          int
	Role        string
	Lvl         string
	Description string
	Block_1     string
	Block_2     string
	Block_3     string
	Block_4     string
	Block_5     string
	Name        string
	Level_skill int
}

// Получить пользователя по login
func GetToId_sql(login string) (User, error) {
	db, err := sql.Open("sqlite", "db.sqlite")
	if err != nil {
		return User{}, errors.New("error connection")
	}
	defer db.Close()
	row := db.QueryRow("select * from users where login = $1", login)
	u := User{}
	var favorites_ []byte
	err = row.Scan(&u.Id, &u.Image, &u.Login, &u.Password, &u.Skill, &u.Role1, &u.Role2, &favorites_)
	if err != nil {
		return User{}, errors.New("error account")
	}
	err = json.Unmarshal(favorites_, &u.Favorites)
	if err != nil {
		return User{}, errors.New("error unmarshal")
	}
	return u, nil
}

// Создать пользователя
func Create_sql(avatar, login, password, role_1, role_2 string, favByte map[string]FavCourses, skill float64) (sql.Result, error) {
	db, err := sql.Open("sqlite", "db.sqlite")
	if err != nil {
		return nil, errors.New("error connection")
	}
	defer db.Close()
	resFavByte, err := json.Marshal(favByte)
	if err != nil {
		return nil, errors.New("error marshal")
	}
	result, err := db.Exec("insert into users (image, login, password, skill, role1, role2, favorites) values ($1, $2, $3, $4, $5, $6, $7)", avatar, login, password, skill, role_1, role_2, resFavByte)
	if err != nil {
		return nil, errors.New("error sql_exec")
	}
	return result, nil
}

// Получить все курсы
func GetAllCourse_sql() ([]Course, error) {
	db, err := sql.Open("sqlite", "db.sqlite")
	if err != nil {
		fmt.Println(err)
		return []Course{}, errors.New("error connection")
	}
	rows, err := db.Query("select * from courses")
	if err != nil {
		return []Course{}, errors.New("error sql_query")
	}
	defer rows.Close()
	var courses []Course
	for rows.Next() {
		c := Course{}
		err := rows.Scan(&c.Id, &c.Role, &c.Lvl, &c.Description, &c.Block_1, &c.Block_2, &c.Block_3, &c.Block_4, &c.Block_5, &c.Name, &c.Level_skill)
		if err != nil {
			return []Course{}, errors.New("error sql_scan")
		}
		courses = append(courses, c)
	}
	return courses, nil
}

// Создать курс
func CreateCourse_sql(Role, Lvl, Description, Block_1, Block_2, Block_3, Block_4, Block_5, Name, Level_skill string) (sql.Result, error) {
	db, err := sql.Open("sqlite", "db.sqlite")
	if err != nil {
		return nil, errors.New("error connection")
	}
	defer db.Close()
	result, err := db.Exec("insert into courses (role, lvl, description, block1, block2, block3, block4, block5, name, lvl_skill) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)", Role, Lvl, Description, Block_1, Block_2, Block_3, Block_4, Block_5, Name, Level_skill)
	if err != nil {
		return nil, errors.New("error sql_exec")
	}
	return result, nil
}

// Получить курс по ID
func GetToCourseId_sql(id int) (Course, error) {
	db, err := sql.Open("sqlite", "db.sqlite")
	if err != nil {
		return Course{}, errors.New("error connection")
	}
	defer db.Close()
	row := db.QueryRow("select * from courses where id = $1", id)
	c := Course{}
	err = row.Scan(&c.Id, &c.Role, &c.Lvl, &c.Description, &c.Block_1, &c.Block_2, &c.Block_3, &c.Block_4, &c.Block_5, &c.Name, &c.Level_skill)
	if err != nil {
		return Course{}, errors.New("error scan")
	}
	return c, nil
}

// Добавление избранных курсов
func UpdateFavorites_sql(id int, login_user string) error {
	db, err := sql.Open("sqlite", "db.sqlite")
	if err != nil {
		return errors.New("error connection")
	}
	defer db.Close()
	new_course, err := GetToCourseId_sql(id)
	if err != nil {
		return err
	}
	user, err := GetUser(login_user)
	if err != nil {
		return err
	}
	if (user.Favorites[strconv.Itoa(new_course.Id)] != FavCourses{}) {
		return errors.New("already in favorites")
	}
	user.Favorites[strconv.Itoa(new_course.Id)] = FavCourses{new_course.Name, strconv.Itoa(new_course.Id)}
	res_m, err := json.Marshal(user.Favorites)
	if err != nil {
		return errors.New("error marshal")
	}
	_, err = db.Exec("update users set favorites = $1 where login = $2", res_m, login_user)
	if err != nil {
		return errors.New("error exec")
	}
	return nil
}

// Удаление курсов из избранных
func DeleteFavorites_sql(id int, login_user string) error {
	db, err := sql.Open("sqlite", "db.sqlite")
	if err != nil {
		return errors.New("error connection")
	}
	defer db.Close()
	new_course, err := GetToCourseId_sql(id)
	if err != nil {
		return err
	}
	user, err := GetUser(login_user)
	if err != nil {
		return err
	}
	if (user.Favorites[strconv.Itoa(new_course.Id)] != FavCourses{}) {
		delete(user.Favorites, strconv.Itoa(new_course.Id))
	} else {
		return errors.New("already in favorites")
	}
	res_m, err := json.Marshal(user.Favorites)
	if err != nil {
		return errors.New("error marshal")
	}
	_, err = db.Exec("update users set favorites = $1 where login = $2", res_m, login_user)
	if err != nil {
		return errors.New("error exec")
	}
	return nil
}

// Изменение skill пользователя
func UpdateSkill_sql(skill int, login_user string) error {
	db, err := sql.Open("sqlite", "db.sqlite")
	if err != nil {
		return errors.New("error connection")
	}
	defer db.Close()
	_, err = db.Exec("update users set skill = $1 where login = $2", skill, login_user)
	if err != nil {
		return errors.New("error exec")
	}
	return nil
}

// // Создать блюдо
// func Create_sql(dd_name, product_name string) (sql.Result, error) {
// 	db, err := sql.Open("sqlite3", "db.sqlite")
// 	if err != nil {
// 		return nil, errors.New("error connection")
// 	}
// 	defer db.Close()
// 	var product = Product{
// 		product_name,
// 		[]Ingredient_Product{},
// 	}
// 	res, err := json.Marshal(product)
// 	if err != nil {
// 		return nil, errors.New("error marshal")
// 	}
// 	var result sql.Result
// 	switch dd_name {
// 	case "dish":
// 		result, err = db.Exec("insert into dish (name, ingredients) values ($1, $2)", product_name, res)
// 	case "drink":
// 		result, err = db.Exec("insert into drink (name, ingredients) values ($1, $2)", product_name, res)
// 	default:
// 		return nil, errors.New("error dd_name")
// 	}
// 	if err != nil {
// 		return nil, errors.New("error sql_exec")
// 	}
// 	return result, nil
// }

// // Получить всё
// func Get_sql(dd_name string) ([]Product_id, error) {
// 	db, err := sql.Open("sqlite3", "db.sqlite")
// 	if err != nil {
// 		return nil, errors.New("error connection")
// 	}
// 	defer db.Close()
// 	products := []Product_id{}
// 	var rows *sql.Rows
// 	switch dd_name {
// 	case "dish":
// 		rows, err = db.Query("select * from dish")
// 	case "drink":
// 		rows, err = db.Query("select * from drink")
// 	default:
// 		return nil, errors.New("error dd_name")
// 	}
// 	defer rows.Close()
// 	if err != nil {
// 		return nil, errors.New("error sql_query")
// 	}
// 	for rows.Next() {
// 		p := Product{}
// 		p_id := Product_id{}
// 		ingrs, _ := json.Marshal(Product{})
// 		err := rows.Scan(&p_id.ID, &p_id.Name, &ingrs)
// 		json.Unmarshal(ingrs, &p)
// 		p_id.Ingredients = p.Ingredients
// 		if err != nil {
// 			fmt.Println(err)
// 			continue
// 		}
// 		products = append(products, p_id)
// 	}
// 	return products, nil
// }

// // Получить по ID
// func GetToId_sql(dd_name string, id uint8) (Product_id, error) {
// 	db, err := sql.Open("sqlite3", "db.sqlite")
// 	if err != nil {
// 		return Product_id{}, errors.New("error connection")
// 	}
// 	defer db.Close()
// 	var row *sql.Row
// 	switch dd_name {
// 	case "dish":
// 		row = db.QueryRow("select * from dish where id = $1", id)
// 	case "drink":
// 		row = db.QueryRow("select * from drink where id = $1", id)
// 	default:
// 		return Product_id{}, errors.New("error dd_name")
// 	}
// 	p := Product{}
// 	p_id := Product_id{}
// 	ingrs, _ := json.Marshal(Product{})
// 	err = row.Scan(&p_id.ID, &p_id.Name, &ingrs)
// 	json.Unmarshal(ingrs, &p)
// 	p_id.Ingredients = p.Ingredients
// 	if err != nil {
// 		return Product_id{}, errors.New("error scan")
// 	}
// 	return p_id, nil
// }

// // Обновление ингредиентов
// func Update_sql(dd_name string, id uint8, ingrs []Ingredient_Product) error {
// 	db, err := sql.Open("sqlite3", "db.sqlite")
// 	if err != nil {
// 		return errors.New("error connection")
// 	}
// 	defer db.Close()
// 	prod, err := GetToId_sql(dd_name, id)
// 	if err != nil {
// 		return errors.New("error get_to_id")
// 	}
// 	var product = Product{
// 		prod.Name,
// 		ingrs,
// 	}
// 	res_m, err := json.Marshal(product)
// 	if err != nil {
// 		return errors.New("error marshal")
// 	}
// 	// var result sql.Result
// 	switch dd_name {
// 	case "dish":
// 		_, err = db.Exec("update dish set ingredients = $1 where id = $2", res_m, id)
// 	case "drink":
// 		_, err = db.Exec("update drink set ingredients = $1 where id = $2", res_m, id)
// 	default:
// 		return errors.New("error dd_name")
// 	}
// 	if err != nil {
// 		return errors.New("error exec")
// 	}
// 	return nil
// }

// // Удаление по id
// func Delete_sql(dd_name string, id uint8) error {
// 	// result, err := db.Exec("delete from Products where id = $1", 1)
// 	db, err := sql.Open("sqlite3", "db.sqlite")
// 	if err != nil {
// 		return errors.New("error connection")
// 	}
// 	defer db.Close()
// 	// var result sql.Result
// 	switch dd_name {
// 	case "dish":
// 		_, err = db.Exec("delete from dish where id = $1", id)
// 	case "drink":
// 		_, err = db.Exec("delete from drink where id = $1", id)
// 	default:
// 		return errors.New("error dd_name")
// 	}
// 	if err != nil {
// 		return errors.New("error exec")
// 	}
// 	return nil
// }
