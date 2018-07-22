package ras

import "errors"

var ErrSegFault = errors.New("segfualt")
var ErrSliceOutOfBounds = errors.New("slice out of bounds")
