package main

import (
	CustomerController "backend/api/controller/customer"
	EmployeeController "backend/api/controller/employee"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	//Employee API Method
	router.GET("/employee", EmployeeController.GET)
	router.POST("/employee", EmployeeController.POST)
	router.GET("/employee/:id", EmployeeController.GetEmployeeByID)  //GET By ID
	router.PUT("/employee", EmployeeController.PUT)
	router.DELETE("/employee", EmployeeController.DELETE)

	//Customer API Method
	router.GET("/customer", CustomerController.GET)
	router.POST("/customer", CustomerController.POST)
	router.PUT("/customer", CustomerController.PUT)
	router.DELETE("/customer", CustomerController.DELETE)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
