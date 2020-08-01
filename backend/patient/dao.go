package patient

import (
	"github.com/davidk81/svelte-golang-demo/backend/patientdb"
	"github.com/davidk81/svelte-golang-demo/backend/patientdb/models"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// var patients = []Patient{{PatientID: "patient1", Name: "Peter"}, {PatientID: "patient2", Name: "Paul"}, {PatientID: "patient3", Name: "Patrick"}}

// GetPatients (list)
func GetPatients(ctx *fasthttp.RequestCtx) (models.PatientSlice, error) {
	return models.Patients(qm.Limit(5)).All(ctx, patientdb.DB())
}

// AddPatientNote adds a new note to the patient
func AddPatientNote(note *models.PatientNote, ctx *fasthttp.RequestCtx) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	note.Noteid = id.String()
	return (*note).Insert(ctx, patientdb.DB(), boil.Infer())
}
