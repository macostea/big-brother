package model

type CollectedDataItem interface {
	Type() string
	GetInfo()
}
