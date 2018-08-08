package api

import "net/http"

func handleQuestions(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		handleQuestionCreate(w, r)
	case "GET":
		params := pathParams(r, "/api/question/:id")
		questionId, ok := params[":id"]
		if ok { // GET /api/questions/ID
			handleQuestionGet(w, r, questionId)
			return
		}
		handleTopQuestions(w, r) // GET /api/questions/
	default:
		http.NotFound(w, r)
	}
}
