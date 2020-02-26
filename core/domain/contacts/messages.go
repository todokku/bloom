package contacts

import (
	"time"
)

type Contacts struct {
	Contacts []Contact `json:"contacts"`
}

type CreateContactParams struct {
	DeviceID      string        `json:"device_id"`
	FirstName     string        `json:"first_name"`
	LastName      string        `json:"last_name"`
	Notes         string        `json:"notes"`
	Birthday      *time.Time    `json:"birthday"`
	Organizations Organizations `json:"organizations"`
	Addresses     Addresses     `json:"addresses"`
	Emails        Emails        `json:"emails"`
	Phones        Phones        `json:"phones"`
	Websites      Websites      `json:"websites"`
}

type DeleteContactParams struct {
	ID string `json:"id"`
}
