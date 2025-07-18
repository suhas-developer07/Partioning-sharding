package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/suhas-developer07/Partioning-sharding/db"
)

type Payload struct{
	Id           int    `json:"id"`
	Session_Time int    `json:"session_time"`
	Country      string `json:"country"`
}

func MountRoutes()*mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		var Payload Payload

			if err:= json.NewDecoder(r.Body).Decode(&Payload);err!=nil{
				w.Header().Set("Content-Type","application/json")
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"msg":"unable to read body",
				})
				return 
			}

		   switch {
		   case Payload.Id<=10 :
			 query := `INSERT INTO user_analytics_p1 (id,session_time,country) VALUES($1,$2,$3)`

			 _,err := db.Shard1.Exec(query,Payload.Id,Payload.Session_Time,Payload.Country)
			 if err!=nil{
				log.Fatalf("database insertion failed")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"msg":"Database insertion failed",
				})
				return 
			 }
			 w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]string{
					"msg":"Insertion completed",
				})	 
			case Payload.Id<=20 :
			 query := `INSERT INTO user_analytics_p2 (id,session_time,country) VALUES($1,$2,$3)`

			 _,err := db.Shard1.Exec(query,Payload.Id,Payload.Session_Time,Payload.Country)
			 if err!=nil{
				log.Fatalf("database insertion failed")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"msg":"Database insertion failed",
				})
				return 
			 }
			 w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]string{
					"msg":"Insertion completed",
				})	 
			
			case Payload.Id<=30 :
			 query := `INSERT INTO user_analytics_p3 (id,session_time,country) VALUES($1,$2,$3)`

			 _,err := db.Shard1.Exec(query,Payload.Id,Payload.Session_Time,Payload.Country)
			 if err!=nil{
				log.Fatalf("database insertion failed")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"msg":"Database insertion failed",
				})
				return 
			 }
			 w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]string{
					"msg":"Insertion completed",
				})	

		    case Payload.Id<=40 :
			 query := `INSERT INTO user_analytics_p4 (id,session_time,country) VALUES($1,$2,$3)`

			 _,err := db.Shard2.Exec(query,Payload.Id,Payload.Session_Time,Payload.Country)
			 if err!=nil{
				log.Fatalf("database insertion failed")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"msg":"Database insertion failed",
				})
				return 
			 }
			 w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]string{
					"msg":"Insertion completed",
				})	 
			 
			 case Payload.Id<=50 :
			 query := `INSERT INTO user_analytics_p4 (id,session_time,country) VALUES($1,$2,$3)`

			 _,err := db.Shard2.Exec(query,Payload.Id,Payload.Session_Time,Payload.Country)
			 if err!=nil{
				log.Fatalf("database insertion failed")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"msg":"Database insertion failed",
				})
				return 
			 }
			 w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]string{
					"msg":"Insertion completed",
				})	
		   }
	}).Methods("POST")



	return router
}