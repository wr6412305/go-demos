package repositories

import (
	"go-demos/iris-demo/movie/datamodels"
	"sync"
)

// Query 代表一种“访客”和它的查询动作
type Query func(datamodels.Movie) bool

// MovieRepository 会处理一些关于movie实例的基本的操作
// 这是一个以测试为目的的接口，即是一个内存中的movie库
// 或是一个连接到数据库的实例
type MovieRepository interface {
	Exec(query Query, action Query, limit int, mode int) (ok bool)
	Select(query Query) (movie datamodels.Movie, found bool)
	SelectMany(query Query, limit int) (results []datamodels.Movie)
	InsertOrUpdate(movie datamodels.Movie) (updatedMovie datamodels.Movie, err error)
	Delete(query Query, limit int) (deleted bool)
}

// NewMovieRepository 返回一个新的基于内存的movie库
// 库的类型在我们的例子中是唯一的
func NewMovieRepository(source map[int64]datamodels.Movie) MovieRepository {
	return &movieMemoryRepository{source: source}
}

// movieMemoryRepository就是一个"MovieRepository"
// 它负责存储于内存中的实例数据(map)
type movieMemoryRepository struct {
	source map[int64]datamodels.Movie
	mu     sync.RWMutex
}

const (
	// ReadOnlyMode will RLock(read) the data .
	ReadOnlyMode = iota
	// ReadWriteMode will Lock(read/write) the data.
	ReadWriteMode
)
