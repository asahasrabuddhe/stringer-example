//go:generate stringer -type=RequestSource

package main

import "fmt"

type RequestSource int

const (
	// Unique Request Sources
	Make RequestSource = iota
	ProjectBoard
	Stories
	YSC
	// Aliased Request Sources
	MakeApp = Make
	ProjectBoardApp = ProjectBoard
	StoriesApp = Stories
)

func main() {
	fmt.Println(Make, MakeApp)
	fmt.Println(ProjectBoard, ProjectBoardApp)
	fmt.Println(Stories, StoriesApp)
	fmt.Println(YSC)
}
