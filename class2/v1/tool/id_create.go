package tool

import (
	"bufio"
	"encoding/binary"
	"os"
	"sync"
)

type ID struct {
	Value int64
	sync.Mutex
}

var (
	id     *ID
	idOnce sync.Once
)

func NewIdInstance() *ID {
	idOnce.Do(
		func() {
			id = &ID{}
			filePath := ".\\data\\id"
			file, _ := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, os.ModeExclusive)
			read := bufio.NewReader(file)
			binary.Read(read, binary.BigEndian, id.Value)
		})
	return id
}

func (p *ID) SaveId() {
	p.Lock()
	filepath := ".\\data\\id"
	file, _ := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, os.ModeExclusive)
	defer file.Close()
	write := bufio.NewWriter(file)
	binary.Write(write, binary.BigEndian, p.Value)
	p.Unlock()
}

func (p *ID) GetID() int64 {
	defer func() {
		p.Lock()
		p.Value++
		p.Unlock()
	}()
	return p.Value
}
