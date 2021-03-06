package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/poommin2543/app/ent"
	"github.com/poommin2543/app/ent/medicaltype"
)

// MedicaltypeController defines the struct for the medicaltype controller
type MedicaltypeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateMedicaltype handles POST requests for adding medicaltype entities
// @Summary Create medicaltype
// @Description Create medicaltype
// @ID create-medicaltype
// @Accept   json
// @Produce  json
// @Param medicaltype body ent.MedicalType true "Medicaltype entity"
// @Success 200 {object} ent.MedicalType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicaltypes [post]
func (ctl *MedicaltypeController) CreateMedicaltype(c *gin.Context) {
	obj := ent.MedicalType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "medicaltype binding failed",
		})
		return
	}

	u, err := ctl.client.MedicalType.
		Create().
		SetName(obj.Name).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetMedicaltype handles GET requests to retrieve a medicaltype entity
// @Summary Get a medicaltype entity by ID
// @Description get medicaltype by ID
// @ID get-medicaltype
// @Produce  json
// @Param id path int true "MedicalType ID"
// @Success 200 {object} ent.MedicalType
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicaltypes/{id} [get]
func (ctl *MedicaltypeController) GetMedicaltype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.MedicalType.
		Query().
		Where(medicaltype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListMedicaltype handles request to get a list of medicaltype entities
// @Summary List medicaltype entities
// @Description list medicaltype entities
// @ID list-medicaltype
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.MedicalType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicaltypes [get]
func (ctl *MedicaltypeController) ListMedicaltype(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	medicaltypes, err := ctl.client.MedicalType.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, medicaltypes)
}

// DeleteMedicaltype handles DELETE requests to delete a medicaltype entity
// @Summary Delete a medicaltype entity by ID
// @Description get medicaltype by ID
// @ID delete-medicaltype
// @Produce  json
// @Param id path int true "Medicaltype ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicaltypes/{id} [delete]
func (ctl *MedicaltypeController) DeleteMedicaltype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.MedicalType.
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

// UpdateMedicaltype handles PUT requests to update a medicaltype entity
// @Summary Update a medicaltype entity by ID
// @Description update medicaltype by ID
// @ID update-medicaltype
// @Accept   json
// @Produce  json
// @Param id path int true "Medicaltype ID"
// @Param medicaltype body ent.MedicalType true "Medicaltype entity"
// @Success 200 {object} ent.MedicalType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicaltypes/{id} [put]
func (ctl *MedicaltypeController) UpdateMedicaltype(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.MedicalType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "medicaltype binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.MedicalType.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewMedicaltypeController creates and registers handles for the medicaltype controller
func NewMedicaltypeController(router gin.IRouter, client *ent.Client) *MedicaltypeController {
	uc := &MedicaltypeController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitMedicaltypeController registers routes to the main engine
func (ctl *MedicaltypeController) register() {
	medicaltypes := ctl.router.Group("/medicaltypes")

	medicaltypes.GET("", ctl.ListMedicaltype)

	// CRUD
	medicaltypes.POST("", ctl.CreateMedicaltype)
	medicaltypes.GET(":id", ctl.GetMedicaltype)
	medicaltypes.PUT(":id", ctl.UpdateMedicaltype)
	medicaltypes.DELETE(":id", ctl.DeleteMedicaltype)
}
