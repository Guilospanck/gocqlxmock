package gocqlxmock

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type sessionXSut struct {
	ctx          context.Context
	stmt         string
	names        []string
	err          error
	errMsg       string
	sessionxmock *SessionxMock
	querymock    *QueryxMock
}

func makeSessionxSut() sessionXSut {
	ctx := context.Background()
	stmt := "statement"
	names := []string{"name1", "name2"}
	errMsg := "sessionx_error"
	err := fmt.Errorf(errMsg)

	sessionxmock := &SessionxMock{}
	querymock := &QueryxMock{}

	return sessionXSut{
		ctx,
		stmt,
		names,
		err,
		errMsg,
		sessionxmock,
		querymock,
	}
}

func Test_Sessionx_ContextQuery(t *testing.T) {
	t.Run("Should call ContextQuery with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeSessionxSut()
		sut.sessionxmock.On("ContextQuery", sut.ctx, sut.stmt, sut.names).Return(sut.querymock)

		// act
		result := sut.sessionxmock.ContextQuery(sut.ctx, sut.stmt, sut.names)

		// assert
		sut.sessionxmock.AssertExpectations(t)
		sut.sessionxmock.AssertCalled(t, "ContextQuery", sut.ctx, sut.stmt, sut.names)
		sut.sessionxmock.AssertNumberOfCalls(t, "ContextQuery", 1)
		assert.Equal(t, result, sut.querymock)
	})
}

func Test_Sessionx_Query(t *testing.T) {
	t.Run("Should call Query with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeSessionxSut()
		sut.sessionxmock.On("Query", sut.stmt, sut.names).Return(sut.querymock)

		// act
		result := sut.sessionxmock.Query(sut.stmt, sut.names)

		// assert
		sut.sessionxmock.AssertExpectations(t)
		sut.sessionxmock.AssertCalled(t, "Query", sut.stmt, sut.names)
		sut.sessionxmock.AssertNumberOfCalls(t, "Query", 1)
		assert.Equal(t, result, sut.querymock)
	})
}

func Test_Sessionx_ExecStmt(t *testing.T) {
	t.Run("Should call ExecStmt with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeSessionxSut()
		sut.sessionxmock.On("ExecStmt", sut.stmt).Return(nil)

		// act
		result := sut.sessionxmock.ExecStmt(sut.stmt)

		// assert
		sut.sessionxmock.AssertExpectations(t)
		sut.sessionxmock.AssertCalled(t, "ExecStmt", sut.stmt)
		sut.sessionxmock.AssertNumberOfCalls(t, "ExecStmt", 1)
		assert.NoError(t, result)
	})

	t.Run("Should call ExecStmt with proper parameters and return error", func(t *testing.T) {
		// arrange
		sut := makeSessionxSut()
		sut.sessionxmock.On("ExecStmt", sut.stmt).Return(sut.err)

		// act
		result := sut.sessionxmock.ExecStmt(sut.stmt)

		// assert
		sut.sessionxmock.AssertExpectations(t)
		sut.sessionxmock.AssertCalled(t, "ExecStmt", sut.stmt)
		sut.sessionxmock.AssertNumberOfCalls(t, "ExecStmt", 1)
		assert.Error(t, result, sut.errMsg)
	})
}

func Test_Sessionx_AwaitSchemaAgreement(t *testing.T) {
	t.Run("Should call AwaitSchemaAgreement with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeSessionxSut()
		sut.sessionxmock.On("AwaitSchemaAgreement", sut.ctx).Return(nil)

		// act
		result := sut.sessionxmock.AwaitSchemaAgreement(sut.ctx)

		// assert
		sut.sessionxmock.AssertExpectations(t)
		sut.sessionxmock.AssertCalled(t, "AwaitSchemaAgreement", sut.ctx)
		sut.sessionxmock.AssertNumberOfCalls(t, "AwaitSchemaAgreement", 1)
		assert.NoError(t, result)
	})

	t.Run("Should call AwaitSchemaAgreement with proper parameters and return error", func(t *testing.T) {
		// arrange
		sut := makeSessionxSut()
		sut.sessionxmock.On("AwaitSchemaAgreement", sut.ctx).Return(sut.err)

		// act
		result := sut.sessionxmock.AwaitSchemaAgreement(sut.ctx)

		// assert
		sut.sessionxmock.AssertExpectations(t)
		sut.sessionxmock.AssertCalled(t, "AwaitSchemaAgreement", sut.ctx)
		sut.sessionxmock.AssertNumberOfCalls(t, "AwaitSchemaAgreement", 1)
		assert.Error(t, result, sut.errMsg)
	})
}

func Test_Sessionx_Close(t *testing.T) {
	t.Run("Should call Close with proper parameters and return void", func(t *testing.T) {
		// arrange
		sut := makeSessionxSut()
		sut.sessionxmock.On("Close").Return()

		// act
		sut.sessionxmock.Close()

		// assert
		sut.sessionxmock.AssertExpectations(t)
		sut.sessionxmock.AssertCalled(t, "Close")
		sut.sessionxmock.AssertNumberOfCalls(t, "Close", 1)
	})
}
