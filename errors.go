package ras

import "errors"

var ErrSegFault = errors.New("segfualt")
var ErrSliceOutOfBounds = errors.New("array out of bounds")
