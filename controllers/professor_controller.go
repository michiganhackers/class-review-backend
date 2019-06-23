package controllers

import (
    "class-review-backend/services"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "strconv"
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

	pc.Routes.GET("", pc.getProfessors)
    pc.Routes.GET("review/:id", pc.getProfessorByReviewID)
    pc.Routes.GET("course/:id", pc.getProfessorsByCourseID)
    pc.Routes.GET("stats", pc.getProfessorStats)
    pc.Routes.GET("stats/:name", pc.getProfessorStatsByName)
	return pc
}

func (pc *ProfessorController) getProfessors(c *gin.Context) {
    professors, err := pc.Services.ProfessorService.GetProfessors()
    if err != nil {
        c.JSON(http.StatusNotFound, "Professors not found")
        return
    }
    c.JSON(http.StatusOK, professors)
    return
}

func (pc *ProfessorController) getProfessorByReviewID(c *gin.Context) {
    idStr := c.Param("id")
    if idStr == "" {
        log.Println("No id in url")
        c.JSON(http.StatusBadRequest, "No id in url")
        return
    }
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        log.Println("Bad id param")
        c.JSON(http.StatusBadRequest, "Bad id param")
        return
    }
    professor, err := pc.Services.ProfessorService.GetProfessorByReviewID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, "No professor with provided review id")
        return
    }
    c.JSON(http.StatusOK, professor)
    return
}

func (pc *ProfessorController) getProfessorsByCourseID(c *gin.Context) {
    idStr := c.Param("id")
    if idStr == "" {
        log.Println("No id in url")
        c.JSON(http.StatusBadRequest, "No id in url")
        return
    }
    id, err := strconv.ParseUint(idStr, 10, 64)
    if err != nil {
        log.Println("Bad id param")
        c.JSON(http.StatusBadRequest, "Bad id param")
        return
    }
    professor, err := pc.Services.ProfessorService.GetProfessorsByCourseID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, "No professor with provided course id")
        return
    }
    c.JSON(http.StatusOK, professor)
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

func (pc *ProfessorController) getProfessorStatsByName(c *gin.Context) {
    name := c.Param("name")
    if name == "" {
        log.Println("No name in url")
        c.JSON(http.StatusBadRequest, "No name in url")
        return
    }
    professorStats, err := pc.Services.ProfessorService.GetProfessorStatsByName(name)
    if err != nil {
        c.JSON(http.StatusNotFound, "No professor stats with provided name")
        return
    }
    c.JSON(http.StatusOK, professorStats)
    return
}