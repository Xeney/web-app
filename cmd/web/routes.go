package main

func initializeRoutes() {
	router.GET("/", SessionCheck, GetHome)
	router.GET("/courses", SessionCheck, GetCourses)
	router.POST("/courses/:id", SessionCheck, PostCoursesTarget)
	router.GET("/courses/:id", SessionCheck, GetCoursesTarget)
	router.GET("/login", SessionCheck, GetLogin)
	router.POST("/login", SessionCheck, PostLogin)
	router.GET("/lk", SessionCheck, GetLk)
	router.POST("/lk", SessionCheck, PostLk)
	router.GET("/exit", SessionCheck, GetExit)
	router.GET("/test", SessionCheck, GetTest)
	router.POST("/test", SessionCheck, PostTest)
}
