package curtime

import (
	"time"

	"github.com/beevik/ntp"
)

func out_curtime() (time.Time, error) {

	return ntp.Time("0.beevik-ntp.pool.ntp.org")
}
