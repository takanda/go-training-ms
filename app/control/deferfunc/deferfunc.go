package deferfunc

import "fmt"
import "io"
import "os"

func Deferfunc() {
	newfile, error := os.Create("learnGo.txt")
	if error != nil {
		fmt.Println("Error: Could not create file.")
		return
	}
	
	defer newfile.Close()
	
	if _, error = io.WriteString(newfile, "Learning Go!!!"); error != nil {
		fmt.Println("Error: Could not write to file.")
		return
	}

	newfile.Sync()
}