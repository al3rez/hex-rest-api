package mem

import "time"

type User struct {
	ID      int
	Name    string
	Created time.Time
}
