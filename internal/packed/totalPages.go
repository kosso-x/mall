package packed

import (
	"context"
)

func TotalPages(ctx context.Context, count, size int) (totalPages int) {
	r1 := count / size
	r2 := count % size
	if r2 > 0 {
		totalPages = r1 + 1
	} else {
		totalPages = r1
	}

	return
}
