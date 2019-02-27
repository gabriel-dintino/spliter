package configurations

import (
	"os"
	"strconv"
)

var Port, LineFileFullpath, SDLEndpoint, fileChunkSizeStr string
var FileChunkSize int

func init() {
	if LineFileFullpath = os.Getenv("LINE_INGESTOR_LINE_FILE_URI"); LineFileFullpath == "" {
		LineFileFullpath = "C:/files/SRV_LINEAS.DAT"
	}
	if Port = os.Getenv("LINE_INGESTOR_SERVER_PORT"); Port == "" {
		Port = "8080"
	}
	if SDLEndpoint = os.Getenv("LINE_INGESTOR_SDL_ENDPOINT"); SDLEndpoint == "" {
		SDLEndpoint = "http://lines-server-int-core-lab.apps-ocp1.cloudtelecom.com.ar/sdl/v1/lines"
	}
	if fileChunkSizeStr = os.Getenv("LINE_INGESTOR_CHUNK_FILE_SIZE"); fileChunkSizeStr == "" {
		FileChunkSize = 10000
	} else {
		var err error
		FileChunkSize, err = strconv.Atoi(fileChunkSizeStr)
		if err != nil {
			panic("Wrong chunk file size environment variable was set.")
		}
	}
}
