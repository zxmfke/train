package hashmap

import "github.com/zxmfke/train/datastruct/list"

type HSMapValue interface {
	Get(key string) (interface{}, bool)

	Set(key string, v interface{})

	Delete(key string)
}

type HSMapListValue map[string]*list.List
