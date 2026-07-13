package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"rest-api/internal/models"
	"rest-api/pkg/utils"
)

var (
	teachers = make(map[int]models.Teacher)
	mutex    = &sync.Mutex{}
	nextID   = 1
)

// Initialize some dummy data
func init() {
	teachers[nextID] = models.Teacher{
		Id:        nextID,
		FirstName: "Oghenemaro",
		LastName:  "Ebedi",
		Class:     "1A",
		Subject:   "Applied Quantum Physics",
	}
	nextID++
	teachers[nextID] = models.Teacher{
		Id:        nextID,
		FirstName: "Kosiso",
		LastName:  "Achalugo",
		Class:     "3A",
		Subject:   "Project Management",
	}
	nextID++
}

func getTeachersHandler(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")
	teacherList := make([]models.Teacher, 0, len(teachers))

	id, err := utils.ExtractId(r.URL.Path, "teachers")
	if err != nil {
		http.Error(w, "Failed to extract ID from URL", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content_Type", "application/json")

	if id != 0 {
		teacher, ok := teachers[id]

		if ok {
			response := struct {
				Status string         `json:"status"`
				Data   models.Teacher `json:"data"`
			}{
				Status: "success",
				Data:   teacher,
			}
			json.NewEncoder(w).Encode(response)
			return
		} else {
			http.Error(w, "Teacher not found.", http.StatusNotFound)
			return
		}
	}

	for _, teacher := range teachers {
		if firstName == "" || teacher.FirstName == firstName && lastName == "" || teacher.LastName == lastName {
			teacherList = append(teacherList, teacher)
		}
	}

	response := utils.BuildMultipleTeacherResponse("success", len(teachers), teacherList)

	json.NewEncoder(w).Encode(response)
}

func addTeacherHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var newTeachers []models.Teacher

	err := json.NewDecoder(r.Body).Decode(&newTeachers)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	addedTeachers := make([]models.Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		newTeacher.Id = nextID
		teachers[nextID] = newTeacher
		addedTeachers[i] = newTeacher
		nextID++
	}

	response := utils.BuildMultipleTeacherResponse("success", len(addedTeachers), addedTeachers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}

func TeachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTeachersHandler(w, r)
	case http.MethodPost:
		addTeacherHandler(w, r)
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on teachers route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on teachers route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on teachers route"))
	}
}
