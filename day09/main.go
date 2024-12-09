package main

import (
	"fmt"
	"os"

	"github.com/aoc-2024-go/io"
	"github.com/aoc-2024-go/utils"
)

var part1Answer, part2Answer int

func main() {
	args := os.Args[1:]
	var lines []string

	if len(args) > 0 {
		lines = io.ReadFile("./sample.txt")
	} else {
		lines = io.ReadFile("./input.txt")
	}
	line := lines[0]
	part1Answer = doPart1(line)
	part2Answer = doPart2(line)
	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
}

func doPart1(line string) int {
	var part1Checksum int
	var leftFileId, rightFileId int = 0, (len(line) + 1) / 2
	var curPos int = 0
	var rightFileBlocksRemaining int = 0
	var curSpaceRemaining int = 0

	for leftFileId < rightFileId {
		if curSpaceRemaining == 0 {
			// Need to advance to the next space, so add the checksum of the current file
			curFileSize := int(line[2*leftFileId] - '0')
			part1Checksum += checksum(curPos, curPos+curFileSize-1, leftFileId)

			// move curPos to the position of the next space
			curPos += curFileSize
			curSpaceRemaining = int(line[2*leftFileId+1] - '0')

			// Pre-emptively move to the next file
			leftFileId += 1
			continue
		}
		if rightFileBlocksRemaining == 0 {
			rightFileId -= 1
			// Need to consider the next file from the right to fragment
			rightFileBlocksRemaining = int(line[2*rightFileId] - '0')
			continue
		}

		// Figure out how many blocks from the right file to move
		var movedBlocks int
		if rightFileBlocksRemaining < curSpaceRemaining {
			movedBlocks = rightFileBlocksRemaining
		} else {
			movedBlocks = curSpaceRemaining
		}
		part1Checksum += checksum(curPos, curPos+movedBlocks-1, rightFileId)

		curPos += movedBlocks
		rightFileBlocksRemaining -= movedBlocks
		curSpaceRemaining -= movedBlocks
	}
	// The for loop exited while there might still be some blocks from the rightFile to place
	if rightFileBlocksRemaining > 0 {
		part1Checksum += checksum(curPos, curPos+rightFileBlocksRemaining-1, rightFileId)
	}
	return part1Checksum
}

func doPart2(line string) int {
	var nFiles = (len(line) + 1) / 2
	var filePos []int = make([]int, nFiles)
	var fileSizes []int = make([]int, nFiles)

	// The positions of each free space in the disk, grouped by size from 0 to 9. (Size zero free spaces are ignored)
	var spaces [][]int = make([][]int, 10)

	var part2Checksum int = 0
	curPos := 0
	for i := 0; i < len(line); i++ {
		curSize := int(line[i] - '0')
		if i%2 == 0 {
			// File
			curId := i / 2
			filePos[curId] = curPos
			fileSizes[curId] = curSize
		} else if curSize != 0 {
			// non-zero free space
			spaces[curSize] = append(spaces[curSize], curPos)
		}
		curPos += curSize
	}

	// Working backward
	for fileId := nFiles - 1; fileId >= 0; fileId-- {
		fileSize := fileSizes[fileId]
		if fileSize == 0 {
			// ignore zero size files; they do not contribute to the checksum
			continue
		}

		pos := filePos[fileId]
		spaceSize := -1

		// Find the leftmost space that can fit
		for size := fileSize; size < 10; size++ {
			if len(spaces[size]) > 0 && spaces[size][0] < pos {
				pos = spaces[size][0]
				spaceSize = size
			}
		}

		if spaceSize != -1 {
			// Move the file to the new position
			filePos[fileId] = pos
			// Recalculate the free space
			spaces[spaceSize] = spaces[spaceSize][1:]
			newSpaceSize := spaceSize - fileSize
			if newSpaceSize != 0 {
				newSpacePos := pos + fileSize
				spaces[newSpaceSize] = insertIntoSlice(spaces[newSpaceSize], newSpacePos)
			}

		}

		// Calculate the checksum of the blocks taken up by the file
		part2Checksum += checksumFromPos(filePos[fileId], fileSizes[fileId], fileId)
	}
	return part2Checksum
}

func insertIntoSlice(slc []int, new int) []int {
	newSlc := make([]int, len(slc)+1)
	idx, _ := utils.Find(slc, func(x int) bool {
		return x > new
	})
	if idx == -1 {
		idx = len(slc)
	}
	copy(newSlc[:idx], slc[:idx])
	newSlc[idx] = new
	copy(newSlc[idx+1:], slc[idx:])
	return newSlc
}

func checksum(start int, end int, id int) int {
	return id * ((end - start + 1) * (end + start)) / 2
}

func checksumFromPos(pos int, size int, id int) int {
	end := pos + size - 1
	return checksum(pos, end, id)
}
