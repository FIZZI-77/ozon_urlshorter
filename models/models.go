package models

import "errors"

type Link struct {
	ID    int64  `json:"id,omitempty"`
	URL   string `json:"url"`
	Short string `json:"short"`
}

var ErrNotFound = errors.New("not found")

type CreateRequest struct {
	URL string `json:"url" binding:"required"`
}

type ResolveRequest struct {
	Short string `json:"short"`
}

type CreateResponse struct {
	Short string `json:"short"`
}

type ResolveResponse struct {
	URL string `json:"url"`
}
