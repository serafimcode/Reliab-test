package dto

import (
	"github.com/google/uuid"
	"time"
)

type MeetingDTO struct {
	ID              uuid.UUID `json:"uuid"`
	Name            string    `json:"name"`
	Place           string    `json:"place"`
	Comment         *string   `json:"comment,omitempty"`
	RecipientEmails [5]string `json:"recipient_emails,omitempty"`
	ApplicantEmail  string    `json:"applicant_email"`
	StartDate       time.Time `json:"start_date"` // dd.mm.yyyy
	StartTime       time.Time `json:"start_time"` // hh:mm
	EndDate         time.Time `json:"end_date"`   // dd.mm.yyyy
	EndTime         time.Time `json:"end_time"`   // hh:mm
	IsFullDay       bool      `json:"is_full_day"`
	IsOnline        bool      `json:"is_online"`
	AuthorEmail     string    `json:"author_email"`
}

type GetMeetingResponse struct {
	MeetingDTO
}

type CreateMeetingRequest struct {
	MeetingDTO
}

type EditMeetingRequest struct {
	MeetingDTO
}
