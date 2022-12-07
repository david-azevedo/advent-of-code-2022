package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

func newFile(name string, size int) *file {
	f := file{name, size}
	return &f
}

type dir struct {
	name   string
	parent *dir
	files  []*file
	dirs   []*dir
}

func newDir(name string, parent *dir) *dir {
	d := dir{name, parent, []*file{}, []*dir{}}
	return &d
}

func (d *dir) addFile(f *file) {
	d.files = append(d.files, f)
}

func (d *dir) addDir(newD *dir) {
	d.dirs = append(d.dirs, newD)
}

func (d *dir) searchDir(name string) *dir {
	for _, dir := range d.dirs {
		if dir.name == name {
			return dir
		}
	}
	fmt.Printf("Did not find dir %v inside %v", name, d.name)
	return nil
}

func main() {
	file, err := os.Open("./day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	root := newDir("/", nil)
	currentDir := root
	for scanner.Scan() { // Process input and build file system
		tokens := strings.Split(scanner.Text(), " ")

		switch tokens[0] {
		case "$":
			switch tokens[1] {
			case "cd":
				switch tokens[2] {
				case "..":
					currentDir = currentDir.parent
				default:
					dirName := tokens[2]
					currentDir = currentDir.searchDir(dirName)
				}
			default:
			}
		case "dir": // dir
			name := tokens[1]
			newDir := newDir(name, currentDir)
			currentDir.addDir(newDir)
		default: // file
			size, _ := strconv.Atoi(tokens[0])
			name := tokens[1]
			file := newFile(name, size)
			currentDir.addFile(file)
		}
	}

	// Debug file system
	// printFileSystem(root, 0)

	// Part 1
	sum := 0
	traversePart1(root, &sum)
	fmt.Println(sum)

	//Part 2

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func traversePart1(d *dir, sum *int) {
	max := 100000
	size := calcDirSize(d)
	if size <= max {
		*sum += size
	}
	for _, dir := range d.dirs {
		traversePart1(dir, sum)
	}
}

func calcDirSize(d *dir) int {
	size := 0
	for _, file := range d.files {
		size += file.size
	}

	for _, dir := range d.dirs {
		size += calcDirSize(dir)
	}

	return size
}

func printFileSystem(root *dir, tabs int) {
	fmt.Println(strings.Repeat(" ", tabs), "-", root.name, "(dir)")
	for _, file := range root.files {
		fmt.Println(strings.Repeat(" ", tabs+1), "-", file.name, "(file, size=", file.size, ")")
	}
	for _, dir := range root.dirs {
		printFileSystem(dir, tabs+1)
	}
}
