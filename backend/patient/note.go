package patient

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/davidk81/svelte-golang-demo/backend/patientdb/models"
	"github.com/davidk81/svelte-golang-demo/backend/session"
	"github.com/valyala/fasthttp"
)

// HandlePatient entrypoint http request handler
func HandlePatientNote(ctx *fasthttp.RequestCtx) {
	session.VerifySession(ctx, "nurse")
	switch string(ctx.Request.Header.Method()) {
	case "POST":
		handleMethodNotePost(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}

// HandlePatientList entrypoint http request handler
func HandlePatientNoteList(ctx *fasthttp.RequestCtx) {
	session.VerifySession(ctx, "nurse")
	switch string(ctx.Request.Header.Method()) {
	case "GET":
		handleMethodNoteGetList(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}

func handleMethodNotePost(ctx *fasthttp.RequestCtx) {
	// decode post body
	var note models.PatientNote
	err := json.Unmarshal(ctx.Request.Body(), &note)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	err = AddPatientNote(&note, ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	// return patient info in response
	b, err := json.Marshal(note)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx.SetBody([]byte(b))
	ctx.SetStatusCode(fasthttp.StatusCreated)
}

func handleMethodNoteGetList(ctx *fasthttp.RequestCtx) {
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
