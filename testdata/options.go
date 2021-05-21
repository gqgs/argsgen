package options

type options struct {
	i, input     string `arg:"input filename,+"`
	o, output    string `arg:"output filename,+"`
	db, database string `arg:"database name"`
	folder       string `arg:"target folder"`
	parallel     uint   `arg:"number of process in parallel"`
	profile      bool
}
