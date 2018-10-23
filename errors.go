package main

import (
	"errors"
	"fmt"
)

const (
	UnsupportedAnnotationFor = "unsupported annotation %s for %s"
)

func UnsupportedAnnotationForError(ann, tokenStr string) error {
	return errors.New(fmt.Sprintf(UnsupportedAnnotationFor, ann, tokenStr))
}
