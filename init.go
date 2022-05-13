package region

import "reflect"

const (
	INDEX_BLOCK_LENGTH  = 12
	TOTAL_HEADER_LENGTH = 8192
)

var (
	regionTypeOf = reflect.TypeOf((*Region)(nil)).String()
)
