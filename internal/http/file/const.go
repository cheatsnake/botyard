package file

const filesFolder = "stock"
const maxFilesAmount = 10
const fileFieldName = "file"

var maxFileSizes = map[string]int64{
	"images": 2 * 1024 * 1024,  // 2 MB
	"audios": 5 * 1024 * 1024,  // 5 MB
	"videos": 25 * 1024 * 1024, // 25 MB
	"files":  10 * 1024 * 1024, // 10 MB
}

var knownContentTypes = map[string]string{
	"image/gif":     "images",
	"image/jpeg":    "images",
	"image/png":     "images",
	"image/svg+xml": "images",
	"image/webp":    "images",

	"video/mp4":       "videos",
	"video/webm":      "videos",
	"video/ogg":       "videos",
	"video/quicktime": "videos",
	"video/x-flv":     "videos",

	"audio/mpeg": "audios",
	"audio/ogg":  "audios",
	"audio/wav":  "audios",
	"audio/aac":  "audios",
}
