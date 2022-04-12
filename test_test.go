package gocqlxmock

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_test(t *testing.T) {
	t.Run("test", func(t *testing.T) {

		const insertStatement = `INSERT INTO tracking.tracking_data ("first_name","last_name","timestamp","location","speed","heat","telepathy_powers") VALUES ('Jim','Jeffries','2017-11-11 08:05+0000','New York',1.0,3.0,17)`
		insertNames := []string{"test"}

		session := &SessionxMock{}
		session.On("Query", insertStatement, insertNames).Return(&Queryx{stmt: insertStatement, names: insertNames})

		insertQuery := session.Query(insertStatement, insertNames).WithContext(context.Background())

		// mock Queryx

		type Test struct {
			Name string
		}

		test := &Test{}

		err := insertQuery.BindStruct(test).ExecRelease()

		assert.Nil(t, err)

	})
}
