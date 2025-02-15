package parser

type Parser func(content []byte) (string, error)
