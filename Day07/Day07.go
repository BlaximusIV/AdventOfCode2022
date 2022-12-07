package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type File struct {
	Name string
	Size int
}

type DirectoryNode struct {
	Name            string
	ParentDirectory *DirectoryNode
	Directories     []*DirectoryNode
	Files           []File
}

var DirectoryTree = DirectoryNode{}
var CurrentNode = &DirectoryTree
var TerminalInteractions = []string{}
var TerminalIndex = 0

// Sizes of directories at or under 100000
var directorySizes = []int{}

func main() {
	DirectoryTree.Name = "Root"

	content, err := os.ReadFile("PuzzleInput.txt")
	if err != nil {
		log.Fatal(err)
	}

	TerminalInteractions = strings.Split(string(content), "\r\n")

	for TerminalIndex < len(TerminalInteractions) {
		executeCommand(TerminalInteractions[TerminalIndex])
	}

	getDirectorySizes(DirectoryTree)

	directorySizeSum := 0
	for _, size := range directorySizes {
		directorySizeSum += size
	}

	log.Printf("Total size of directories: %d\n", directorySizeSum)
}

func executeCommand(command string) {
	commandParts := strings.Split(command, " ")

	// if it's a cd
	if len(commandParts) == 3 {
		parameter := commandParts[2]
		if parameter == "/" {
			CurrentNode = &DirectoryTree
			TerminalIndex++
		} else if parameter == ".." {
			CurrentNode = CurrentNode.ParentDirectory
			TerminalIndex++
		} else {
			CurrentNode = getChildNode(parameter)
			TerminalIndex++
		}
		// It's an ls
	} else {
		items := getDirectoryItems()
		populateDirectory(items)
	}

}

func getChildNode(nodeName string) (childNode *DirectoryNode) {
	for _, node := range CurrentNode.Directories {
		// Names are unique
		if node.Name == nodeName {
			childNode = node
			break
		}
	}

	return
}

func getDirectoryItems() (items []string) {
	TerminalIndex++

	// append until we get to the next command

	for !isCommandLine() {
		items = append(items, TerminalInteractions[TerminalIndex])
		TerminalIndex++

		if TerminalIndex >= len(TerminalInteractions) {
			break
		}
	}

	return
}

func isCommandLine() (isCommandLine bool) {
	isCommandLine = string(TerminalInteractions[TerminalIndex][0]) == "$"
	return
}

func populateDirectory(directoryItems []string) {
	for _, item := range directoryItems {
		itemParts := strings.Split(item, " ")

		// If it's a directory
		if itemParts[0] == "dir" {
			subDirectory := DirectoryNode{}
			subDirectory.ParentDirectory = CurrentNode
			subDirectory.Name = itemParts[1]
			CurrentNode.Directories = append(CurrentNode.Directories, &subDirectory)
		} else {
			size, _ := strconv.Atoi(itemParts[0])
			CurrentNode.Files = append(CurrentNode.Files, File{itemParts[1], size})
		}
	}
}

func getDirectorySizes(node DirectoryNode) (size int) {
	for _, file := range node.Files {
		size += file.Size
	}

	for _, childNode := range node.Directories {
		size += getDirectorySizes(*childNode)
	}

	if size <= 100000 {
		directorySizes = append(directorySizes, size)
	}

	return
}
