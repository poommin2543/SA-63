package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/poommin2543/app/ent"
	"github.com/poommin2543/app/ent/systemequipment"
)

// SystemequipmentController defines the struct for the systemequipment controller
type SystemequipmentController struct {
	client *ent.Client
	router gin.IRouter
}
type Medicalequipments struct {
	Medicalequipment []Medicalequipment
}

type Medicalequipment struct {
	Name  string
	Stock int
}

// CreateSystemequipment handles POST requests for adding systemequipment entities
// @Summary Create systemequipment
// @Description Create systemequipment
// @ID create-systemequipment
// @Accept   json
// @Produce  json
// @Param systemequipment body ent.Systemequipment true "Systemequipment entity"
// @Success 200 {object} ent.Systemequipment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /systemequipments [post]
func (ctl *SystemequipmentController) CreateSystemequipment(c *gin.Context) {
	obj := ent.Systemequipment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "systemequipment binding failed",
		})
		return
	}

	u, err := ctl.client.Systemequipment.
		Create().
		SetAddedtime(obj.Addedtime).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetSystemequipment handles GET requests to retrieve a systemequipment entity
// @Summary Get a systemequipment entity by ID
// @Description get systemequipment by ID
// @ID get-systemequipment
// @Produce  json
// @Param id path int true "Systemequipment ID"
// @Success 200 {object} ent.Systemequipment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /systemequipments/{id} [get]
func (ctl *SystemequipmentController) GetSystemequipment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.Systemequipment.
		Query().
		Where(systemequipment.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListSystemequipment handles request to get a list of systemequipment entities
// @Summary List systemequipment entities
// @Description list systemequipment entities
// @ID list-systemequipment
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Systemequipment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /systemequipments [get]
func (ctl *SystemequipmentController) ListSystemequipment(c *gin.Context) {
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

	systemequipments, err := ctl.client.Systemequipment.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, systemequipments)
}

// DeleteSystemequipment handles DELETE requests to delete a systemequipment entity
// @Summary Delete a systemequipment entity by ID
// @Description get systemequipment by ID
// @ID delete-systemequipment
// @Produce  json
// @Param id path int true "Systemequipment ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /systemequipments/{id} [delete]
func (ctl *SystemequipmentController) DeleteSystemequipment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Systemequipment.
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

// UpdateSystemequipment handles PUT requests to update a systemequipment entity
// @Summary Update a systemequipment entity by ID
// @Description update systemequipment by ID
// @ID update-systemequipment
// @Accept   json
// @Produce  json
// @Param id path int true "Systemequipment ID"
// @Param systemequipment body ent.Systemequipment true "Systemequipment entity"
// @Success 200 {object} ent.Systemequipment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /systemequipments/{id} [put]
func (ctl *SystemequipmentController) UpdateSystemequipment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Systemequipment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "systemequipment binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Systemequipment.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewSystemequipmentController creates and registers handles for the systemequipment controller
func NewSystemequipmentController(router gin.IRouter, client *ent.Client) *SystemequipmentController {
	uc := &SystemequipmentController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitSystemequipmentController registers routes to the main engine
func (ctl *SystemequipmentController) register() {
	systemequipments := ctl.router.Group("/systemequipments")

	systemequipments.GET("", ctl.ListSystemequipment)

	// CRUD
	systemequipments.POST("", ctl.CreateSystemequipment)
	systemequipments.GET(":id", ctl.GetSystemequipment)
	systemequipments.PUT(":id", ctl.UpdateSystemequipment)
	systemequipments.DELETE(":id", ctl.DeleteSystemequipment)
}
