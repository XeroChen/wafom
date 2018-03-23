package omdata

import (
	"../omdb"
	"fmt"
	"github.com/go-xorm/core"
	"math"
	"sync"
)

type IdAllocator struct {
	Name      string `xorm:"pk varchar(255) notnull"`
	Tag       []byte
	Max       uint32     `xorm:"notnull default(0)"`
	Candidate uint32     `xorm:"-"`
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
			defer omdb.CloseDB(DBHandle)
			if errcode, err := DBHandle.ID(core.PK{p.Name}).Cols("Tag").Update(p); err != nil {
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
	defer omdb.CloseDB(DBHandle)
	if errcode, err := DBHandle.ID(core.PK{p.Name}).Cols("Tag").Update(p); err != nil {
		fmt.Println("IdAllocator.Free() returns error:", string(errcode), err.Error())
		return false
	}
	p.Candidate = idx
	return true
}

func DestroyIdAllocator(name string) bool {
	DBHandle := omdb.GetDB()
	defer omdb.CloseDB(DBHandle)

	if errcode, err := DBHandle.ID(core.PK{name}).Delete(new(IdAllocator)); err != nil {
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
	defer omdb.CloseDB(DBHandle)
	/* Table does not exist, create it first */
	if ext, _ := DBHandle.IsTableExist(new(IdAllocator)); ext != true {
		err := DBHandle.CreateTables(new(IdAllocator))
		if err != nil {
			fmt.Println("DBHandle.Insert() returns error:", err.Error())
			return nil
		}
	}
	NewAllocator := new(IdAllocator)
	NewAllocator.Mutex = sync.Mutex{}
	NewAllocator.Name = name
	NewAllocator.Max = max
	NewAllocator.Candidate = 0
	NewAllocator.Tag = make([]byte, NewAllocator.Max)
	if errcode, err := DBHandle.Insert(NewAllocator); err != nil {
		fmt.Println("DBHandle.Insert() returns error:", string(errcode), err.Error())
		return nil
	}
	return NewAllocator
}

func GetIdAllocator(name string) *IdAllocator {
	var Allocator IdAllocator
	DBHandle := omdb.GetDB()
	defer omdb.CloseDB(DBHandle)

	if has, err := DBHandle.ID(core.PK{name}).Get(&Allocator); has == false {
		fmt.Println("GetIdAllocator returns error:", err.Error())
		return nil
	}
	Allocator.Mutex = sync.Mutex{}
	return &Allocator
}
