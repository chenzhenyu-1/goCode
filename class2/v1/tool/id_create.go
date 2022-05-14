package tool

import (
	"class2/repository"
	"fmt"
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
			id.Value = repository.NewPostDaoInstance().GetPostMaxId()
			fmt.Println(id.Value)
		})
	return id
}

func (p *ID) GetID() int64 {
	defer func() {
		p.Lock()
		p.Value++
		p.Unlock()
	}()
	return p.Value
}
