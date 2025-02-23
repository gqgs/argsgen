package options

type options struct {
	db     string `arg:"database name"`
	limit  int    `arg:"max of characters to use"`
	song   string `arg:"song that will be used in the video"`
	frame  int    `arg:"length of each frame"`
	folder string `arg:"folder which output will be uploaded"`
	debug  bool   `arg:"enable debug logs"`
}
