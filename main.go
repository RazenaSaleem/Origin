package main

import (
	"fmt"
	api "openapiEMP/api"

	"github.com/labstack/echo/v4"

	// api "openapiEMP/api/server"
	"openapiEMP/pkg/server"

	"github.com/kube-tarian/kad/agent/pkg/logging"
)

func main() {
	log := logging.NewLogger()
	log.Infof("Starting server")

	serverInstance := server.NewServer(log)

	e := echo.New()
	wrapper := &api.ServerInterfaceWrapper{
		Handler: serverInstance,
	}
	// wrapper := &server.ServerInterfaceWrapper{
	//     Handler: serverInstance,
	// }

	e.GET("/employee", wrapper.GetEmployees)
	e.POST("/employee", wrapper.CreateEmployee)
	e.PUT("/employee/:EmployeeID", wrapper.UpdateEmployee)
	e.DELETE("/employee/:EmployeeID", wrapper.DeleteEmployee)

	if err := e.Start(fmt.Sprintf("%s:%d", "0.0.0.0", 52002)); err != nil {
		log.Errorf("server run returned, reason: %s", err.Error())
	}

	log.Infof("Stopped server")
}

/*
package main

import (
	"fmt"
	//"net/http"

	"github.com/labstack/echo/v4"
	//"github.com/gin-gonic/gin"
	"Newfolder/openapiEMP/api"
	"Newfolder/openapiEMP/pkg/server"
	"github.com/kube-tarian/kad/agent/pkg/logging"
)

func main() {
	log := logging.NewLogger()
	log.Infof("Starting server")

	serverInstance := server.NewServer(log)

	r := gin.Default()

	//api.RegisterHandlers(r, serverInstance)

	wrapper := &server.ServerInterfaceWrapper{
		Handler: serverInstance,
	}


	r.POST("/employee", wrapper.CreateEmployee)
	r.GET("/employee", wrapper.GetEmployees)
	r.PUT("/employee/:EmployeeID", wrapper.UpdateEmployee)
	r.DELETE("/employee/:EmployeeID", wrapper.DeleteEmployee)

	if err := r.Run(fmt.Sprintf("%s:%d", "0.0.0.0", 52002)); err != nil {
		log.Errorf("server run returned, reason: %s", err.Error())
	}

	log.Infof("Stopped server")
}*/
