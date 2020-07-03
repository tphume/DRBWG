package reminder

import "time"

// Represent connection to data source
type InsertRepo interface {
	Insert(t time.Time, name string, g string, c string) error
}
