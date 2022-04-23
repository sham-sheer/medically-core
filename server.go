package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

// Server is an http server that handles REST requests.
type Server struct {
	db *gorm.DB
}

// NewServer creates a new instance of a Server.
func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}

// RegisterRouter registers a router onto the Server.
func (s *Server) RegisterRouter(router *gin.Engine) {
	router.GET("/ping", s.ping)

	router.GET("/user", s.getUsers)
	router.POST("/user", s.createUser)
	router.GET("/user/:userID", s.getUser)
	router.PUT("/user/:userID", s.updateUser)
	router.DELETE("/user/:userID", s.deleteUser)

	router.GET("/med", s.getMeds)
	router.POST("/med", s.createMed)
	router.GET("/med/:medID", s.getMed)
	router.PUT("/med/:medID", s.updateMed)
	router.DELETE("/med/:medID", s.deleteMed)

	router.GET("/disease", s.getDiseases)
	router.POST("/disease", s.createDisease)
	router.GET("/disease/:diseaseID", s.getDisease)
	router.PUT("/disease/:diseaseID", s.updateDisease)
	router.DELETE("/disease/:diseasesID", s.deleteDisease)

	router.GET("/clinic", s.getClinics)
	router.POST("/clinic", s.createClinic)
	router.GET("/clinic/:clinicID", s.getClinic)
	router.PUT("/clinic/:clinicID", s.updateClinic)
	router.DELETE("/clinic/:clinicID", s.deleteClinic)
}

// ------------------------------- User Server Methods ------------------------------------//
func (s *Server) getUsers(c *gin.Context) {
	var users []User
	if err := s.db.Find(&users).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	}
	c.JSON(http.StatusOK, users)
}

func (s *Server) createUser(c *gin.Context) {
	var user User
	if err := BindJSON(c, &user); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		return
	}

	if err := s.db.Create(&user).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	}
	c.JSON(http.StatusOK, user)

}

func (s *Server) getUser(c *gin.Context) {
	id := c.Param("userID")
	var user User
	if err := s.db.Find(&user, id).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (s *Server) updateUser(c *gin.Context) {
	var user User
	if err := BindJSON(c, &user); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		return
	}

	if err := s.db.Save(user).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (s *Server) deleteUser(c *gin.Context) {
	userID := c.Param("userID")
	req := s.db.Delete(User{}, "ID = ?", userID)
	if err := req.Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else if req.RowsAffected == 0 {
		c.String(http.StatusNotFound, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, gin.H{"userId": userID})
	}
}

// ------------------------------- ------------------- ------------------------------------//

// ----------------------------  Medication Server Methods ---------------------------------//
func (s *Server) getMeds(c *gin.Context) {
	var meds []Med
	if err := s.db.Find(&meds).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	}
	c.JSON(http.StatusOK, meds)
}

func (s *Server) createMed(c *gin.Context) {
	var med Med
	if err := BindJSON(c, &med); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		return
	}

	if err := s.db.Create(&med).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	}
	c.JSON(http.StatusOK, med)
}

func (s *Server) getMed(c *gin.Context) {
	var med Med
	if err := s.db.Find(&med, c.Param("medID")).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, med)
	}
}

func (s *Server) updateMed(c *gin.Context) {
	var med Med
	if err := BindJSON(c, &med); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		return
	}

	if err := s.db.Save(med).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, med)
	}
}

func (s *Server) deleteMed(c *gin.Context) {
	medID := c.Param("medID")
	req := s.db.Delete(Med{}, "ID = ?", medID)
	if err := req.Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else if req.RowsAffected == 0 {
		c.String(http.StatusNotFound, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, medID)
	}
}

// ------------------------------- ------------------- ------------------------------------//

// ----------------------------  Disease Server Methods ---------------------------------//
func (s *Server) getDiseases(c *gin.Context) {
	var diseases []Disease
	if err := s.db.Find(&diseases).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, diseases)
	}
}

func (s *Server) createDisease(c *gin.Context) {
	var disease Disease
	if err := BindJSON(c, &disease); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		return
	}

	if err := s.db.Create(&disease).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	}
	c.JSON(http.StatusOK, disease)
}

func (s *Server) getDisease(c *gin.Context) {
	var disease Disease
	if err := s.db.Find(&disease, c.Param("diseaseID")).Error; err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, disease)
	}
}

func (s *Server) updateDisease(c *gin.Context) {
	var disease Disease
	if err := BindJSON(c, &disease); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		return
	}

	if err := s.db.Save(disease).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, disease)
	}
}

func (s *Server) deleteDisease(c *gin.Context) {
	diseaseID := c.Param("diseaseID")
	req := s.db.Delete(Disease{}, "ID = ?", diseaseID)
	if err := req.Error; err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
	} else if req.RowsAffected == 0 {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, diseaseID)
	}
}
// ------------------------------- ------------------- ------------------------------------//

// ---------------------------- Clinic Server Methods ---------------------------------//
func (s *Server) getClinics(c *gin.Context) {
	var clinics []Clinic
	if err := s.db.Find(&clinics).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, clinics)
	}
}

func (s *Server) createClinic(c *gin.Context) {
	var clinic Clinic
	if err := BindJSON(c, &clinic); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		return
	}

	if err := s.db.Create(&clinic).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, clinic)
	}
}

func (s *Server) getClinic(c *gin.Context) {
	var clinic Clinic
	if err := s.db.Find(&clinic, c.Param("clinicID")).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, clinic)
	}
}

func (s *Server) updateClinic(c *gin.Context) {
	var clinic Clinic
	if err := BindJSON(c, &clinic); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
		return
	}

	if err := s.db.Save(clinic).Error; err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, clinic)
	}
}

func (s *Server) deleteClinic(c *gin.Context) {
	clinicId := c.Param("clinicID")
	req := s.db.Delete(Clinic{}, "ID = ?", clinicId)
	if err := req.Error; err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err))
	} else if req.RowsAffected == 0 {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
	} else {
		c.JSON(http.StatusOK, clinicId)
	}
}
// ------------------------------- ------------------- ------------------------------------//


func (s *Server) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "service ready to go!"})
}

func BindJSON(c *gin.Context, obj interface{}) (err error) {
	if err = c.ShouldBindWith(obj, binding.JSON); err != nil {
		return err
	}
	return
}