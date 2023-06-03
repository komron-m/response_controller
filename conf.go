package resonse_controller

import (
	"time"
)

// deadline of 2 seconds everywhere
func deadline() time.Time {
	return time.Now().Add(time.Second * 2)
}
