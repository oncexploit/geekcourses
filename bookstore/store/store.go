//store/store.go	创建抽象数据类型Book和针对Book存取的接口类型Store
package store

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("exist")
)

type Book struct { //json两侧的符号是数字1左边的`(波浪键),而不是单引号，由于此原因我多次报错
	Id      string   `json:"id"`      // 图书ISBN ID
	Name    string   `json:"name"`    // 图书名称
	Authors []string `json:"authors"` // 图书作者
	Press   string   `json:"press"`   // 出版社

}

type Store interface {
	Create(*Book) error       //创建一个新图书条目
	Update(*Book) error       //更新某图书条目
	Get(string) (Book, error) //获取某图书条目的信息
	GetAll() ([]Book, error)  //获取所有图书条目的信息
	Delete(string) error      //删除某图书条目的信息
}
