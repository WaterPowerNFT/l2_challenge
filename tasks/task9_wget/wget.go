package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func getSites() []string {
	return os.Args[1:]
}

func downloadSite(index int, siteToDownload string) error {
	resp, err := http.Get(siteToDownload)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	f, err := os.Create(fmt.Sprintf("download/%d.html", index))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return err
}

func main() {
	sites := getSites()
	err := os.MkdirAll("download", 0777)
	if err != nil {
		panic(err)
	}
	for i, site := range sites {
		err := downloadSite(i, site)
		if err != nil {
			fmt.Printf("error while proccessing web site %v\n", err)
		} else {
			fmt.Printf("web site %s downloaded\n", site)
		}
	}
	fmt.Println("job is done")
}
