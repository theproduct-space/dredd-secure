package keeper

type Exec struct {
	ID       string
	Function func(...interface{}) interface{}
	Args     []interface{}
	DelayS   int64
}
