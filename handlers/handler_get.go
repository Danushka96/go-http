package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/pickme-go/log"
	"io/ioutil"
	"net/http"
)

type HandleGet struct {

}

type getRequestBody struct {
	Id int64 `json:"id,omitempty"`
}

func (HandleGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body);
	if err!=nil{
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, err)
	}

	var p getRequestBody

	err = json.Unmarshal(data, &p)

	if err!=nil{
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, err)
	}

	var person = PersonMap[p.Id]

	if person.Id==0{
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, "person not found")
		return
	}

	response,_ := json.Marshal(person)

	_,_ = w.Write(response)

}
