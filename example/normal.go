package normal

import (
	"time"
)

type Normal struct {
	Date time.Time `toml:"date"`
	Name string    `toml:"name"`
}
