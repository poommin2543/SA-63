package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/poommin2543/app/ent"
   "github.com/poommin2543/app/ent/medicalequipment"
   "github.com/gin-gonic/gin"
)
 
// MedicalequipmentController defines the struct for the medicalequipment controller
type MedicalequipmentController struct {
   client *ent.Client
   router gin.IRouter
}
// CreateMedicalequipment handles POST requests for adding medicalequipment entities
// @Summary Create medicalequipment
// @Description Create medicalequipment
// @ID create-medicalequipment
// @Accept   json
// @Produce  json
// @Param medicalequipment body ent.Medicalequipment true "Medicalequipment entity"
// @Success 200 {object} ent.Medicalequipment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalequipments [post]
func (ctl *MedicalequipmentController) CreateMedicalequipment(c *gin.Context) {
	obj := ent.Medicalequipment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "medicalequipment binding failed",
		})
		return
	}
  
	u, err := ctl.client.Medicalequipment.
		Create().
		SetMedicalID(obj.MedicalID).
		SetMedicalNAME(obj.MedicalNAME).
		SetMedicalStock(obj.MedicalStock).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, u)
 }
 // GetMedicalequipment handles GET requests to retrieve a medicalequipment entity
// @Summary Get a medicalequipment entity by ID
// @Description get medicalequipment by ID
// @ID get-medicalequipment
// @Produce  json
// @Param id path int true "Medicalequipment ID"
// @Success 200 {object} ent.Medicalequipment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalequipments/{id} [get]
func (ctl *MedicalequipmentController) GetMedicalequipment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	u, err := ctl.client.Medicalequipment.
		Query().
		Where(medicalequipment.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, u)
 }
 // ListMedicalequipment handles request to get a list of medicalequipment entities
// @Summary List medicalequipment entities
// @Description list medicalequipment entities
// @ID list-medicalequipment
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Medicalequipment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalequipments [get]
func (ctl *MedicalequipmentController) ListMedicalequipment(c *gin.Context) {
   limitQuery := c.Query("limit")
   limit := 10
   if limitQuery != "" {
       limit64, err := strconv.ParseInt(limitQuery, 10, 64)
       if err == nil {limit = int(limit64)}
   }
 
   offsetQuery := c.Query("offset")
   offset := 0
   if offsetQuery != "" {
       offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
       if err == nil {offset = int(offset64)}
   }
 
   medicalequipments, err := ctl.client.Medicalequipment.
       Query().
       Limit(limit).
       Offset(offset).
       All(context.Background())
   	if err != nil {
       c.JSON(400, gin.H{"error": err.Error(),})
       return
   }
 
   c.JSON(200, medicalequipments)
}
// DeleteMedicalequipment handles DELETE requests to delete a medicalequipment entity
// @Summary Delete a medicalequipment entity by ID
// @Description get medicalequipment by ID
// @ID delete-medicalequipment
// @Produce  json
// @Param id path int true "Medicalequipment ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalequipments/{id} [delete]
func (ctl *MedicalequipmentController) DeleteMedicalequipment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Medicalequipment.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
 }
 // UpdateMedicalequipment handles PUT requests to update a medicalequipment entity
// @Summary Update a medicalequipment entity by ID
// @Description update medicalequipment by ID
// @ID update-medicalequipment
// @Accept   json
// @Produce  json
// @Param id path int true "Medicalequipment ID"
// @Param medicalequipment body ent.Medicalequipment true "Medicalequipment entity"
// @Success 200 {object} ent.Medicalequipment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicalequipments/{id} [put]
func (ctl *MedicalequipmentController) UpdateMedicalequipment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Medicalequipment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "medicalequipment binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Medicalequipment.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, u)
 }
 // NewMedicalequipmentController creates and registers handles for the medicalequipment controller
func NewMedicalequipmentController(router gin.IRouter, client *ent.Client) *MedicalequipmentController {
	uc := &MedicalequipmentController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitMedicalequipmentController registers routes to the main engine
 func (ctl *MedicalequipmentController) register() {
	medicalequipments := ctl.router.Group("/medicalequipments")
  
	medicalequipments.GET("", ctl.ListMedicalequipment)
  
	// CRUD
	medicalequipments.POST("", ctl.CreateMedicalequipment)
	medicalequipments.GET(":id", ctl.GetMedicalequipment)
	medicalequipments.PUT(":id", ctl.UpdateMedicalequipment)
	medicalequipments.DELETE(":id", ctl.DeleteMedicalequipment)
 }
 
 