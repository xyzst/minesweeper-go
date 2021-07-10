package board

type RowError struct {
	Err string
}

func (e RowError) Error() string {
	return e.Err
}

type ColumnError struct {
	Err string
}

func (e ColumnError) Error() string {
	return e.Err
}

type LogicalError struct {
	Err string
}

func (e LogicalError) Error() string {
	return e.Err
}