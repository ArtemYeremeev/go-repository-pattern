package interfaces

// IDbHandler descrybes database interface
type IDbHandler interface {
	Execute(st string)
	Query(st string) (IRow, error)
}

// IRow descrybes row interface
type IRow interface {
	Scan(dst ...interface{}) error
	Next() bool
}
