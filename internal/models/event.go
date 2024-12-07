package models

import "time"

type Event struct {
	Id          int64     `orm:"column(id);auto"`
	Name        string    `orm:"column(name)"`
	Description string    `orm:"column(description)"`
	Date        string    `orm:"column(date)"`
	Location    string    `orm:"column(location)"`
	CreatedOn   time.Time `orm:"column(createdOn);type(datetime);null;auto_now_add"`
	UpdatedOn   time.Time `orm:"column(updatedOn);type(datetime);null;auto_now"`
}
