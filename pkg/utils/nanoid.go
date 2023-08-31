package utils

import (
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func MustNewID() string {
	return gonanoid.MustGenerate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", 15)
}
