// Package lock provides an interesting lock mechanism I thought about
// however I have not found a single use case for it yet in software.
package lock

import "sync/atomic"

import "runtime"

import "strconv"
import "github.com/mariusvanderwijden/go-snippets/search"

type Lock struct {
	lock bool
	// mu locks routines and enterValues if we need to append
	mu sync.Mutex
	routines []Routines
}

type Routines struct {
	rID int
	value bool
}

func (l *Lock) Enter() {
	id, err := getGoroutineID()
	if err != nil {
		panic(err)
	}
	l.mu.RLock()
	index := search.Search(len(l.routines), func (i int) bool {
		return l.routines[i] == id
	})
	l.mu.Unlock()
	if index == -1 {
		l.mu.Lock()
		l.routines = append(l.routines, Routines{id, false})
		index = len(l.routines) - 1
		l.mu.Unlock()
	}
	l.routines[index].value = atomic.LoadBool(&l.lock)
}

func (l *Lock) Exit() {
	id, err := getGoroutineID()
	if err != nil {
		panic(err)
	}
	l.mu.RLock()
	index := search.Search(len(l.routines), func (i int) bool {
		return l.routines[i] == id
	})
	l.mu.Unlock()
	atomic.StoreBool(&l.lock, l.routines[index].value)
}

func (l *Lock) ReadLock() bool {
	return atomic.LoadBool(&l.lock)
}


func getGoroutineID() (int, error) {
	var buffer [64] byte
	size := runtime.Stack(buffer[:], false)
	id := strings.Fields(strings.TrimPrefix(string(buffer[:size]), "goroutine "))[0]
	return strconv.Atoi(id)
}
