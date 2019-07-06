package controllers

import (
    "class-review-backend/models"
    "class-review-backend/services"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
)

type ProfessorController struct {
    Services *services.Services
    Routes   *gin.RouterGroup
}

func DefaultProfessorController(eng *gin.Engine, services *services.Services) *ProfessorController {
    pc := &ProfessorController{
        Routes:   eng.Group("/professor"),
        Services: services,
    }

	pc.Routes.GET("", pc.getAllProfessors)
    pc.Routes.GET("/:uniqname", pc.getProfessorByUniqname)
    pc.Routes.POST("", pc.postProfessor)
    pc.Routes.PUT("/:uniqname", pc.updateProfessor)
    pc.Routes.DELETE("/:uniqname", pc.deleteProfessor)
    // pc.Routes.GET("stats", pc.getProfessorStats)
    // pc.Routes.GET("stats/:uniqname", pc.getProfessorStatsByUniqname)
	return pc
}

func (pc *ProfessorController) getAllProfessors(c *gin.Context) {
    professors, err := pc.Services.ProfessorService.GetAllProfessors()
    if err != nil {
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
        c.JSON(http.StatusNotFound, "No professor with provided uniqname")
        return
    }
    c.JSON(http.StatusOK, professor)
    return
}

func (pc *ProfessorController) postProfessor(c *gin.Context) {
    var professorInput models.Professor
    err := c.BindJSON(&professorInput)
    if err != nil {
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
    if err != nil {
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
        c.JSON(http.StatusNotFound, "No professor with provided uniqname")
        return
    }
    c.JSON(http.StatusOK, "OK")
    return
}

func (pc *ProfessorController) getProfessorStats(c *gin.Context) {
    professorStats, err := pc.Services.ProfessorService.GetProfessorStats()
    if err != nil {
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
        c.JSON(http.StatusNotFound, "No professor stats with provided uniqname")
        return
    }
    c.JSON(http.StatusOK, professorStats)
    return
}