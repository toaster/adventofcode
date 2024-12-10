package amphipod

import (
	"fmt"
	"math"
	"slices"
)

// Disk describes a disk of an amphipod (https://adventofcode.com/2024/day/9).
type Disk struct {
	blocks    []uint
	fileCount uint
	files     map[uint]file
	freeSpace []file
	size      uint
}

// InitializeDisk reads the given disk map and initializes a new Disk from it.
func InitializeDisk(diskMap string) Disk {
	var blocks []uint
	var fileID uint
	files := map[uint]file{}
	var freeSpace []file
	var pos uint
	for i, c := range diskMap {
		length := uint(c - '0')
		if i%2 == 0 {
			files[fileID] = file{pos: pos, length: length}
			for j := uint(0); j < length; j++ {
				blocks = append(blocks, fileID)
			}
			fileID++
		} else if length > 0 {
			freeSpace = append(freeSpace, file{pos: pos, length: length})
			for j := uint(0); j < length; j++ {
				blocks = append(blocks, empty)
			}
		}
		pos += uint(length)
	}
	return Disk{
		blocks:    blocks,
		files:     files,
		freeSpace: freeSpace,
		size:      pos,
	}
}

// Checksum computes the current checksum of the disk.
func (d Disk) Checksum() uint {
	var checksum uint
	for i := uint(0); i < d.size; i++ {
		if d.blocks[i] == empty {
			continue
		}

		checksum += d.blocks[i] * i
	}
	return checksum
}

// Compact moves occupied blocks from the end to foremost free space until there is no free space
// except of the end of the disk.
func (d Disk) Compact() {
	fsIndex := 0
	var fsBlockIndex uint
	var last uint
	for i := d.size - 1; i > 0; i-- {
		if d.blocks[i] == empty {
			continue
		}

		fs := d.freeSpace[fsIndex]
		target := fs.pos + fsBlockIndex
		if target >= i {
			last = i
			break
		}

		d.blocks[target] = d.blocks[i]
		d.blocks[i] = empty
		fsBlockIndex++
		if fsBlockIndex >= fs.length {
			fsBlockIndex = 0
			fsIndex++
			if fsIndex >= len(d.freeSpace) {
				last = i
				break
			}
		}
	}
	d.freeSpace = []file{{
		length: d.size - (last + 1),
		pos:    last + 1,
	}}
}

// CompactFiles moves all files to free space before them if they fit starting with the last file.
func (d Disk) CompactFiles() {
	for id := uint(len(d.files) - 1); id > 0; id-- {
		f := d.files[id]
		for i, fs := range d.freeSpace {
			if fs.pos > f.pos {
				break
			}
			if fs.length < f.length {
				continue
			}

			for j := uint(0); j < f.length; j++ {
				d.blocks[fs.pos+j] = id
				d.blocks[f.pos+j] = empty
			}
			d.files[id] = file{pos: fs.pos, length: f.length}
			if f.length == fs.length {
				d.freeSpace = append(d.freeSpace[:i], d.freeSpace[i+1:]...)
			} else {
				d.freeSpace[i].length -= f.length
				d.freeSpace[i].pos += f.length
			}
			for j := i; j < len(d.freeSpace); j++ {
				fsAfter := d.freeSpace[j]
				if fsAfter.pos < f.pos {
					continue
				}

				var fsBefore file
				if j > 0 {
					fsBefore = d.freeSpace[j-1]
				}
				if j == 0 || fsBefore.pos+fsBefore.length < f.pos {
					// no free space before old file position
					if fsAfter.pos > f.pos+f.length {
						// no free space after old file position
						d.freeSpace = slices.Insert(d.freeSpace, j, f)
					} else {
						d.freeSpace[j].pos = f.pos
						d.freeSpace[j].length += f.length
					}
				} else {
					if fsAfter.pos > f.pos+f.length {
						// no free space after old file position
						d.freeSpace[j-1].length += f.length
					} else {
						d.freeSpace[j-1].length += f.length + fsAfter.length
						d.freeSpace = append(d.freeSpace[:j], d.freeSpace[j+1:]...)
					}
				}
				break
			}
			break
		}
	}
}

// Print prints the current disk content.
func (d Disk) Print() {
	for _, block := range d.blocks {
		if block == empty {
			fmt.Print(".")
		} else {
			fmt.Printf("%d", block)
		}
		fmt.Print(" ")
	}
	fmt.Println()
}

const empty = math.MaxUint

type file struct {
	length uint
	pos    uint
}
