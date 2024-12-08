package models

import (
	"time"

	"github.com/beego/beego/orm"
)

// User represents a user in the system
type User struct {
	Id        int       `orm:"column(id);auto"`
	Username  string    `orm:"column(username);size(255)"`
	Password  string    `orm:"column(password);size(255)"`
	Email     string    `orm:"column(email);size(255)"`
	Role      string    `orm:"column(role);size(255)"`
	Mobile    string    `orm:"column(mobile);size(255)"`
	UpdatedAt time.Time `orm:"column(updatedAt);type(datetime);null;auto_now"`     // auto_now is used to set the time when the record is updated
	CreatedAt time.Time `orm:"column(createdAt);type(datetime);null;auto_now_add"` // auto_now_add is used to set the time when the record is created
}

// TableName returns the name of the table in the database
func (u *User) TableName() string {
	return "ev_user"
}

// init registers the model with the ORM
func init() {
	orm.RegisterModel(new(User))
}

// CreateUser creates a new user in the database
func (u *User) CreateUser() (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(u)
	return id, err
}
