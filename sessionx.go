package gocqlxmock

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/stretchr/testify/mock"
)

type ISessionx interface {
	ContextQuery(ctx context.Context, stmt string, names []string) IQueryx
	Query(stmt string, names []string) IQueryx
	ExecStmt(stmt string) error
}

type SessionxMock struct {
	*gocql.Session
	mock.Mock
}

func (mock *SessionxMock) ContextQuery(ctx context.Context, stmt string, names []string) IQueryx {
	args := mock.Called(ctx, stmt, names)

	return args.Get(0).(IQueryx)
}

func (mock *SessionxMock) Query(stmt string, names []string) IQueryx {
	args := mock.Called(stmt, names)

	return args.Get(0).(IQueryx)
}

func (mock *SessionxMock) ExecStmt(stmt string) error {
	args := mock.Called(stmt)

	return args.Error(0)
}
