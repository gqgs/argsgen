package options

import "time"

type options struct {
	i, input     string        `arg:"input filename,positional"`
	o, output    string        `arg:"output filename,positional"`
	db, database string        `arg:"database name"`
	folder       string        `arg:"target folder,required"`
	parallel     uint          `arg:"number of process in parallel"`
	limit        int           `arg:"limit of something,required"`
	real         float64       `arg:"float of something"`
	profile      bool          `arg:"should it profile?"`
	duration     time.Duration `arg:"some duration,required"`
}
