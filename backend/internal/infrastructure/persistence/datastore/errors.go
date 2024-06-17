package datastore

import "golang.org/x/xerrors"

func NewSQLError(err error) error {
	return xerrors.Errorf("failed to SQL execution: %w", err)
}
