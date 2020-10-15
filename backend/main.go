package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/poommin2543/app/controllers"
	_ "github.com/poommin2543/app/docs"
	"github.com/poommin2543/app/ent"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Medicalequipments struct {
	Medicalequipment []Medicalequipment
}

type Medicalequipment struct {
	Name  string
	Stock int
}

type Medicaltypes struct {
	Medicaltype []Medicaltype
}

type Medicaltype struct {
	Name string
}

type Physicians struct {
	Physician []Physician
}

type Physician struct {
	Name  string
	Email string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewMedicalequipmentController(v1, client)
	controllers.NewMedicaltypeController(v1, client)
	controllers.NewPhysicianController(v1, client)
	controllers.NewSystemequipmentController(v1, client)

	// Set Physicians Data
	physicians := Physicians{
		Physician: []Physician{
			Physician{"Chanwit Kaewkasi", "chanwit@gmail.com"},
			Physician{"Name Surname", "me@example.com"},
		},
	}

	for _, ph := range physicians.Physician {
		client.Physician.
			Create().
			SetEmail(ph.Email).
			SetName(ph.Name).
			Save(context.Background())
	}

	medicaltypes := Medicaltypes{
		Medicaltype: []Medicaltype{
			Medicaltype{"ใช้แล้วทิ้ง"},
			Medicaltype{"นำกลับมาใช้ใหม่"},
		},
	}

	for _, mt := range medicaltypes.Medicaltype {
		client.MedicalType.
			Create().
			SetName(mt.Name).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
