package omdata

import (
	"../omdb"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	testSetup()
	m.Run()
	testTearDown()
}

func testSetup() {
	db := omdb.GetDB()
	if db == nil {
		fmt.Println("testSetup failed!")
		os.Exit(1)
	}
	db.ShowSQL(true)
	db.DropTables(new(IdAllocator))
	omdb.CloseDB(db)
}

func testTearDown() {
	db := omdb.GetDB()
	if db == nil {
		fmt.Println("testTearDown failed!")
		os.Exit(1)
	}
	db.DropTables(new(IdAllocator))
	omdb.CloseDB(db)
}

func TestCreateIdAllocator(t *testing.T) {
	IdAllocator := CreateIdAllocator("TestCreateIdAllocator", 1)
	if IdAllocator == nil {
		t.Errorf("CreateIdAllocator failed.")
	}

	var Idx uint32
	Idx = IdAllocator.Alloc()
	if Idx != 0 {
		t.Errorf("IdAllocator.Alloc() Idx=%v.", Idx)
	}

	Idx = IdAllocator.Alloc()
	if Idx != INVALID_INDEX {
		t.Errorf("IdAllocator.Alloc() Idx=%v.", Idx)
	}

	succeed := IdAllocator.Free(Idx)
	if succeed != true {
		t.Errorf("IdAllocator.Free() failed.")
	}

	Idx = IdAllocator.Alloc()
	if Idx != 0 {
		t.Errorf("IdAllocator.Alloc() Idx=%v.", Idx)
	}
}

func TestAlloc(t *testing.T) {
	IDalct := CreateIdAllocator("TestAlloc", 2)
	if IDalct == nil {
		t.Errorf("CreateIdAllocator failed.")
		return
	}
	var Idx uint32
	Idx = IDalct.Alloc()
	if Idx != 0 {
		t.Errorf("IdAllocator.Alloc() Idx=%v.", Idx)
	}

	Idx = IDalct.Alloc()
	if Idx != 1 {
		t.Errorf("IdAllocator.Alloc() Idx=%v.", Idx)
	}

	Idx = IDalct.Alloc()
	if Idx != INVALID_INDEX {
		t.Errorf("IdAllocator.Alloc() Idx=%v.", Idx)
	}
}

func TestFree(t *testing.T) {
	IDalct := CreateIdAllocator("TestFree", 2)
	if IDalct == nil {
		t.Errorf("CreateIdAllocator failed.")
	}
	var Idx uint32
	Idx = IDalct.Alloc()
	if Idx != 0 {
		t.Errorf("IdAllocator.Alloc() Idx=%v.", Idx)
	}

	Idx = IDalct.Alloc()
	if Idx != 1 {
		t.Errorf("IdAllocator.Alloc() Idx=%v.", Idx)
	}

	succeed := IDalct.Free(Idx)
	if succeed != true {
		t.Errorf("IdAllocator.Free() failed.")
	}

	Idx = IDalct.Alloc()
	if Idx != 1 {
		t.Errorf("IdAllocator.Alloc() Idx=%v.", Idx)
	}
}

func TestGet(t *testing.T) {
	IDalct := CreateIdAllocator("TestGet", 2)
	if IDalct == nil {
		t.Errorf("CreateIdAllocator failed.")
	}
	var Idx uint32
	Idx = IDalct.Alloc()
	if Idx != 0 {
		t.Errorf("IdAllocator.Alloc() Idx=%v.", Idx)
	}

	AnotherIDalct := GetIdAllocator("TestGet")
	if AnotherIDalct == nil {
		t.Errorf("GetIdAllocator failed.")
	}

	Idx = AnotherIDalct.Alloc()
	if Idx != 1 {
		t.Errorf("IdAllocator.Alloc() Idx=%v.", Idx)
	}

	succeed := AnotherIDalct.Free(Idx)
	if succeed != true {
		t.Errorf("IdAllocator.Free() failed.")
	}

	Idx = IDalct.Alloc()
	if Idx != 1 {
		t.Errorf("IdAllocator.Alloc() Idx=%v.", Idx)
	}
}
