package main

import (
	"app/pkg/logic"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var GlobalLoginUser string
var GlobalSkillUser int

func SessionCheck(c *gin.Context) {
	s := sessions.Default(c)
	if user := s.Get("User"); user != nil {
		if GlobalLoginUser == "" {
			GetExit(c)
			location := url.URL{Path: "/login"}
			c.Redirect(http.StatusFound, location.RequestURI())
			return
		}
		res, err := logic.GetUser(GlobalLoginUser)
		if err != nil {
			panic(err)
		}
		GlobalSkillUser = res.Skill
		c.Keys["User"] = res
		c.Keys["Avatar"] = Avatar
	}
}

func GetTest(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("User") == nil {
		location := url.URL{Path: "/login"}
		c.Redirect(http.StatusFound, location.RequestURI())
	}
	c.HTML(http.StatusOK, "test.html", c.Keys)
}

func PostTest(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("User") == nil {
		location := url.URL{Path: "/login"}
		c.Redirect(http.StatusFound, location.RequestURI())
	}
	if c.PostForm("what-1") != "value-1" {
		err := logic.UpdateSkill(1, GlobalLoginUser)
		if err != nil {
			panic(err)
		}
		c.Keys["err"] = "Неверный ответ, ваш skill понижен до 1"
		c.HTML(http.StatusInternalServerError, "test.html", c.Keys)
		c.Keys["err"] = nil
		return
	}
	if c.PostForm("what-2") != "value-1" {
		err := logic.UpdateSkill(1, GlobalLoginUser)
		if err != nil {
			panic(err)
		}
		c.Keys["err"] = "Неверный ответ, ваш skill понижен до 1"
		c.HTML(http.StatusInternalServerError, "test.html", c.Keys)
		c.Keys["err"] = nil
		return
	}
	if c.PostForm("what-3") != "value-1" {
		err := logic.UpdateSkill(1, GlobalLoginUser)
		if err != nil {
			panic(err)
		}
		c.Keys["err"] = "Неверный ответ, ваш skill понижен до 1"
		c.HTML(http.StatusInternalServerError, "test.html", c.Keys)
		c.Keys["err"] = nil
		return
	}
	if c.PostForm("what-4") != "value-1" {
		err := logic.UpdateSkill(1, GlobalLoginUser)
		if err != nil {
			panic(err)
		}
		c.Keys["err"] = "Неверный ответ, ваш skill понижен до 1"
		c.HTML(http.StatusInternalServerError, "test.html", c.Keys)
		c.Keys["err"] = nil
		return
	}
	if c.PostForm("what-5") != "value-1" {
		err := logic.UpdateSkill(1, GlobalLoginUser)
		if err != nil {
			panic(err)
		}
		c.Keys["err"] = "Неверный ответ, ваш skill понижен до 1"
		c.HTML(http.StatusInternalServerError, "test.html", c.Keys)
		c.Keys["err"] = nil
		return
	}
	if c.PostForm("what-6") != "value-1" {
		err := logic.UpdateSkill(1, GlobalLoginUser)
		if err != nil {
			panic(err)
		}
		c.Keys["err"] = "Неверный ответ, ваш skill понижен до 1"
		c.HTML(http.StatusInternalServerError, "test.html", c.Keys)
		c.Keys["err"] = nil
		return
	}
	if c.PostForm("what-7") != "value-1" {
		err := logic.UpdateSkill(1, GlobalLoginUser)
		if err != nil {
			panic(err)
		}
		c.Keys["err"] = "Неверный ответ, ваш skill понижен до 1"
		c.HTML(http.StatusInternalServerError, "test.html", c.Keys)
		c.Keys["err"] = nil
		return
	}
	if c.PostForm("what-8") != "value-1" {
		err := logic.UpdateSkill(1, GlobalLoginUser)
		if err != nil {
			panic(err)
		}
		c.Keys["err"] = "Неверный ответ, ваш skill понижен до 1"
		c.HTML(http.StatusInternalServerError, "test.html", c.Keys)
		c.Keys["err"] = nil
		return
	}
	if c.PostForm("what-9") != "value-1" {
		err := logic.UpdateSkill(1, GlobalLoginUser)
		if err != nil {
			panic(err)
		}
		c.Keys["err"] = "Неверный ответ, ваш skill понижен до 1"
		c.HTML(http.StatusInternalServerError, "test.html", c.Keys)
		c.Keys["err"] = nil
		return
	}
	err := logic.UpdateSkill(10, GlobalLoginUser)
	if err != nil {
		panic(err)
	}
	c.Keys["res"] = "Успешно! Ваш skill повышен до 10"
	c.HTML(http.StatusInternalServerError, "test.html", c.Keys)
	c.Keys["res"] = nil
}

func GetHome(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", c.Keys)
}

func GetCourses(c *gin.Context) {
	res, err := logic.GetAllCourse_sql()
	if err != nil {
		c.Keys["error"] = err
		c.HTML(http.StatusInternalServerError, "error.html", c.Keys)
		c.Keys["error"] = nil
		return
	}
	c.Keys["res"] = res
	c.HTML(http.StatusOK, "courses.html", c.Keys)
	c.Keys["res"] = nil
}

func GetCoursesTarget(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("User") == nil {
		location := url.URL{Path: "/login"}
		c.Redirect(http.StatusFound, location.RequestURI())
	}
	num, _ := strconv.Atoi(c.Param("id"))
	data, err := logic.GetToCourseId_sql(num)
	if GlobalSkillUser < data.Level_skill {
		c.Keys["error"] = "Ваш SKILL слишком низкий"
		c.HTML(http.StatusInternalServerError, "error.html", c.Keys)
		c.Keys["error"] = nil
		return
	}
	if err != nil {
		c.Keys["error"] = err
		c.HTML(http.StatusInternalServerError, "error.html", c.Keys)
		c.Keys["error"] = nil
		return
	}
	c.Keys["data"] = data
	c.HTML(http.StatusOK, "course-target.html", c.Keys)
	c.Keys["data"] = nil
}

func PostCoursesTarget(c *gin.Context) {
	id_course := c.Param("id")
	if c.Request.Method == http.MethodPost {
		new_id_course, _ := strconv.Atoi(id_course)
		err := logic.AddFavoritesCourse(new_id_course, GlobalLoginUser)
		if err != nil {
			c.Keys["error"] = err
			c.HTML(http.StatusInternalServerError, "error.html", c.Keys)
			c.Keys["error"] = nil
			return
		}
	}
	location := url.URL{Path: "/courses/" + id_course}
	c.Redirect(http.StatusFound, location.RequestURI())
}

func PostLk(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("User") == nil {
		c.Keys["error"] = "Session nil"
		c.HTML(http.StatusOK, "error.html", c.Keys)
		c.Keys["error"] = nil
		return
	}
	id, _ := strconv.Atoi(c.PostForm("id"))
	err := logic.DeleteFavorites_sql(id, GlobalLoginUser)
	if err != nil {
		c.Keys["error"] = err
		c.HTML(http.StatusInternalServerError, "error.html", c.Keys)
		c.Keys["error"] = nil
		return
	}
	location := url.URL{Path: "/lk"}
	c.Redirect(http.StatusFound, location.RequestURI())
}

func GetLk(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("User") != nil {
		c.HTML(http.StatusOK, "lk.html", c.Keys)
		return
	}
	location := url.URL{Path: "/login"}
	c.Redirect(http.StatusFound, location.RequestURI())
}

func GetExit(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("User") != nil {
		session.Set("User", nil)
		session.Save()
		GlobalLoginUser = "none"
		GlobalSkillUser = 0
		c.Keys["User"] = nil
	}
	location := url.URL{Path: "/login"}
	c.Redirect(http.StatusFound, location.RequestURI())
}

func GetLogin(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("User") != nil {
		location := url.URL{Path: "/lk"}
		c.Redirect(http.StatusFound, location.RequestURI())
	}
	c.HTML(http.StatusOK, "login.html", c.Keys)
}

func PostLogin(c *gin.Context) {
	session := sessions.Default(c)
	if c.Request.Method == http.MethodPost {
		user, err := logic.GetUser(c.PostForm("login"))
		if err != nil {
			c.HTML(http.StatusOK, "login.html", gin.H{"error": err})
			return
		}
		if c.PostForm("password") != user.Password {
			c.HTML(http.StatusOK, "login.html", gin.H{"error": "Неправильный пароль"})
			return
		}
		if session.Get("User") == nil {
			session.Set("User", user)
			GlobalLoginUser = user.Login
			GlobalSkillUser = int(user.Skill)
			err := session.Save()
			if err != nil {
				fmt.Println(err)
			}
		}
		location := url.URL{Path: "/lk"}
		c.Redirect(http.StatusFound, location.RequestURI())
	}
}

// c.HTML(http.StatusOK, "index.html", gin.H{"id": c.Param("id")})

// var (
// 	name = c.PostForm("name")
// 	typs = c.PostForm("type")
// 	err  error
// )
// _, err = logic.ProdCreate(typs, name)
// c.HTML(http.StatusOK, "index.html", gin.H{"err": err})

// q := url.Values{}
// q.Set("id", c.Param("id"))
// location := url.URL{Path: "/dish/" + c.Param("id"), RawQuery: q.Encode()}
// c.Redirect(http.StatusFound, location.RequestURI())

// q := url.Values{}
// location := url.URL{Path: "/drink/" + c.Param("id"), RawQuery: q.Encode()}
// c.Redirect(http.StatusFound, location.RequestURI())
