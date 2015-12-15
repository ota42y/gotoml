package normal

import (
	"time"
)

type Normal struct {
	Date    time.Time   `toml:"date"`
	Dates   []time.Time `toml:"dates"`
	Name    string      `toml:"name"`
	Num     int64       `toml:"num"`
	Numbers []int64     `toml:"numbers"`
}
