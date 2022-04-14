package gocqlxmock

import (
	"context"
	"fmt"
	"testing"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type Potato struct {
	Name string
}

func makeArg(name string) *Potato {
	return &Potato{
		Name: name,
	}
}

func makeMapArg(name string) map[string]interface{} {
	mapped := make(map[string]interface{})

	potato := &Potato{
		Name: name,
	}

	mapped[name] = potato

	return mapped
}

func makeCustomPayload(name string) map[string][]byte {
	mapped := make(map[string][]byte)

	mapped[name] = []byte{}

	return mapped
}

type queryXSut struct {
	tr          gocqlx.Transformer
	err         error
	errMsg      string
	boolVar     bool
	iterxmock   *IterxMock
	queryxmock  *QueryxMock
	consistency gocql.Consistency
	trace       gocql.Tracer
	observer    gocql.QueryObserver
	pagesize_n  int
}

type tracer struct{}

func (tracer *tracer) Trace(traceId []byte) {}

type observerstruct struct{}

func (observer *observerstruct) ObserveQuery(context.Context, gocql.ObservedQuery) {}

func makeQueryxSut() queryXSut {
	tr := func(name string, val interface{}) interface{} { return nil }
	errMsg := "queryx_error"
	err := fmt.Errorf(errMsg)
	boolVar := true
	iterxmock := &IterxMock{}
	queryxmock := &QueryxMock{}
	consistency := gocql.Consistency(10)
	var trace gocql.Tracer = &tracer{}
	var observer gocql.QueryObserver = &observerstruct{}
	pagesize_n := 10

	return queryXSut{
		tr,
		err,
		errMsg,
		boolVar,
		iterxmock,
		queryxmock,
		consistency,
		trace,
		observer,
		pagesize_n,
	}
}

func Test_Queryx_WithBindTransformer(t *testing.T) {
	t.Run("Should call WithBindTransformer with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		trExpect := mock.AnythingOfType("gocqlx.Transformer")
		sut.queryxmock.On("WithBindTransformer", trExpect).Return(sut.queryxmock)

		// act
		result := sut.queryxmock.WithBindTransformer(sut.tr)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertNumberOfCalls(t, "WithBindTransformer", 1)
		assert.Equal(t, result, sut.queryxmock)
	})
}

func Test_Queryx_BindStruct(t *testing.T) {
	t.Run("Should call BindStruct with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		arg := makeArg("potato")
		sut.queryxmock.On("BindStruct", arg).Return(sut.queryxmock)

		// act
		result := sut.queryxmock.BindStruct(arg)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "BindStruct", arg)
		sut.queryxmock.AssertNumberOfCalls(t, "BindStruct", 1)
		assert.Equal(t, result, sut.queryxmock)
	})
}

func Test_Queryx_BindStructMap(t *testing.T) {
	t.Run("Should call BindStructMap with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		arg0 := makeArg("potato")
		arg1 := makeMapArg("larry")
		sut.queryxmock.On("BindStructMap", arg0, arg1).Return(sut.queryxmock)

		// act
		result := sut.queryxmock.BindStructMap(arg0, arg1)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "BindStructMap", arg0, arg1)
		sut.queryxmock.AssertNumberOfCalls(t, "BindStructMap", 1)
		assert.Equal(t, result, sut.queryxmock)
	})
}

func Test_Queryx_BindMap(t *testing.T) {
	t.Run("Should call BindMap with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		arg := makeMapArg("larry")
		sut.queryxmock.On("BindMap", arg).Return(sut.queryxmock)

		// act
		result := sut.queryxmock.BindMap(arg)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "BindMap", arg)
		sut.queryxmock.AssertNumberOfCalls(t, "BindMap", 1)
		assert.Equal(t, result, sut.queryxmock)
	})
}

func Test_Queryx_Bind(t *testing.T) {
	t.Run("Should call Bind with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		arg := makeArg("larry")
		sut.queryxmock.On("Bind", []interface{}{arg}).Return(sut.queryxmock)

		// act
		result := sut.queryxmock.Bind(arg)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "Bind", []interface{}{arg})
		sut.queryxmock.AssertNumberOfCalls(t, "Bind", 1)
		assert.Equal(t, result, sut.queryxmock)
	})
}

func Test_Queryx_Err(t *testing.T) {
	t.Run("Should call Err with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("Err").Return(nil)

		// act
		result := sut.queryxmock.Err()

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "Err")
		sut.queryxmock.AssertNumberOfCalls(t, "Err", 1)
		assert.NoError(t, result)
	})

	t.Run("Should call Err with proper parameters and return err", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("Err").Return(sut.err)

		// act
		result := sut.queryxmock.Err()

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "Err")
		sut.queryxmock.AssertNumberOfCalls(t, "Err", 1)
		assert.Error(t, result, sut.errMsg)
	})
}

func Test_Queryx_Exec(t *testing.T) {
	t.Run("Should call Exec with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("Exec").Return(nil)

		// act
		result := sut.queryxmock.Exec()

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "Exec")
		sut.queryxmock.AssertNumberOfCalls(t, "Exec", 1)
		assert.NoError(t, result)
	})

	t.Run("Should call Exec with proper parameters and return err", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("Exec").Return(sut.err)

		// act
		result := sut.queryxmock.Exec()

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "Exec")
		sut.queryxmock.AssertNumberOfCalls(t, "Exec", 1)
		assert.Error(t, result, sut.errMsg)
	})
}

func Test_Queryx_ExecRelease(t *testing.T) {
	t.Run("Should call ExecRelease with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("ExecRelease").Return(nil)

		// act
		result := sut.queryxmock.ExecRelease()

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "ExecRelease")
		sut.queryxmock.AssertNumberOfCalls(t, "ExecRelease", 1)
		assert.NoError(t, result)
	})

	t.Run("Should call ExecRelease with proper parameters and return err", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("ExecRelease").Return(sut.err)

		// act
		result := sut.queryxmock.ExecRelease()

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "ExecRelease")
		sut.queryxmock.AssertNumberOfCalls(t, "ExecRelease", 1)
		assert.Error(t, result, sut.errMsg)
	})
}

func Test_Queryx_ExecCAS(t *testing.T) {
	t.Run("Should call ExecCAS with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("ExecCAS").Return(sut.boolVar, nil)

		// act
		result, err := sut.queryxmock.ExecCAS()

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "ExecCAS")
		sut.queryxmock.AssertNumberOfCalls(t, "ExecCAS", 1)
		assert.NoError(t, err)
		assert.Equal(t, sut.boolVar, result)
	})

	t.Run("Should call ExecCAS with proper parameters and return err", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("ExecCAS").Return(sut.boolVar, sut.err)

		// act
		result, err := sut.queryxmock.ExecCAS()

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "ExecCAS")
		sut.queryxmock.AssertNumberOfCalls(t, "ExecCAS", 1)
		assert.Error(t, err, sut.errMsg)
		assert.Equal(t, sut.boolVar, result)
	})
}
