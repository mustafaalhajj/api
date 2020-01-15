package main
import(
	"net/http"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)
type Roll struct {
	ID           string `json:"id"`
	ImageNumber string `json:" imageNumber "`
	Name         string `json:" name"`
	Ingredients  string `json:"ingredients"`
}

func getRolls(w http.ResponseWriter,r*http.Request){
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(rolls)
}
func getRoll(w http.ResponseWriter,r*http.Request){
	w.Header().Set("content-type", "application/json")
	params :=mux.Vars(r)
	for _,item:=range rolls {
	if item.ID == params["id"]{
	json.NewEncoder(w).Encode(item)
	return
		}
	}
  }

func createRoll(w http.ResponseWriter,r*http.Request){
	w.Header().Set("content-type", "application/json")
	
	var newRoll Roll
	json.NewDecoder(r.Body).Decode(&newRoll)
	newRoll.ID=strconv.Itoa(len(rolls)+1)
	rolls = append(rolls,newRoll)

	json.NewEncoder(w).Encode(newRoll)	
}
func updateRoll(w http.ResponseWriter,r*http.Request){
	w.Header().Set("content-type", "application/json")
	params :=mux.Vars(r)
	for i, item:=range rolls{
	if item.ID == params["id"]{
	rolls=append(rolls[:i], rolls[i+1:]...)
	var newRoll Roll
	json.NewDecoder(r.Body).Decode(&newRoll)
	newRoll.ID = params["id"]
	rolls = append(rolls,newRoll){
	json.NewEncoder(w).Encode(newRoll)
	return
		}
	}
}
func deleteRoll(w http.ResponseWriter,r*http.Request){
	w.Header().Set("content-type", "application/json")
	params :=mux.Vars(r)
	for i, item:=range rolls{
		if item.ID == params["id"]{
		rolls=append(rolls[:i], rolls[i+1:]...)
		break
		}
	}
	json.NewEncoder(w).Encode(newRoll)	
}
var rolls []Roll

func main() {

	rolls = append(rolls, Roll{ID: "1", ImageNumber: "8", Name: 		"Spicy Tuna Roll", Ingredients: "Tuna, Chili sauce, Nori, 		Rice"})

	router := mux.NewRouter();

	router.HandleFunc("/sushi",getRolls).Methods("get")
	router.HandleFunc("/sushi{id}",getRoll).Methods("get")
	router.HandleFunc("/sushi",createRoll).Methods("POST")
	router.HandleFunc("/sushi{id}",updateRoll).Methods("POST")
	router.HandleFunc("/sushi{id}",getRoll).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":5000",router))

}
