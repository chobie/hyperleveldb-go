package main

/*
#cgo LDFLAGS: -lhyperleveldb
#include <string.h>
#include <hyperleveldb/c.h>

static void CmpDestroy(void* arg) { }

static int CmpCompare(void* arg, const char* a, size_t alen,
                      const char* b, size_t blen) {
  int n = (alen < blen) ? alen : blen;
  int r = memcmp(a, b, n);
  if (r == 0) {
    if (alen < blen) r = -1;
    else if (alen > blen) r = +1;
  }
  return r;
}

static const char* CmpName(void* arg) {
  return "foo";
}

static hyperleveldb_comparator_t* CmpFooNew() {
  return hyperleveldb_comparator_create(NULL, CmpDestroy, CmpCompare, CmpName);
}

*/
import "C"

type Comparator struct {
	Comparator *C.hyperleveldb_comparator_t
}

func NewFooComparator() *Comparator {
	return &Comparator{C.CmpFooNew()}
}

func (cmp *Comparator) Close() {
	C.hyperleveldb_comparator_destroy(cmp.Comparator)
}

func main() {
	NewFooComparator().Close()
}
