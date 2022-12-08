package main

import (
	"log"
	"os"
	"sort"
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
var TerminalLines = []string{}
var TerminalIndex = 0

// Sizes of directories at or under 100000
var constrainedDirectorySizes = []int{}
var directorySizes = []int{}

func main() {
	DirectoryTree.Name = "Root"

	content, err := os.ReadFile("PuzzleInput.txt")
	if err != nil {
		log.Fatal(err)
	}

	TerminalLines = strings.Split(string(content), "\r\n")

	for TerminalIndex < len(TerminalLines) {
		executeCommand(TerminalLines[TerminalIndex])
	}

	getDirectorySizes(DirectoryTree)

	// Part 1
	directorySizeSum := 0
	for _, size := range constrainedDirectorySizes {
		directorySizeSum += size
	}

	log.Printf("Total size of directories: %d\n", directorySizeSum)

	// Part 2
	sort.Slice(directorySizes, func(i, j int) bool {
		return directorySizes[i] < directorySizes[j]
	})

	smallestPossibleDirectorySize := getSmallestPossibleDirectorySize()

	log.Printf("Size of smallest possible directory to delete to free up enough space: %d\n", smallestPossibleDirectorySize)
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
	} else /* It's an ls */ {
		items := getDirectoryItems()
		populateDirectory(items)
	}

}

func getChildNode(nodeName string) (childNode *DirectoryNode) {
	for _, node := range CurrentNode.Directories {
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
		items = append(items, TerminalLines[TerminalIndex])
		TerminalIndex++

		if TerminalIndex >= len(TerminalLines) {
			break
		}
	}

	return
}

func isCommandLine() (isCommandLine bool) {
	isCommandLine = string(TerminalLines[TerminalIndex][0]) == "$"
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
		constrainedDirectorySizes = append(constrainedDirectorySizes, size)
	}

	directorySizes = append(directorySizes, size)

	return
}

func getSmallestPossibleDirectorySize() (directorySize int) {
	totalDiskSpace := 70000000
	updateSize := 30000000
	usedSpace := directorySizes[len(directorySizes)-1]
	freeSpace := totalDiskSpace - usedSpace
	requiredDirectorySize := updateSize - freeSpace

	// There's probably some magic to do this somewhere, but I don't know it yet
	index := 0
	for directorySize < requiredDirectorySize {
		directorySize = directorySizes[index]

		index++
	}

	return
}
