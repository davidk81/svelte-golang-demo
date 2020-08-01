package patient

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/davidk81/svelte-golang-demo/backend/session"
	"github.com/valyala/fasthttp"
)

//Patient struct for (de)marshalling
type Patient struct {
	PatientID string `json:"patientId"`
	Name      string `json:"name"`
}

// HandlePatient entrypoint http request handler
func HandlePatient(ctx *fasthttp.RequestCtx) {
	session.VerifySession(ctx, "nurse")
	switch string(ctx.Request.Header.Method()) {
	case "POST":
		handleMethodPost(ctx)
	case "DELETE":
		handleMethodDelete(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}

// HandlePatientList entrypoint http request handler
func HandlePatientList(ctx *fasthttp.RequestCtx) {
	session.VerifySession(ctx, "nurse")
	switch string(ctx.Request.Header.Method()) {
	case "GET":
		handleMethodGetList(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}

func handleMethodDelete(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func handleMethodPost(ctx *fasthttp.RequestCtx) {
	// decode post body
	var patient Patient
	err := json.Unmarshal(ctx.Request.Body(), &patient)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	// TODO: validate data
	// TODO: insert to db

	// return patient info in response
	b, err := json.Marshal(patient)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx.SetBody([]byte(b))
	ctx.SetStatusCode(fasthttp.StatusCreated)
}

func handleMethodGetList(ctx *fasthttp.RequestCtx) {
	// return patient info in response
	p, err := GetPatients(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := json.Marshal(p)
	log.Println(string(b))
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx.SetBody([]byte(b))
	ctx.SetStatusCode(fasthttp.StatusCreated)
}
