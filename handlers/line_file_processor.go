package handlers

import "sync"

func Run() {
	/*
		fileName := configurations.LineFileFullpath
		processStartTime := time.Now()
		fileContent, err := ioutil.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		lineContent := strings.Split(string(fileContent), "\n")
		lines := make([]string, len(lineContent))
		copy(lines, lineContent[:])
		fmt.Println(len(lines))
		fmt.Println(cap(lines))
		chunkLen := configurations.FileChunkSize
		chunkStart := 0
		chunkEnd := chunkLen
		quantityChunk := len(lines) / chunkLen
		var wg sync.WaitGroup
		wg.Add(quantityChunk)
		for i := 1; i <= quantityChunk; i++ {
			sendDataToSDL(lines, chunkStart, chunkEnd, &wg) //, &wg
			chunkStart = chunkEnd
			chunkEnd = chunkLen * i
		}
		fmt.Println(lines[chunkEnd : len(lines)-1])
		wg.Wait()
		processEndTime := time.Now()
		fmt.Println(processStartTime)
		fmt.Println(processEndTime)
		fmt.Println(quantityChunk)
	*/
}

func sendDataToSDL(lines []string, chunkStart int, chunkEnd int, wg *sync.WaitGroup) {
	/*
		defer wg.Done()
		lineChunk := lines[chunkStart:chunkEnd]
		var line string
		for _, line = range lineChunk {
			r := models.SDLRequest{Data: line}
			_, err := resty.R().
				SetHeader("Content-Type", "application/json").
				SetBody(r).
				Post(configurations.SDLEndpoint)
			if err != nil {
				fmt.Println(fmt.Sprintf("Error ingesting line: %s", line))
			}
		}
	*/
}
