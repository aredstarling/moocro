package moocro

import (
	"os"
	"strconv"
)

func concurrency() int {
	value, err := strconv.Atoi(os.Getenv("MOOCRO_CONCURRENCY"))
	if err != nil || value < 1 {
		return 1
	}

	return value
}
