// Створити вебсервер для перегляду інформації щодо класу школи.
// Користувач повинен мати можливість отримувати загальну інформацію про клас (список учнів, назва класу).
// Додаткові вимоги:
// інформація про учнів має зберігатися в оперативній пам'яті та бути доступною під час кожного запиту;
// отримання інформації про учня (наприклад, середній бал по предметах) має здійснюватись методом GET на адресі "/student/{id}",
// де {id} — унікальний ідентифікатор учня;
// дані можна отримати, лише якщо користувач є вчителем у цьому класі

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Student struct {
	ID     int               `json:"id"`
	Name   string            `json:"name"`
	Grades map[string]float64 `json:"grades"`
}

type Class struct {
	Name     string    `json:"name"`
	Teacher  string    `json:"teacher"`
	Students []Student `json:"students"`
}

var class = Class{
	Name:    "Математика",
	Teacher: "Бойко Ольга",
	Students: []Student{
		{ID: 1, Name: "Настя", Grades: map[string]float64{"математика": 85, "труди": 90}},
		{ID: 2, Name: "Катя", Grades: map[string]float64{"математика": 75, "біологія": 80}},
		{ID: 3, Name: "Петя", Grades: map[string]float64{"математика": 95, "географія": 85}},
	},
}

func main() {
	mux := http.NewServeMux()

	// Обробник для отримання інформації про клас
	mux.HandleFunc("/class", getClassInfo)
	mux.HandleFunc("/student/", getStudentInfo)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error happened: ", err.Error())
		return
	}
}

func getClassInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(class)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error encoding class info:", err)
		return
	}
}

// Обробник для отримання інформації про учня за його ID
func getStudentInfo(w http.ResponseWriter, r *http.Request) {
	userRole := r.Header.Get("Role") 	// Перевірка ролі користувача
	if userRole != "teacher" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Wrong student ID", http.StatusBadRequest)
		return
	}

	// Пошук учня за ID та повернення його інформації
	for _, student := range class.Students {
		if student.ID == id {
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(student)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				fmt.Println("Error encoding student info:", err)
				return
			}
			return
		}
	}
	http.Error(w, "Student not found", http.StatusNotFound)

}
