package datafile3

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
    "sync/atomic"
)

type Data []byte

type DataFile interface {
	Read() (rsn int64, d Data, err error)
	Write(d Data) (wsn int64, err error)
	Rsn() int64
	Wsn() int64
	DataLen() uint32
}

type myDataFile struct {
	file     *os.File
	fileLock sync.RWMutex
    rcond    *sync.Cond
	roffset  int64
	woffset  int64
	dataLen  uint32
}

func NewMyDataFile(filePath string, dataLen uint32) (DataFile, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	if dataLen <= 0 {
		return nil, errors.New(fmt.Sprintf("Invalid dataLen %d", dataLen))
	}
    var fileLock sync.RWMutex
    rcond := sync.NewCond(fileLock.RLocker())
	return &myDataFile{
		file:    file,
		dataLen: dataLen,
        fileLock:fileLock,
        rcond:rcond,
	},nil
}

func (self *myDataFile) Read() (rsn int64, d Data, err error) {
    var offset int64
    for {
        offset = atomic.LoadInt64(&self.roffset)
        if atomic.CompareAndSwapInt64(&self.roffset,offset,(offset+int64(self.dataLen))) {
            break
        }
    }

	rsn = offset / int64(self.dataLen)
	data := make([]byte, self.dataLen)

    self.fileLock.RLock()
    defer self.fileLock.RUnlock()
	for {
		_, err = self.file.ReadAt(data, offset)
		if err != nil {
			if err == io.EOF {
                self.rcond.Wait()
				continue
			}
			return
		}
		d = data
		return
	}
}

func (self *myDataFile) Write(d Data) (wsn int64, err error) {
	var offset int64
    for{
        offset = atomic.LoadInt64(&self.woffset)
        if atomic.CompareAndSwapInt64(&self.woffset, offset, (offset+int64(self.dataLen))) {
            break
        }
    }

	wsn = offset / int64(self.dataLen)
	var data []byte
	if len(d) > int(self.dataLen) {
		data = d[0:self.dataLen]
	} else {
		data = d
	}
	self.fileLock.Lock()
	defer self.fileLock.Unlock()
	_, err = self.file.Write(data)
    self.rcond.Signal()
	return

}

func (self *myDataFile) Rsn() int64 {
    offset:=atomic.LoadInt64(&self.roffset)
	return offset / int64(self.dataLen)
}
func (self *myDataFile) Wsn() int64 {
    offset:=atomic.LoadInt64(&self.woffset)
	return offset / int64(self.dataLen)
}
func (self *myDataFile) DataLen() uint32 {
	return self.dataLen
}
