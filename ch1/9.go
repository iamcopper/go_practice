package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	const header = "http://"
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(strings.ToLower(url), header) {
			url = header + url
			fmt.Printf("URL no header, after add header url: %s\n", url)
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("http response status: [%d] %s\n", resp.StatusCode, resp.Status)
		written, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copy %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("\n%d bytes written into output.\n", written)
	}
}
