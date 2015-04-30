package dirty

import "testing"

func TestNewDirty(t *testing.T) {
	d := NewDirty("example.json")
	if d.docs == nil {
		t.Errorf("got %v", d.docs)
	}
}

func TestSetVal(t *testing.T) {
	d := NewDirty("example.json")
	if d.docs == nil {
		t.Errorf("got %v", d.docs)
	}
	d.Set("C", "three")
	if d.Get("C") != "three" {
		t.Errorf("got %v\n", d.Get("C"))
	}
	d.Flush()
}

func TestSetOtherVal(t *testing.T) {
	d := NewDirty("example.json")
	if d.docs == nil {
		t.Errorf("got %v", d.docs)
	}

	d.Set("D", "four")
	d.Set("E", "five")
	if d.Get("D") != "four" {
		t.Errorf("got %v\n", d.Get("C"))
	}
}

func TestKeys(t *testing.T) {
	d := NewDirty("example.json")
	k := d.Keys()
	if len(k) != 3 {
		t.Errorf("got %v\n", len(k))
	}
}

func TestAll(t *testing.T) {
	d := NewDirty("example.json")
	a := d.All()
	if len(a) != 3 {
		t.Errorf("got %v\n", len(a))
	}
}

func TestSetNil(t *testing.T) {
	d := NewDirty("example.json")
	d.Set("C", nil)

	a := d.All()
	if len(a) != 2 {
		t.Errorf("got %v\n", len(a))
	}
}

func TestNoPath(t *testing.T) {
	d := NewDirty("")
	if d.Get("A") != nil {
		t.Errorf("got %v\n", d.Get("A"))
	}
}

func TestNonExistingPath(t *testing.T) {
	d := NewDirty("non-existing.json")
	if d.Get("A") != nil {
		t.Errorf("got %v\n", d.Get("A"))
	}
}
