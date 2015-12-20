package normal

import (
	"time"
)

type Normal struct {
	Date    time.Time   `toml:"date"`
	Dates   []time.Time `toml:"dates"`
	Group   Group       `toml:"group"`
	Name    string      `toml:"name"`
	Num     int64       `toml:"num"`
	Numbers []int64     `toml:"numbers"`
}

type Group struct {
	Name   string `toml:"name"`
	School string `toml:"school"`
}
