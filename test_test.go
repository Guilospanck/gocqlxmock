package gocqlxmock

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_test(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		// arrange
		const stmt = `INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Jim','Jeffries','2017-11-11 08:05+0000','New York',1.0,3.0,17)`
		names := []string{"test"}
		ctx := context.Background()

		queryMock := &QueryxMock{
			Ctx:   ctx,
			Stmt:  stmt,
			Names: names,
		}
		queryMock.On("WithContext", ctx).Return(queryMock)

		sessionMock := &SessionxMock{}
		sessionMock.On("Query", stmt, names).Return(queryMock)

		// act
		var session ISessionx = sessionMock
		result := session.Query(stmt, names).WithContext(ctx)

		// assert
		queryMock.AssertExpectations(t)
		sessionMock.AssertExpectations(t)
		assert.Equal(t, stmt, result.(*QueryxMock).Stmt)
		assert.Equal(t, ctx, result.(*QueryxMock).Ctx)
		assert.Equal(t, names, result.(*QueryxMock).Names)
	})
}
