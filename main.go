package dbsimpl

import (
	"fmt"
	"strings"
	"database/sql"
)

/*
The aim of this package is to make writing queries against the
database as easy as possible...


There is very little hand-holding in this package, and it's aim is to
provide an improvement over basic string-concatenation without having
to resort to a full ORM.

This project has a few simple goals

1. Select, Delete, Insert and Update are the only supported operations
2. There will be no explicit handling of table joins, or foreign keys
3. It should be possible to construct a query and modify/extend it if
'dbsimpl' can't do the job.

Things it doesn't/won't do
1. Make the actual query, this you can do on your own

Roadmap
1. pull returned data into a structure (although this is on the roadmap)
2. Allow for querying for a (single/the first) element
*/

type segment struct {
	keyword string
	value string
}

func (s *segment) String() string {
	str := fmt.Sprintf("%s %s", s.keyword, s.value)
	return strings.TrimSpace(str)
}

// Builder struct for building a Query
type Builder struct {
	segments []segment // a list of segments built a query
	args  []interface{} // a list of arguments to be passed into the builder
	iface interface{}
}

// NewRaw ...
func NewRaw() Builder {
	return Builder{
		segments: make([]segment, 0),
		args: make([]interface{}, 0),
	}
}

// New ...
func New(str interface{}) Builder {
	fmt.Println("interface needs to be reflected for great good!")
	return NewRaw()
}

func (b *Builder) rawS(keyword string, value string, args []interface{}) {
	b.segments = append(b.segments, segment{
		keyword: keyword,
		value: value,
		
	})
	b.args = append(b.args, args)
}

// RawS ...
func (b *Builder) RawS(value string, args ...interface{}) *Builder {
	b.rawS("", value, args)
	return b
}

// Select builder
func (b *Builder) Select(value string, args ...interface{}) *Builder {
	b.rawS("SELECT", value, args)
	return b
}

// From builder
func (b *Builder) From(value string, args ...interface{}) *Builder {
	b.rawS("FROM", value, args)
	return b
}

// Where builder
func (b *Builder) Where(value string, args ...interface{}) *Builder {
	b.rawS("WHERE", value, args)
	return b
}

// And builder
func (b *Builder) And(value string, args ...interface{}) *Builder {
	b.rawS("AND", value, args)
	return b
}

// Build builds and returns the raw query
func (b * Builder) Build() string {
	var queryList = []string{}
	for _, segment := range b.segments {
		queryList = append(queryList, segment.String())
	}

	return strings.Join(queryList, " ")
}

// ExecuteRaw the query on the passed in database
func (b *Builder) ExecuteRaw(db *sql.DB) (*sql.Rows, error) {
	query := b.Build()
	return db.Query(query, b.args)
}
