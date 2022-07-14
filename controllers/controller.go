package controllers

import (
	"context"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/honestbank/tech-assignment-backend-engineer/manager"
	"github.com/honestbank/tech-assignment-backend-engineer/models"
	"io/ioutil"
	"log"
	"net/http"
)

type controller struct {
	ctx context.Context
	m   manager.Manager
}

func NewController(context context.Context, mgr manager.Manager) *controller {
	return &controller{ctx: context, m: mgr}
}
func (c *controller) ProcessData(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		body, _ := ioutil.ReadAll(req.Body)
		var recordData models.RecordData
		json.Unmarshal(body, &recordData)

		response := c.m.AddCreditDetails(&recordData)
		resp.WriteHeader(http.StatusOK)
		resp.Header().Set("Content-Type", "application/json")
		json.NewEncoder(resp).Encode(response)
	default:
		log.Println("error no 404")
		resp.WriteHeader(http.StatusNotFound)
		fmt.Fprint(resp, "not found")
	}

}
