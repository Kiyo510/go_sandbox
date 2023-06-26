package http

import (
	"fmt"
	"github.com/Kiyo510/go_sandbox/pkg/domain/service"
	"github.com/Kiyo510/go_sandbox/pkg/infra"
	"github.com/Kiyo510/go_sandbox/pkg/infra/mysql"
	"github.com/Kiyo510/go_sandbox/pkg/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	apiVersion      = "/v1"
	healthCheckRoot = "/health_check"
	// student
	studentsAPIRoot = apiVersion + "/students"
	studentIDParam  = "student_id"
)

func InitRouter() *echo.Echo {
	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	// health check
	healthCheckGroup := e.Group(healthCheckRoot)
	{
		relativePath := ""
		healthCheckGroup.GET(relativePath, healthCheck)
	}

	// student
	mySQLConn := infra.NewMySQLConnector()
	studentRepository := mysql.NewStudentRepository(mySQLConn.Conn)
	studentService := service.NewStudentService(studentRepository)
	studentUsecase := usecase.NewUserUsecase(studentService)

	studentGroup := e.Group(studentsAPIRoot)
	{
		handler := NewStudentHandler(studentUsecase)
		// v1/students
		relativePath := ""
		studentGroup.GET(relativePath, handler.FindAllStudents())

		//v1/students/{student_id}
		relativePath = fmt.Sprintf("/:%s", studentIDParam)
		studentGroup.GET(relativePath, handler.FindStudentByID())
	}

	return e
}
