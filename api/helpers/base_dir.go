package helpers

import (
	"os"
	"regexp"
)

func GetBaseDir(name string) string {
	regName := regexp.MustCompile(`^(.*` + name + `)`)
	currentDir, _ := os.Getwd()
	return string(regName.Find([]byte(currentDir)))
}
