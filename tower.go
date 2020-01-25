package main

import (
	"fmt"
)

type Tower struct {
	disks []int
	index int
	callback func()
}

func NewTower(index int) *Tower {
	t := new(Tower)
	t.index = index
	return t
}

func (t *Tower) Index() int {
	return t.index
}

func (t *Tower) SetCallback(callback func()) {
	t.callback = callback
}

func (t *Tower) Add(disk int) {
	if len(t.disks) > 0 && t.disks[len(t.disks)-1] <= disk {
		panic(fmt.Sprintf("Tower[%d] cannot add disk (%d) on top of disk (%d)\n", t.index, t.disks[len(t.disks)-1], disk))
	} else {
		t.disks = append(t.disks, disk)
	}
}

func (t *Tower) MoveTopTo(destTower *Tower) {
	fmt.Printf("Moving Top disk %d from T[%d] to T[%d]\n", t.disks[len(t.disks)-1], t.Index(), destTower.Index())
	var disk int
	disk, t.disks = t.disks[len(t.disks)-1], t.disks[:len(t.disks)-1]
	destTower.Add(disk)

	t.callback()
}

func (t *Tower) MoveDisks(n int, destTower *Tower, buffTower *Tower) {
	if n > 0 {
		fmt.Printf("Move %d disks from T[%d] to T[%d] using T[%d]\n", n, t.Index(), destTower.Index(), buffTower.Index())
		t.MoveDisks(n-1, buffTower, destTower)
		t.MoveTopTo(destTower)
		buffTower.MoveDisks(n-1, destTower, t)
	}
}

func (t *Tower) ShowTower() {
	towerContent := fmt.Sprintf("Tower[%d]: ", t.Index())
	for _, d := range t.disks {
		towerContent += fmt.Sprintf("<%d>", d)
	}
	fmt.Println(towerContent)
}
