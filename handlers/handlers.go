package handlers

import (
	"goApi/db"
)

type UserGet struct {
	Parce []db.Data `json:"parce"`
 }

 type BaseHandler struct {
	db *db.DB
 }

func NewBaseHandler(db *db.DB) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}