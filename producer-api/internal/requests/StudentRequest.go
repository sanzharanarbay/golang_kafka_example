package requests

import (
	"github.com/sanzharanarbay/golang_kafka_example/producer-api/internal/models"
	"net/url"
)

type StudentRequest struct {
	model *models.Student
}

func NewStudentRequest(model *models.Student) *StudentRequest {
	return &StudentRequest{
		model: model,
	}
}

func (s *StudentRequest) ValidateRequest() url.Values {
	errs := url.Values{}

	if s.model.ID == 0 {
		errs.Add("id", "The id field is required!")
	}

	if s.model.FIO == "" {
		errs.Add("fio", "The fio field is required!")
	}

	if len(s.model.FIO) < 5 || len(s.model.FIO) > 255 {
		errs.Add("fio", "The fio field must be between 3-255 chars!")
	}

	if s.model.Group == "" {
		errs.Add("group", "The group field is required!")
	}

	if s.model.Major == "" {
		errs.Add("major", "The major field is required!")
	}

	if s.model.Gpa == 0 {
		errs.Add("gpa", "The gpa field is required!")
	}
	return errs
}
