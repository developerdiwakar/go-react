package helpers

import "os"

const (
	MAILS_DIR_PATH string = "/app/views/mails/"
)

// Get the the file path for
func GetMailTemplateFilePath(filename string) string {
	// Get the Current working diretory
	cwd, _ := os.Getwd()
	extn := "html"
	return cwd + MAILS_DIR_PATH + filename + "." + extn
}

// The below function checks if a regular file (not directory)
// with a given filepath exist
func IsFileExist(filePath string) bool {
	fileInfo, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		return false
	}
	// will return false if the fileinfo says the filepath is a directory
	return !fileInfo.IsDir()
}
