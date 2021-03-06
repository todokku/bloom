package contacts

import (
	"time"

	"gitlab.com/bloom42/bloom/core/db"
)

func UpdateContact(contact Contact) (Contact, error) {
	// TODO: validators
	cleanContactCollections(&contact)
	var err error
	now := time.Now().UTC()

	contact.UpdatedAt = now

	query := `
		UPDATE contacts SET
			updated_at = $1,
			first_name = $2,
			last_name = $3,
			notes = $4,
			addresses = $5,
			birthday = $6,
			organizations = $7,
			emails = $8,
			phones = $9,
			websites = $10,
			device_id = $11,
			bloom_username = $12
		WHERE id = $13
	`
	_, err = db.DB.Exec(query,
		contact.UpdatedAt,
		contact.FirstName,
		contact.LastName,
		contact.Notes,
		contact.Addresses,
		contact.Birthday,
		contact.Organizations,
		contact.Emails,
		contact.Phones,
		contact.Websites,
		contact.DeviceID,
		contact.BloomUsername,
		contact.ID,
	)

	return contact, err
}
