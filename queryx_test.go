package gocqlxmock

import (
	"context"
	"fmt"
	"testing"
	"time"

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
	timestamp   int64
	routingKey  []byte
	ctx         context.Context
	p           float64
	retry       gocql.RetryPolicy
	sp          gocql.SpeculativeExecutionPolicy
	cons        gocql.SerialConsistency
	state       []byte
}

type tracer struct{}

func (tracer *tracer) Trace(traceId []byte) {}

type observerstruct struct{}

func (observer *observerstruct) ObserveQuery(context.Context, gocql.ObservedQuery) {}

type retrypolicystruct struct{}

func (rp *retrypolicystruct) Attempt(gocql.RetryableQuery) bool  { return true }
func (rp *retrypolicystruct) GetRetryType(error) gocql.RetryType { return gocql.RetryType(10) }

type spstruct struct{}

func (sp *spstruct) Attempts() int        { return 1 }
func (sp *spstruct) Delay() time.Duration { return time.Duration(1) }

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
	timestamp := int64(10)
	routingKey := []byte{}
	ctx := context.Background()
	p := float64(10.1)
	retry := &retrypolicystruct{}
	sp := &spstruct{}
	cons := gocql.SerialConsistency(10)
	state := []byte{}

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
		timestamp,
		routingKey,
		ctx,
		p,
		retry,
		sp,
		cons,
		state,
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

func Test_Queryx_ExecCASRelease(t *testing.T) {
	t.Run("Should call ExecCASRelease with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("ExecCASRelease").Return(sut.boolVar, nil)

		// act
		result, err := sut.queryxmock.ExecCASRelease()

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "ExecCASRelease")
		sut.queryxmock.AssertNumberOfCalls(t, "ExecCASRelease", 1)
		assert.NoError(t, err)
		assert.Equal(t, sut.boolVar, result)
	})

	t.Run("Should call ExecCASRelease with proper parameters and return err", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("ExecCASRelease").Return(sut.boolVar, sut.err)

		// act
		result, err := sut.queryxmock.ExecCASRelease()

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "ExecCASRelease")
		sut.queryxmock.AssertNumberOfCalls(t, "ExecCASRelease", 1)
		assert.Error(t, err, sut.errMsg)
		assert.Equal(t, sut.boolVar, result)
	})
}

func Test_Queryx_Get(t *testing.T) {
	t.Run("Should call Get with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		arg := makeArg("potato")
		sut.queryxmock.On("Get", arg).Return(nil)

		// act
		err := sut.queryxmock.Get(arg)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "Get", arg)
		sut.queryxmock.AssertNumberOfCalls(t, "Get", 1)
		assert.NoError(t, err)
	})

	t.Run("Should call Get with proper parameters and return err", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		arg := makeArg("potato")
		sut.queryxmock.On("Get", arg).Return(sut.err)

		// act
		err := sut.queryxmock.Get(arg)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "Get", arg)
		sut.queryxmock.AssertNumberOfCalls(t, "Get", 1)
		assert.Error(t, err, sut.errMsg)
	})
}

func Test_Queryx_GetRelease(t *testing.T) {
	t.Run("Should call GetRelease with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		arg := makeArg("potato")
		sut.queryxmock.On("GetRelease", arg).Return(nil)

		// act
		err := sut.queryxmock.GetRelease(arg)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "GetRelease", arg)
		sut.queryxmock.AssertNumberOfCalls(t, "GetRelease", 1)
		assert.NoError(t, err)
	})

	t.Run("Should call GetRelease with proper parameters and return err", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		arg := makeArg("potato")
		sut.queryxmock.On("GetRelease", arg).Return(sut.err)

		// act
		err := sut.queryxmock.GetRelease(arg)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "GetRelease", arg)
		sut.queryxmock.AssertNumberOfCalls(t, "GetRelease", 1)
		assert.Error(t, err, sut.errMsg)
	})
}

func Test_Queryx_GetCAS(t *testing.T) {
	t.Run("Should call GetCAS with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		arg := makeArg("potato")
		sut.queryxmock.On("GetCAS", arg).Return(sut.boolVar, nil)

		// act
		result, err := sut.queryxmock.GetCAS(arg)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "GetCAS", arg)
		sut.queryxmock.AssertNumberOfCalls(t, "GetCAS", 1)
		assert.Equal(t, sut.boolVar, result)
		assert.NoError(t, err)
	})

	t.Run("Should call GetCAS with proper parameters and return err", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		arg := makeArg("potato")
		sut.queryxmock.On("GetCAS", arg).Return(sut.boolVar, sut.err)

		// act
		result, err := sut.queryxmock.GetCAS(arg)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "GetCAS", arg)
		sut.queryxmock.AssertNumberOfCalls(t, "GetCAS", 1)
		assert.Equal(t, sut.boolVar, result)
		assert.Error(t, err, sut.errMsg)
	})
}

func Test_Queryx_GetCASRelease(t *testing.T) {
	t.Run("Should call GetCASRelease with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		arg := makeArg("potato")
		sut.queryxmock.On("GetCASRelease", arg).Return(sut.boolVar, nil)

		// act
		result, err := sut.queryxmock.GetCASRelease(arg)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "GetCASRelease", arg)
		sut.queryxmock.AssertNumberOfCalls(t, "GetCASRelease", 1)
		assert.Equal(t, sut.boolVar, result)
		assert.NoError(t, err)
	})

	t.Run("Should call GetCASRelease with proper parameters and return err", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		arg := makeArg("potato")
		sut.queryxmock.On("GetCASRelease", arg).Return(sut.boolVar, sut.err)

		// act
		result, err := sut.queryxmock.GetCASRelease(arg)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "GetCASRelease", arg)
		sut.queryxmock.AssertNumberOfCalls(t, "GetCASRelease", 1)
		assert.Equal(t, sut.boolVar, result)
		assert.Error(t, err, sut.errMsg)
	})
}

func Test_Queryx_Select(t *testing.T) {
	t.Run("Should call Select with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		arg := makeArg("potato")
		sut.queryxmock.On("Select", arg).Return(nil)

		// act
		err := sut.queryxmock.Select(arg)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "Select", arg)
		sut.queryxmock.AssertNumberOfCalls(t, "Select", 1)
		assert.NoError(t, err)
	})

	t.Run("Should call Select with proper parameters and return err", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		arg := makeArg("potato")
		sut.queryxmock.On("Select", arg).Return(sut.err)

		// act
		err := sut.queryxmock.Select(arg)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "Select", arg)
		sut.queryxmock.AssertNumberOfCalls(t, "Select", 1)
		assert.Error(t, err, sut.errMsg)
	})
}

func Test_Queryx_SelectRelease(t *testing.T) {
	t.Run("Should call SelectRelease with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		arg := makeArg("potato")
		sut.queryxmock.On("SelectRelease", arg).Return(nil)

		// act
		err := sut.queryxmock.SelectRelease(arg)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "SelectRelease", arg)
		sut.queryxmock.AssertNumberOfCalls(t, "SelectRelease", 1)
		assert.NoError(t, err)
	})

	t.Run("Should call SelectRelease with proper parameters and return err", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		arg := makeArg("potato")
		sut.queryxmock.On("SelectRelease", arg).Return(sut.err)

		// act
		err := sut.queryxmock.SelectRelease(arg)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "SelectRelease", arg)
		sut.queryxmock.AssertNumberOfCalls(t, "SelectRelease", 1)
		assert.Error(t, err, sut.errMsg)
	})
}

func Test_Queryx_Iter(t *testing.T) {
	t.Run("Should call Iter with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("Iter").Return(sut.iterxmock)

		// act
		result := sut.queryxmock.Iter()

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "Iter")
		sut.queryxmock.AssertNumberOfCalls(t, "Iter", 1)
		assert.Equal(t, sut.iterxmock, result)
	})
}

func Test_Queryx_Consistency(t *testing.T) {
	t.Run("Should call Consistency with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("Consistency", sut.consistency).Return(sut.queryxmock)

		// act
		result := sut.queryxmock.Consistency(sut.consistency)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "Consistency", sut.consistency)
		sut.queryxmock.AssertNumberOfCalls(t, "Consistency", 1)
		assert.Equal(t, sut.queryxmock, result)
	})
}

func Test_Queryx_CustomPayload(t *testing.T) {
	t.Run("Should call CustomPayload with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		customPayload := makeCustomPayload("potato")
		sut.queryxmock.On("CustomPayload", customPayload).Return(sut.queryxmock)

		// act
		result := sut.queryxmock.CustomPayload(customPayload)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "CustomPayload", customPayload)
		sut.queryxmock.AssertNumberOfCalls(t, "CustomPayload", 1)
		assert.Equal(t, sut.queryxmock, result)
	})
}

func Test_Queryx_Trace(t *testing.T) {
	t.Run("Should call Trace with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("Trace", sut.trace).Return(sut.queryxmock)

		// act
		result := sut.queryxmock.Trace(sut.trace)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "Trace", sut.trace)
		sut.queryxmock.AssertNumberOfCalls(t, "Trace", 1)
		assert.Equal(t, sut.queryxmock, result)
	})
}

func Test_Queryx_Observer(t *testing.T) {
	t.Run("Should call Observer with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("Observer", sut.observer).Return(sut.queryxmock)

		// act
		result := sut.queryxmock.Observer(sut.observer)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "Observer", sut.observer)
		sut.queryxmock.AssertNumberOfCalls(t, "Observer", 1)
		assert.Equal(t, sut.queryxmock, result)
	})
}

func Test_Queryx_PageSize(t *testing.T) {
	t.Run("Should call PageSize with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("PageSize", sut.pagesize_n).Return(sut.queryxmock)

		// act
		result := sut.queryxmock.PageSize(sut.pagesize_n)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "PageSize", sut.pagesize_n)
		sut.queryxmock.AssertNumberOfCalls(t, "PageSize", 1)
		assert.Equal(t, sut.queryxmock, result)
	})
}

func Test_Queryx_DefaultTimestamp(t *testing.T) {
	t.Run("Should call DefaultTimestamp with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("DefaultTimestamp", sut.boolVar).Return(sut.queryxmock)

		// act
		result := sut.queryxmock.DefaultTimestamp(sut.boolVar)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "DefaultTimestamp", sut.boolVar)
		sut.queryxmock.AssertNumberOfCalls(t, "DefaultTimestamp", 1)
		assert.Equal(t, sut.queryxmock, result)
	})
}

func Test_Queryx_WithTimestamp(t *testing.T) {
	t.Run("Should call WithTimestamp with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("WithTimestamp", sut.timestamp).Return(sut.queryxmock)

		// act
		result := sut.queryxmock.WithTimestamp(sut.timestamp)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "WithTimestamp", sut.timestamp)
		sut.queryxmock.AssertNumberOfCalls(t, "WithTimestamp", 1)
		assert.Equal(t, sut.queryxmock, result)
	})
}

func Test_Queryx_RoutingKey(t *testing.T) {
	t.Run("Should call RoutingKey with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("RoutingKey", sut.routingKey).Return(sut.queryxmock)

		// act
		result := sut.queryxmock.RoutingKey(sut.routingKey)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "RoutingKey", sut.routingKey)
		sut.queryxmock.AssertNumberOfCalls(t, "RoutingKey", 1)
		assert.Equal(t, sut.queryxmock, result)
	})
}

func Test_Queryx_WithContext(t *testing.T) {
	t.Run("Should call WithContext with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("WithContext", sut.ctx).Return(sut.queryxmock)

		// act
		result := sut.queryxmock.WithContext(sut.ctx)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "WithContext", sut.ctx)
		sut.queryxmock.AssertNumberOfCalls(t, "WithContext", 1)
		assert.Equal(t, sut.queryxmock, result)
	})
}

func Test_Queryx_Prefetch(t *testing.T) {
	t.Run("Should call Prefetch with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("Prefetch", sut.p).Return(sut.queryxmock)

		// act
		result := sut.queryxmock.Prefetch(sut.p)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "Prefetch", sut.p)
		sut.queryxmock.AssertNumberOfCalls(t, "Prefetch", 1)
		assert.Equal(t, sut.queryxmock, result)
	})
}

func Test_Queryx_RetryPolicy(t *testing.T) {
	t.Run("Should call RetryPolicy with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("RetryPolicy", sut.retry).Return(sut.queryxmock)

		// act
		result := sut.queryxmock.RetryPolicy(sut.retry)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "RetryPolicy", sut.retry)
		sut.queryxmock.AssertNumberOfCalls(t, "RetryPolicy", 1)
		assert.Equal(t, sut.queryxmock, result)
	})
}

func Test_Queryx_SetSpeculativeExecutionPolicy(t *testing.T) {
	t.Run("Should call SetSpeculativeExecutionPolicy with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("SetSpeculativeExecutionPolicy", sut.sp).Return(sut.queryxmock)

		// act
		result := sut.queryxmock.SetSpeculativeExecutionPolicy(sut.sp)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "SetSpeculativeExecutionPolicy", sut.sp)
		sut.queryxmock.AssertNumberOfCalls(t, "SetSpeculativeExecutionPolicy", 1)
		assert.Equal(t, sut.queryxmock, result)
	})
}

func Test_Queryx_Idempotent(t *testing.T) {
	t.Run("Should call Idempotent with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("Idempotent", sut.boolVar).Return(sut.queryxmock)

		// act
		result := sut.queryxmock.Idempotent(sut.boolVar)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "Idempotent", sut.boolVar)
		sut.queryxmock.AssertNumberOfCalls(t, "Idempotent", 1)
		assert.Equal(t, sut.queryxmock, result)
	})
}

func Test_Queryx_SerialConsistency(t *testing.T) {
	t.Run("Should call SerialConsistency with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("SerialConsistency", sut.cons).Return(sut.queryxmock)

		// act
		result := sut.queryxmock.SerialConsistency(sut.cons)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "SerialConsistency", sut.cons)
		sut.queryxmock.AssertNumberOfCalls(t, "SerialConsistency", 1)
		assert.Equal(t, sut.queryxmock, result)
	})
}

func Test_Queryx_PageState(t *testing.T) {
	t.Run("Should call PageState with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("PageState", sut.state).Return(sut.queryxmock)

		// act
		result := sut.queryxmock.PageState(sut.state)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "PageState", sut.state)
		sut.queryxmock.AssertNumberOfCalls(t, "PageState", 1)
		assert.Equal(t, sut.queryxmock, result)
	})
}

func Test_Queryx_NoSkipMetadata(t *testing.T) {
	t.Run("Should call NoSkipMetadata with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("NoSkipMetadata").Return(sut.queryxmock)

		// act
		result := sut.queryxmock.NoSkipMetadata()

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "NoSkipMetadata")
		sut.queryxmock.AssertNumberOfCalls(t, "NoSkipMetadata", 1)
		assert.Equal(t, sut.queryxmock, result)
	})
}

func Test_Queryx_Release(t *testing.T) {
	t.Run("Should call Release with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		sut.queryxmock.On("Release").Return(sut.queryxmock)

		// act
		sut.queryxmock.Release()

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "Release")
		sut.queryxmock.AssertNumberOfCalls(t, "Release", 1)
	})
}

func Test_Queryx_Scan(t *testing.T) {
	t.Run("Should call Scan with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		arg := makeArg("potato")
		sut.queryxmock.On("Scan", arg).Return(nil)

		// act
		err := sut.queryxmock.Scan(arg)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "Scan", arg)
		sut.queryxmock.AssertNumberOfCalls(t, "Scan", 1)
		assert.NoError(t, err)
	})

	t.Run("Should call Scan with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeQueryxSut()
		arg := makeArg("potato")
		sut.queryxmock.On("Scan", arg).Return(sut.err)

		// act
		err := sut.queryxmock.Scan(arg)

		// assert
		sut.queryxmock.AssertExpectations(t)
		sut.queryxmock.AssertCalled(t, "Scan", arg)
		sut.queryxmock.AssertNumberOfCalls(t, "Scan", 1)
		assert.Error(t, err, sut.errMsg)
	})
}
