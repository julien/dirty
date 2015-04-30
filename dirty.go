package dirty

import (
	"bufio"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sync"
)

// Dirty is a JSON based storage
type Dirty struct {
	Path string

	lock sync.RWMutex
	docs map[string]interface{}
}

// NewDirty returns a new Dirty instance
// and loads data given a file path
func NewDirty(Path string) *Dirty {
	d := &Dirty{
		Path: Path,
		docs: make(map[string]interface{}),
	}
	d.Read()
	return d
}

// Set a key/value
func (d *Dirty) Set(key string, val interface{}) {
	if val == nil {
		delete(d.docs, key)
	} else {
		if _, ok := d.docs[key]; ok == false {
			d.lock.Lock()
			d.docs[key] = val
			d.lock.Unlock()
		}
	}
}

// Flush saves all the key/values to a .json file
func (d *Dirty) Flush() error {

	f, err := os.Create(d.Path)
	defer f.Close()
	if err != nil {
		return err
	}

	d.lock.Lock()
	out, err := json.MarshalIndent(d.docs, "", " ")
	d.lock.Unlock()

	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(out))
	w.Flush()
	return nil
}

// Read loads all the key/values for a json file
func (d *Dirty) Read() error {
	if d.Path == "" {
		return errors.New("No path")
	}

	data, err := ioutil.ReadFile(d.Path)
	if err != nil {
		return err
	}

	d.lock.Lock()
	err = json.Unmarshal(data, &d.docs)
	d.lock.Unlock()
	if err != nil {
		return err
	}

	return nil
}

// Get a key given it's name
func (d *Dirty) Get(key string) interface{} {
	return d.docs[key]
}

// Keys returns all keys
func (d *Dirty) Keys() []string {
	var keys []string
	for k := range d.docs {
		keys = append(keys, k)
	}
	return keys
}

// All return all key/values
func (d *Dirty) All() map[string]interface{} {
	return d.docs
}
