package util

import (
	"bytes"
	"strconv"
)

func PartitionToBytes(partition []int8) []byte {
	var buffer bytes.Buffer
	var end_index = len(partition) - 1
	for index, part := range partition {
		buffer.WriteString(strconv.Itoa(int(part)))
		if index != end_index {
			buffer.WriteString(",")
		}
	}
	return buffer.Bytes()
}
