package gocqlxmock

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_test(t *testing.T) {
	t.Run("test", func(t *testing.T) {

		const stmt = `INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Jim','Jeffries','2017-11-11 08:05+0000','New York',1.0,3.0,17)`
		names := []string{"test"}
		ctx := context.Background()

		session := &SessionxMock{}
		queryMock := &QueryxMock{
			Ctx:   ctx,
			Stmt:  stmt,
			Names: names,
		}

		session.On("Query", stmt, names).Return(queryMock)
		queryMock.On("WithContext", context.Background()).Return(queryMock)

		result := session.Query(stmt, names).WithContext(context.Background())

		session.AssertExpectations(t)
		queryMock.AssertExpectations(t)
		assert.Equal(t, stmt, result.(*QueryxMock).Stmt)
		assert.Equal(t, names, result.(*QueryxMock).Names)
		assert.Equal(t, ctx, result.(*QueryxMock).Ctx)

	})
}
