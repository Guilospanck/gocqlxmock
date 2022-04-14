package gocqlxmock

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type sessionXSut struct {
	ctx          context.Context
	stmt         string
	names        []string
	sessionxmock *SessionxMock
	querymock    *QueryxMock
}

func makeSessionxSut() sessionXSut {
	ctx := context.Background()
	stmt := "statement"
	names := []string{"name1", "name2"}

	sessionxmock := &SessionxMock{}
	querymock := &QueryxMock{}

	return sessionXSut{
		ctx,
		stmt,
		names,
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
