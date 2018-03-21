package wafdata

import (
	"../omdb"
	"fmt"
	"math"
	"sync"
)

type IdAllocator struct {
	Name      string `xorm:"pk varchar(255) notnull"`
	Tag       []byte
	Candidate uint32     `xorm:"-"`
	Max       uint32     `xorm:"notnull default(0)"`
	Current   uint32     `xorm:"notnull default(0)"`
	Invalid   uint32     `xorm:"notnull default(0)"`
	Mutex     sync.Mutex `xorm:"-"`
}

const (
	// max index is math.MaxUint32 - 1
	INVALID_INDEX uint32 = math.MaxUint32 - 1
)

func (p *IdAllocator) Alloc() uint32 {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()
	var i uint32
	for i = p.Candidate; i < p.Max; i++ {
		if p.Tag[i] == 0 {
			p.Tag[i] = 1
			DBHandle := omdb.GetDB()
			if errcode, err := DBHandle.Update(p); err != nil {
				fmt.Println("IdAllocator.Alloc() returns error:", string(errcode), err.Error())
				return INVALID_INDEX
			}
			p.Candidate = i + 1
			return i
		}
	}
	return INVALID_INDEX
}

func (p *IdAllocator) Free(idx uint32) bool {
	if idx >= p.Max {
		fmt.Printf("IdAllocator.Free() returns error. idx=%v, Max=%v", idx, p.Max)
		return false
	}

	p.Mutex.Lock()
	defer p.Mutex.Unlock()

	p.Tag[idx] = 0
	DBHandle := omdb.GetDB()
	if errcode, err := DBHandle.Update(p); err != nil {
		fmt.Println("IdAllocator.Free() returns error:", string(errcode), err.Error())
		return false
	}
	p.Candidate = idx
	return true
}

func DestroyIdAllocator(name string) bool {
	DBHandle := omdb.GetDB()
	if errcode, err := DBHandle.Delete(&IdAllocator{Name: name}); err != nil {
		fmt.Println("IdAllocator.Alloc() returns error:", string(errcode), err.Error())
		return false
	}
	return true
}

func CreateIdAllocator(name string, max uint32) *IdAllocator {
	if name == "" {
		fmt.Println("CreateIdAllocator() error: cannot null name")
		return nil
	}

	if max >= math.MaxUint32 {
		fmt.Println("CreateIdAllocator() error: cannot accept max value >=", string(math.MaxUint32))
		return nil
	}

	DBHandle := omdb.GetDB()
	NewAllocator := new(IdAllocator)
	NewAllocator.Mutex = sync.Mutex{}
	NewAllocator.Name = name
	NewAllocator.Max = max
	NewAllocator.Candidate = 0
	NewAllocator.Invalid = INVALID_INDEX
	NewAllocator.Current = NewAllocator.Invalid
	NewAllocator.Tag = make([]byte, 0, NewAllocator.Max)
	if errcode, err := DBHandle.Insert(NewAllocator); err != nil {
		fmt.Println("DBHandle.Insert() returns error:", string(errcode), err.Error())
		return nil
	}
	return NewAllocator
}

func GetIdAllocator(name string) *IdAllocator {
	var Allocator = IdAllocator{Name: name}
	DBHandle := omdb.GetDB()
	if has, err := DBHandle.Get(&Allocator); has == false {
		fmt.Println("GetIdAllocator returns error:", err.Error())
		return nil
	}
	Allocator.Mutex = sync.Mutex{}
	return &Allocator
}