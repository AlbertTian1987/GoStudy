package cache
import "testing"

type CacheValue struct {
    size int
}

func (cv *CacheValue) Size() int{
    return cv.size
}

func TestInitialState(t *testing.T){
    cache := NewLRUCache(5)
    l,sz,c,_ := cache.Stats()
    if l != 0 {
        t.Errorf("length = %v, want 0", l)
    }
    if sz != 0 {
        t.Errorf("size = %v, want 0", sz)
    }
    if c != 5 {
        t.Errorf("capacity = %v, want 5", c)
    }
}