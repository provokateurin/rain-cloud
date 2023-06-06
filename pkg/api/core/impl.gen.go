package coreapi

type CoreImpl struct{}

var _ StrictServerInterface = (*CoreImpl)(nil)
