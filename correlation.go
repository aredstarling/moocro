package moocro

import (
	"github.com/gofrs/uuid"
)

func generateCorrelationID() (string, error) {
	v4uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return uuid.NewV5(v4uuid, "moocro").String(), nil
}
