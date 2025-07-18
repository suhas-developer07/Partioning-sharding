package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
			 query := `INSERT INTO user_analytics_p5 (id,session_time,country) VALUES($1,$2,$3)`

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


	router.HandleFunc("/{id}",func(w http.ResponseWriter, r *http.Request) {
	 vars := mux.Vars(r)
	 Id := vars["id"]
	 log.Print(Id)

	 userId ,_ :=strconv.Atoi(Id)

	 log.Print(userId)

	 var data Payload

	 switch  {

	 case userId<=10 :
		query := `SELECT session_time,country FROM user_analytics_p1 WHERE id=$1`

		err := db.Shard1.QueryRow(query,userId).Scan(
			&data.Session_Time,
			&data.Country,
		)
		if err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"msg":"unable to get the data",
					"error":err.Error(),
				})
				return 
		}

	     w.WriteHeader(http.StatusOK)
		 json.NewEncoder(w).Encode(data)
		 
	case userId<=20 :
		query := `SELECT session_time,country FROM user_analytics_p2 WHERE id=$1`

		err := db.Shard1.QueryRow(query,userId).Scan(
			&data.Session_Time,
			&data.Country,
		)
		if err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"msg":"unable to get the data",
				})
				return 
		}

	     w.WriteHeader(http.StatusOK)
		 json.NewEncoder(w).Encode(data)	

	 case userId<=30 :
		query := `SELECT session_time,country FROM user_analytics_p3 WHERE id=$1`

		err := db.Shard1.QueryRow(query,userId).Scan(
			&data.Session_Time,
			&data.Country,
		)
		if err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"msg":"unable to get the data",
				})
				return 
		}

	     w.WriteHeader(http.StatusOK)
		 json.NewEncoder(w).Encode(data)	
	
	 case userId<=40 :
		query := `SELECT session_time,country FROM user_analytics_p4 WHERE id=$1`

		err := db.Shard2.QueryRow(query,userId).Scan(
			&data.Session_Time,
			&data.Country,
		)
		if err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"msg":"unable to get the data",
				})
				return 
		}

	     w.WriteHeader(http.StatusOK)
		 json.NewEncoder(w).Encode(data)	

	 case userId<=50 :
		query := `SELECT session_time,country FROM user_analytics_p5 WHERE id=$1`

		err := db.Shard2.QueryRow(query,userId).Scan(
			&data.Session_Time,
			&data.Country,
		)
		if err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]string{
					"msg":"unable to get the data",
				})
				return 
		}

	     w.WriteHeader(http.StatusOK)
		 json.NewEncoder(w).Encode(data)	
	 }

	}).Methods("GET")



	return router
}