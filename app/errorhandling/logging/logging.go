package logging

// import "fmt"
// import "log"
// import "os"
import (
	"github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

// func LearnLogging() {
// 	log.SetPrefix("LearnLogging(): ")
// 	log.Print("I'm a log")
// 	log.Panic("Heeeello log!")
// 	fmt.Println("Can you see me?")
// }

// func RecordLogging() {
// 	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer file.Close()

// 	log.SetOutput(file)
// 	log.Print("Hey, I'm a log!!!")
// }

func RunZeroLog() {
    zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Debug().
		Int("EmployeeID", 1001).
		Msg("Getting employee information")

}