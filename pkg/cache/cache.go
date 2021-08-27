package cache

import (
	"html/template"
)

type Cache struct {
	Templates map[string]*template.Template
}

func NewCache(templates map[string]*template.Template) *Cache {
	return &Cache{Templates: templates}
}
