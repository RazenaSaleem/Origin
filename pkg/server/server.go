package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/kube-tarian/kad/agent/pkg/logging"
)

type Employee struct {
	EmployeeID   string `json:"EmployeeID"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Department   string `json:"Department"`
	Position     string `json:"Position"`
}
type Server struct {
	log         logging.Logger
	EmpDetails  map[string]Employee
}

func NewServer(log logging.Logger) *Server {
	return &Server{
		log:        log,
		EmpDetails: make(map[string]Employee),
	}
}

func (s *Server) CreateEmployee(c echo.Context) error {
	var employee Employee
	if err := c.Bind(&employee); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	s.EmpDetails[employee.EmployeeID] = employee

	return c.JSON(http.StatusCreated, employee)
}

func (s *Server) GetEmployees(c echo.Context) error {
	employees := make([]Employee, 0, len(s.EmpDetails))

	for _, employee := range s.EmpDetails {
		employees = append(employees, employee)
	}

	return c.JSON(http.StatusOK, employees)
}
func (s *Server) UpdateEmployee(c echo.Context,EmployeeId string) error {
	employeeID := c.Param("EmployeeID")
	employee, exists := s.EmpDetails[employeeID]
	if !exists {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Employee not found"})
	}

	var updatedEmployee Employee
	if err := c.Bind(&updatedEmployee); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	employee.FirstName = updatedEmployee.FirstName
	employee.LastName = updatedEmployee.LastName
	employee.Department = updatedEmployee.Department
	employee.Position = updatedEmployee.Position

	s.EmpDetails[employeeID] = employee
	return c.JSON(http.StatusOK, employee)
}

func (s *Server) DeleteEmployee(c echo.Context,EmployeeId string) error {
	employeeID := c.Param("EmployeeID")
	_, exists := s.EmpDetails[employeeID]
	if !exists {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Employee not found"})
	}

	delete(s.EmpDetails, employeeID)
	return c.JSON(http.StatusOK, echo.Map{"message": "Employee deleted successfully"})
}



/*
package server

import (
	//"encoding/json"
	"net/http"
	"github.com/labstack/echo/v4"
	//"github.com/gin-gonic/gin"
	//"openapiEMP/pkg/model"
	//"openapiEMP/api"
	"github.com/kube-tarian/kad/agent/pkg/logging"
)

type Employee struct {
	EmployeeID   string `json:"EmployeeID"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Department   string `json:"Department"`
	Position     string `json:"Position"`
}
type Server struct {
	log         logging.Logger
	EmpDetails  map[string]Employee
}

func NewServer(log logging.Logger) *Server {
	return &Server{
		log:        log,
		EmpDetails: make(map[string]Employee),
	}
}

func (s *Server) CreateEmployee(c echo.Context) {
	var employee Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s.EmpDetails[employee.EmployeeID] = employee

	c.JSON(http.StatusCreated, employee)
}

func (s *Server) GetEmployees(c echo.Context) {
	employees := make([]Employee, 0, len(s.EmpDetails))

	for _, employee := range s.EmpDetails {
		employees = append(employees, employee)
	}

	c.JSON(http.StatusOK, employees)
}

func (s *Server) UpdateEmployee(c *gin.Context) {
	employeeID := c.Param("EmployeeID")
	employee, exists := s.EmpDetails[employeeID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	var updatedEmployee Employee
	if err := c.ShouldBindJSON(&updatedEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	employee.FirstName = updatedEmployee.FirstName
	employee.LastName = updatedEmployee.LastName
	employee.Department = updatedEmployee.Department
	employee.Position = updatedEmployee.Position

	s.EmpDetails[employeeID] = employee
	c.JSON(http.StatusOK, employee)
}

func (s *Server) DeleteEmployee(c *gin.Context) {
	employeeID := c.Param("EmployeeID")
	_, exists := s.EmpDetails[employeeID]
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	delete(s.EmpDetails, employeeID)
	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})

}*/