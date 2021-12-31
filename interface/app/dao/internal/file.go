// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// FileDao is the manager for logic model data accessing and custom defined data operations functions management.
type FileDao struct {
	gmvc.M             // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	C      fileColumns // C is the short type for Columns, which contains all the column names of Table for convenient usage.
	DB     gdb.DB      // DB is the raw underlying database management object.
	Table  string      // Table is the underlying table name of the DAO.
}

// FileColumns defines and stores column names for table s_file.
type fileColumns struct {
	Id        string //
	Name      string //
	Group     string //
	Status    string //
	CreatedAt string //
	UpdatedAt string //
}

// NewFileDao creates and returns a new DAO object for table data access.
func NewFileDao() *FileDao {
	columns := fileColumns{
		Id:        "id",
		Name:      "name",
		Group:     "group",
		Status:    "status",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	}
	return &FileDao{
		C:     columns,
		M:     g.DB("default").Model("s_file").Safe(),
		DB:    g.DB("default"),
		Table: "s_file",
	}
}
