package handler

import (
	"assignment-3/entity"
	"assignment-3/service"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func StartApp() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	statusService := service.NewStatusService(r)

	go statusService.GenerateStatusData(1, 100)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, "Error occurred while trying to parse template", http.StatusInternalServerError)
			return
		}

		var data = entity.Data{Status: entity.Status{}}

		b, err := ioutil.ReadFile("data/status.json")
		if err != nil {
			http.Error(w, "Error occurred while trying to read file", http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(b, &data)
		if err != nil {
			http.Error(w, "Error occurred while trying to unmarshal JSON data", http.StatusInternalServerError)
			return
		}

		err = tpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Error occurred while trying to render template", http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
