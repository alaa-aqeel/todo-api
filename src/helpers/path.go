package helpers

import (
	"os"
	"regexp"
)

func BaseDir(name string) string {
	regName := regexp.MustCompile(`^(.*` + name + `)`)
	currentDir, _ := os.Getwd()

	return string(regName.Find([]byte(currentDir)))
}
