// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package internal

import (
	"database/sql"
	"focus/app/model"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
	"time"
)

// UserDetailDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type UserDetailDao struct {
	gmvc.M
	Table   string
	Columns userDetailColumns
}

// UserDetailColumns defines and stores column names for table gf_user_detail.
type userDetailColumns struct {
	UserId     string //                                                
    TrueName   string // 真实姓名                                       
    CardId     string // 身份证号                                       
    City       string // 城市                                           
    Birth      string // 出生日期\n(字符串，例如：1986-10-07 00:00:00)  
    Phone      string // 手机号码                                       
    Qq         string // QQ                                             
    Email      string // 邮件                                           
    From       string // 用户来源                                       
    Brief      string // 用户说明                                       
    CreatedAt  string //                                                
    UpdatedAt  string //
}

var (
	// UserDetail is globally public accessible object for table gf_user_detail operations.
	UserDetail = &UserDetailDao{
		M:     g.DB("default").Table("gf_user_detail").Safe(),
		Table: "gf_user_detail",
		Columns: userDetailColumns{
			UserId:    "user_id",     
            TrueName:  "true_name",   
            CardId:    "card_id",     
            City:      "city",        
            Birth:     "birth",       
            Phone:     "phone",       
            Qq:        "qq",          
            Email:     "email",       
            From:      "from",        
            Brief:     "brief",       
            CreatedAt: "created_at",  
            UpdatedAt: "updated_at",
		},
	}
)

// As sets an alias name for current table.
func (d *UserDetailDao) As(as string) *UserDetailDao {
	return &UserDetailDao{M:d.M.As(as)}
}

// TX sets the transaction for current operation.
func (d *UserDetailDao) TX(tx *gdb.TX) *UserDetailDao {
	return &UserDetailDao{M:d.M.TX(tx)}
}

// Master marks the following operation on master node.
func (d *UserDetailDao) Master() *UserDetailDao {
	return &UserDetailDao{M:d.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (d *UserDetailDao) Slave() *UserDetailDao {
	return &UserDetailDao{M:d.M.Slave()}
}

// Args sets custom arguments for model operation.
func (d *UserDetailDao) Args(args ...interface{}) *UserDetailDao {
	return &UserDetailDao{M:d.M.Args(args ...)}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *UserDetailDao) LeftJoin(table ...string) *UserDetailDao {
	return &UserDetailDao{M:d.M.LeftJoin(table...)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *UserDetailDao) RightJoin(table ...string) *UserDetailDao {
	return &UserDetailDao{M:d.M.RightJoin(table...)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *UserDetailDao) InnerJoin(table ...string) *UserDetailDao {
	return &UserDetailDao{M:d.M.InnerJoin(table...)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *UserDetailDao) Fields(fieldNamesOrMapStruct ...interface{}) *UserDetailDao {
	return &UserDetailDao{M:d.M.Fields(fieldNamesOrMapStruct...)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *UserDetailDao) FieldsEx(fieldNamesOrMapStruct ...interface{}) *UserDetailDao {
	return &UserDetailDao{M:d.M.FieldsEx(fieldNamesOrMapStruct...)}
}

// Option sets the extra operation option for the model.
func (d *UserDetailDao) Option(option int) *UserDetailDao {
	return &UserDetailDao{M:d.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (d *UserDetailDao) OmitEmpty() *UserDetailDao {
	return &UserDetailDao{M:d.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (d *UserDetailDao) Filter() *UserDetailDao {
	return &UserDetailDao{M:d.M.Filter()}
}

// Where sets the condition statement for the model. The parameter <where> can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
// Eg:
// Where("uid=10000")
// Where("uid", 10000)
// Where("money>? AND name like ?", 99999, "vip_%")
// Where("uid", 1).Where("name", "john")
// Where("status IN (?)", g.Slice{1,2,3})
// Where("age IN(?,?)", 18, 50)
// Where(User{ Id : 1, UserName : "john"})
func (d *UserDetailDao) Where(where interface{}, args ...interface{}) *UserDetailDao {
	return &UserDetailDao{M:d.M.Where(where, args...)}
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (d *UserDetailDao) WherePri(where interface{}, args ...interface{}) *UserDetailDao {
	return &UserDetailDao{M:d.M.WherePri(where, args...)}
}

// And adds "AND" condition to the where statement.
func (d *UserDetailDao) And(where interface{}, args ...interface{}) *UserDetailDao {
	return &UserDetailDao{M:d.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (d *UserDetailDao) Or(where interface{}, args ...interface{}) *UserDetailDao {
	return &UserDetailDao{M:d.M.Or(where, args...)}
}

// Group sets the "GROUP BY" statement for the model.
func (d *UserDetailDao) Group(groupBy string) *UserDetailDao {
	return &UserDetailDao{M:d.M.Group(groupBy)}
}

// Order sets the "ORDER BY" statement for the model.
func (d *UserDetailDao) Order(orderBy ...string) *UserDetailDao {
	return &UserDetailDao{M:d.M.Order(orderBy...)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (d *UserDetailDao) Limit(limit ...int) *UserDetailDao {
	return &UserDetailDao{M:d.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (d *UserDetailDao) Offset(offset int) *UserDetailDao {
	return &UserDetailDao{M:d.M.Offset(offset)}
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (d *UserDetailDao) Page(page, limit int) *UserDetailDao {
	return &UserDetailDao{M:d.M.Page(page, limit)}
}

// Batch sets the batch operation number for the model.
func (d *UserDetailDao) Batch(batch int) *UserDetailDao {
	return &UserDetailDao{M:d.M.Batch(batch)}
}

// Cache sets the cache feature for the model. It caches the result of the sql, which means
// if there's another same sql request, it just reads and returns the result from cache, it
// but not committed and executed into the database.
//
// If the parameter <duration> < 0, which means it clear the cache with given <name>.
// If the parameter <duration> = 0, which means it never expires.
// If the parameter <duration> > 0, which means it expires after <duration>.
//
// The optional parameter <name> is used to bind a name to the cache, which means you can later
// control the cache like changing the <duration> or clearing the cache with specified <name>.
//
// Note that, the cache feature is disabled if the model is operating on a transaction.
func (d *UserDetailDao) Cache(duration time.Duration, name ...string) *UserDetailDao {
	return &UserDetailDao{M:d.M.Cache(duration, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (d *UserDetailDao) Data(data ...interface{}) *UserDetailDao {
	return &UserDetailDao{M:d.M.Data(data...)}
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*model.UserDetail.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *UserDetailDao) All(where ...interface{}) ([]*model.UserDetail, error) {
	all, err := d.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.UserDetail
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *model.UserDetail.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *UserDetailDao) One(where ...interface{}) (*model.UserDetail, error) {
	one, err := d.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.UserDetail
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (d *UserDetailDao) FindOne(where ...interface{}) (*model.UserDetail, error) {
	one, err := d.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.UserDetail
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (d *UserDetailDao) FindAll(where ...interface{}) ([]*model.UserDetail, error) {
	all, err := d.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.UserDetail
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// Chunk iterates the table with given size and callback function.
func (d *UserDetailDao) Chunk(limit int, callback func(entities []*model.UserDetail, err error) bool) {
	d.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*model.UserDetail
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (d *UserDetailDao) LockUpdate() *UserDetailDao {
	return &UserDetailDao{M:d.M.LockUpdate()}
}

// LockShared sets the lock in share mode for current operation.
func (d *UserDetailDao) LockShared() *UserDetailDao {
	return &UserDetailDao{M:d.M.LockShared()}
}

// Unscoped enables/disables the soft deleting feature.
func (d *UserDetailDao) Unscoped() *UserDetailDao {
	return &UserDetailDao{M:d.M.Unscoped()}
}