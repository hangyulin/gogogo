package router

import (
	"github.com/gorilla/mux"
	"go_server/middleware"
)

func Router() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/qna", middleware.GetQuestionAndAnswerResult).Methods("POST")

	return router
}