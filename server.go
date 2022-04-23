package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/medically-core/model"
	"github.com/julienschmidt/httprouter"
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
func (s *Server) RegisterRouter(router *httprouter.Router) {
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
func (s *Server) getUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var users []model.User
	if err := s.db.Find(&users).Error; err != nil {
		http.Error(w, err.Error(), errToStatusCode(err))
	} else {
		writeJSONResult(w, users)
	}
}

func (s *Server) createUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), errToStatusCode(err))
		return
	}

	if err := s.db.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), errToStatusCode(err))
	} else {
		writeJSONResult(w, user)
	}
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user model.User
	if err := s.db.Find(&user, ps.ByName("userID")).Error; err != nil {
		http.Error(w, err.Error(), errToStatusCode(err))
	} else {
		writeJSONResult(w, user)
	}
}

func (s *Server) updateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), errToStatusCode(err))
		return
	}

	if err := s.db.Save(user).Error; err != nil {
		http.Error(w, err.Error(), errToStatusCode(err))
	} else {
		writeJSONResult(w, user)
	}
}

func (s *Server) deleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := ps.ByName("userID")
	req := s.db.Delete(model.User{}, "ID = ?", userID)
	if err := req.Error; err != nil {
		http.Error(w, err.Error(), errToStatusCode(err))
	} else if req.RowsAffected == 0 {
		http.Error(w, "", http.StatusNotFound)
	} else {
		writeTextResult(w, "ok")
	}
}

// ------------------------------- ------------------- ------------------------------------//

// ----------------------------  Medication Server Methods ---------------------------------//
// ------------------------------- ------------------- ------------------------------------//

// ----------------------------  Disease Server Methods ---------------------------------//
// ------------------------------- ------------------- ------------------------------------//

// ---------------------------- Clinic Server Methods ---------------------------------//
// ------------------------------- ------------------- ------------------------------------//


func (s *Server) ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	writeTextResult(w, "go/gorm")
}

func writeTextResult(w http.ResponseWriter, res string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, res)
}

func writeJSONResult(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

func writeMissingParamError(w http.ResponseWriter, paramName string) {
	http.Error(w, fmt.Sprintf("missing query param %q", paramName), http.StatusBadRequest)
}

func errToStatusCode(err error) int {
	switch err {
	case gorm.ErrRecordNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}