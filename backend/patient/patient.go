package patient

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

// Patient struct for (de)marshalling
type Patient struct {
	PatientID string `json:"patientId"`
	Name      string `json:"name"`
}

var patients = []Patient{{PatientID: "patient1", Name: "Peter"}, {PatientID: "patient2", Name: "Paul"}, {PatientID: "patient3", Name: "Patrick"}}

// HandlePatient entrypoint http request handler
func HandlePatient(ctx *fasthttp.RequestCtx) {

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
	b, err := json.Marshal(patients)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx.SetBody([]byte(b))
	ctx.SetStatusCode(fasthttp.StatusCreated)
}
