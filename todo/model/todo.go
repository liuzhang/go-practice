package model

type Todo struct {
	Id      int    `json:"id" yaml:"id"`
	Content string `json:"content" yaml:"content"`
	Done    bool   `json:"done" yaml:"done"`
}
