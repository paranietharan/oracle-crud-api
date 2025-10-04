package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oracle-crud-api/internal/model"
	"oracle-crud-api/internal/service"
)

type PersonHandler struct {
	service *service.PersonService
}

func NewPersonHandler(service *service.PersonService) *PersonHandler {
	return &PersonHandler{service: service}
}

func (h *PersonHandler) Create(c *gin.Context) {
	var p model.Person
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create person"})
		return
	}
	c.JSON(http.StatusCreated, p)
}

func (h *PersonHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	p, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "person not found"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func (h *PersonHandler) GetAll(c *gin.Context) {
	people, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch persons"})
		return
	}
	c.JSON(http.StatusOK, people)
}

func (h *PersonHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var p model.Person
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	p.Id = id

	if err := h.service.Update(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update person"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func (h *PersonHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete person"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "person deleted"})
}
