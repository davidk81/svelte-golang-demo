package patient

import (
	"database/sql"

	"github.com/davidk81/svelte-golang-demo/backend/db/models"
	"github.com/valyala/fasthttp"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// var patients = []Patient{{PatientID: "patient1", Name: "Peter"}, {PatientID: "patient2", Name: "Paul"}, {PatientID: "patient3", Name: "Patrick"}}

// GetPatients (list)
func GetPatients(ctx *fasthttp.RequestCtx) (models.PatientSlice, error) {
	// Open handle to database like normal
	db, err := sql.Open("postgres", "host=localhost dbname=patientdb user=docker password=docker sslmode=disable")
	if err != nil {
		return nil, err
	}

	return models.Patients(qm.Limit(5)).All(ctx, db)
}
