package options

type options struct {
	i, input     string  `arg:"input filename,+"`
	o, output    string  `arg:"output filename,+"`
	db, database string  `arg:"database name"`
	folder       string  `arg:"target folder,required"`
	parallel     uint    `arg:"number of process in parallel"`
	limit        int     `arg:"limit of something"`
	real         float64 `arg:"float of something"`
	profile      bool    `arg:"should it profile?"`
}
