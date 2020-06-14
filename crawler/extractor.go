package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/imthaghost/goclone/parser"
)

var (
	extensionDir = map[string]string{
		".css":  "css",
		".js":   "js",
		".jpg":  "imgs",
		".jpeg": "imgs",
		".gif":  "imgs",
		".png":  "imgs",
		".svg":  "imgs",
	}
)

// Extractor visits a link determines if its a page or sublink
// downloads the contents to a correct directory in project folder
func Extractor(link string, projectPath string) {
	fmt.Println("Extracting --> ", link)

	// get the html body
	resp, err := http.Get(link)
	if err != nil {
		panic(err)
	}

	// Closure
	defer resp.Body.Close()
	// file base
	base := parser.URLFilename(link)
	// file extension
	ext := parser.URLExtension(link)

	// checks if there was a valid extension
	if ext != "" {
		// checks if that extension has a directory path name associated with it
		// from the extensionDir map
		dirPath := extensionDir[ext]
		if dirPath != "" {
			// If extension and path are valid, move on to writeFileToPath
			writeFileToPath(projectPath, base, ext, ext, dirPath, resp)
		}
	}
}

func writeFileToPath(projectPath, base, oldFileExt, newFileExt, fileDir string, resp *http.Response) {
	var name = base[0 : len(base)-len(oldFileExt)]
	document := name + newFileExt

	// get the project name and path we use the path to
	f, err := os.OpenFile(projectPath+"/"+fileDir+"/"+document, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	htmlData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	f.Write(htmlData)
}
