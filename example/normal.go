package normal

import (
	"time"
)

type Normal struct {
	Date time.Time `toml:"date"`
	Name string    `toml:"name"`
	Num  int64     `toml:"num"`
}
