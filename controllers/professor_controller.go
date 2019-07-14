package controllers

import (
	"class-review-backend/models"
	"class-review-backend/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfessorController struct {
	Services *services.Services
	Routes   *Routes
}

func DefaultProfessorController(routes *Routes, services *services.Services) *ProfessorController {
	pc := &ProfessorController{
		Routes:   routes,
		Services: services,
	}

	pc.Routes.Public.GET("/professor/names", pc.getAllProfessors)
	pc.Routes.Public.GET("/professor/names/:uniqname", pc.getProfessorByUniqname)
	pc.Routes.Private.POST("/professor/names", pc.postProfessor)
	pc.Routes.Private.PUT("/professor/names/:uniqname", pc.updateProfessor)
	pc.Routes.Private.DELETE("/professor/names/:uniqname", pc.deleteProfessor)
	pc.Routes.Private.GET("/professor/stats", pc.getProfessorStats)
	pc.Routes.Private.GET("/professor/stats/:uniqname", pc.getProfessorStatsByUniqname)
	return pc
}

func (pc *ProfessorController) getAllProfessors(c *gin.Context) {
	professors, err := pc.Services.ProfessorService.GetAllProfessors()
	if err != nil {
		log.Println("Professors not found")
		c.JSON(http.StatusNotFound, "Professors not found")
		return
	}
	c.JSON(http.StatusOK, professors)
	return
}

func (pc *ProfessorController) getProfessorByUniqname(c *gin.Context) {
	uniqname := c.Param("uniqname")
	if uniqname == "" {
		log.Println("No uniqname in url")
		c.JSON(http.StatusBadRequest, "No uniqname in url")
		return
	}
	professor, err := pc.Services.ProfessorService.GetProfessorByUniqname(uniqname)
	if err != nil {
		log.Println("No professor with provided uniqname")
		c.JSON(http.StatusNotFound, "No professor with provided uniqname")
		return
	}
	c.JSON(http.StatusOK, professor)
	return
}

func (pc *ProfessorController) postProfessor(c *gin.Context) {
	var professorInput models.Professor
	err := c.BindJSON(&professorInput)
	if err != nil || professorInput.Name == "" || professorInput.Uniqname == "" {
		log.Println("Invalid request body")
		c.JSON(http.StatusBadRequest, "Invalid request body")
		return
	}
	err = pc.Services.ProfessorService.PostProfessor(&professorInput)
	if err != nil {
		log.Println("POST request failed")
		c.JSON(http.StatusNotFound, "POST request failed")
		return
	}
	c.JSON(http.StatusOK, "OK")
	return
}

func (pc *ProfessorController) updateProfessor(c *gin.Context) {
	uniqname := c.Param("uniqname")
	if uniqname == "" {
		log.Println("No uniqname in url")
		c.JSON(http.StatusBadRequest, "No uniqname in url")
		return
	}
	var professorInput models.Professor
	err := c.BindJSON(&professorInput)
	if err != nil || professorInput.Name == "" {
		log.Println("Invalid request body")
		c.JSON(http.StatusBadRequest, "Invalid request body")
		return
	}
	professor, err := pc.Services.ProfessorService.UpdateProfessor(&professorInput, uniqname)
	if err != nil {
		log.Println("PUT request failed")
		c.JSON(http.StatusNotFound, "PUT request failed")
		return
	}
	c.JSON(http.StatusOK, professor)
	return
}

func (pc *ProfessorController) deleteProfessor(c *gin.Context) {
	uniqname := c.Param("uniqname")
	if uniqname == "" {
		log.Println("No uniqname in url")
		c.JSON(http.StatusBadRequest, "No uniqname in url")
		return
	}
	err := pc.Services.ProfessorService.DeleteProfessor(uniqname)
	if err != nil {
		log.Println("DELETE request failed")
		c.JSON(http.StatusNotFound, "DELETE request failed")
		return
	}
	c.JSON(http.StatusOK, "OK")
	return
}

func (pc *ProfessorController) getProfessorStats(c *gin.Context) {
	professorStats, err := pc.Services.ProfessorService.GetProfessorStats()
	if err != nil {
		log.Println("Professor stats not found")
		c.JSON(http.StatusNotFound, "Professor stats not found")
		return
	}
	c.JSON(http.StatusOK, professorStats)
	return
}

func (pc *ProfessorController) getProfessorStatsByUniqname(c *gin.Context) {
	uniqname := c.Param("uniqname")
	if uniqname == "" {
		log.Println("No uniqname in url")
		c.JSON(http.StatusBadRequest, "No uniqname in url")
		return
	}
	professorStats, err := pc.Services.ProfessorService.GetProfessorStatsByUniqname(uniqname)
	if err != nil {
		log.Println("No professor stats with provided uniqname")
		c.JSON(http.StatusNotFound, "No professor stats with provided uniqname")
		return
	}
	c.JSON(http.StatusOK, professorStats)
	return
}
