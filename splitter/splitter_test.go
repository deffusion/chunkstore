package splitter

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"testing"
)

func TestSplitter(t *testing.T) {
	file, err := os.Open("test.pdf")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if SplitIntoFiles(file) != nil {
		fmt.Println("split err:", err)
	}
}

func TestUserDir(t *testing.T) {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(currentUser.HomeDir)
}
