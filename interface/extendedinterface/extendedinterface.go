package extendedinterface

import (
    "fmt"
    "io"
    "net/http"
    "os"
	"encoding/json"
)

type customWriter struct{}

type GitHubResponse []struct {
	FullName string `json:"full_name"`
}

func (w customWriter) Write(p []byte) (n int, err error) {
	var res GitHubResponse
	json.Unmarshal(p, &res)
	for _, r := range res {
        fmt.Println(r.FullName)
    }
    return len(p), nil
}

func LearnExInterface() {
	res, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=3")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	writer := customWriter{}
	io.Copy(writer, res.Body)
}
