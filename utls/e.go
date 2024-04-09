package utls

import "os"

var (
	LocalBaseURL string = os.Getenv("LOCAL_BASE_URL")
)
