package pkg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// musixmatch API base URL
const BASE_URL string = "http://api.musixmatch.com/ws/1.1/"

// your api key
// created here - https://developer.musixmatch.com/admin
// requires login
var API_KEY string = os.Getenv("MUSIX_MATCH_API_KEY")

func init() {
	// initialize environment variables
	if envErr := godotenv.Load(); envErr != nil {
		log.Fatal("Error loading .env file")
	}
}
