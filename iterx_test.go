package gocqlxmock

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type iterXSut struct {
	boolVar   bool
	err       error
	errMsg    string
	iterxmock *IterxMock
}

func makeIterxSut() iterXSut {
	boolVar := true
	errMsg := "sessionx_error"
	err := fmt.Errorf(errMsg)

	iterxmock := &IterxMock{}

	return iterXSut{
		boolVar,
		err,
		errMsg,
		iterxmock,
	}
}

func Test_Iterx_Unsafe(t *testing.T) {
	t.Run("Should call Unsafe with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeIterxSut()
		sut.iterxmock.On("Unsafe").Return(sut.iterxmock)

		// act
		result := sut.iterxmock.Unsafe()

		// assert
		sut.iterxmock.AssertExpectations(t)
		sut.iterxmock.AssertCalled(t, "Unsafe")
		sut.iterxmock.AssertNumberOfCalls(t, "Unsafe", 1)
		assert.Equal(t, result, sut.iterxmock)
	})
}

func Test_Iterx_StructOnly(t *testing.T) {
	t.Run("Should call StructOnly with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeIterxSut()
		sut.iterxmock.On("StructOnly").Return(sut.iterxmock)

		// act
		result := sut.iterxmock.StructOnly()

		// assert
		sut.iterxmock.AssertExpectations(t)
		sut.iterxmock.AssertCalled(t, "StructOnly")
		sut.iterxmock.AssertNumberOfCalls(t, "StructOnly", 1)
		assert.Equal(t, result, sut.iterxmock)
	})
}

func Test_Iterx_Get(t *testing.T) {
	t.Run("Should call Get with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeIterxSut()
		arg := makeArg("potato")
		sut.iterxmock.On("Get", arg).Return(nil)

		// act
		err := sut.iterxmock.Get(arg)

		// assert
		sut.iterxmock.AssertExpectations(t)
		sut.iterxmock.AssertCalled(t, "Get", arg)
		sut.iterxmock.AssertNumberOfCalls(t, "Get", 1)
		assert.NoError(t, err)
	})

	t.Run("Should call Get with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeIterxSut()
		arg := makeArg("potato")
		sut.iterxmock.On("Get", arg).Return(sut.err)

		// act
		err := sut.iterxmock.Get(arg)

		// assert
		sut.iterxmock.AssertExpectations(t)
		sut.iterxmock.AssertCalled(t, "Get", arg)
		sut.iterxmock.AssertNumberOfCalls(t, "Get", 1)
		assert.Error(t, err, sut.errMsg)
	})
}

func Test_Iterx_Select(t *testing.T) {
	t.Run("Should call Select with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeIterxSut()
		arg := makeArg("potato")
		sut.iterxmock.On("Select", arg).Return(nil)

		// act
		err := sut.iterxmock.Select(arg)

		// assert
		sut.iterxmock.AssertExpectations(t)
		sut.iterxmock.AssertCalled(t, "Select", arg)
		sut.iterxmock.AssertNumberOfCalls(t, "Select", 1)
		assert.NoError(t, err)
	})

	t.Run("Should call Select with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeIterxSut()
		arg := makeArg("potato")
		sut.iterxmock.On("Select", arg).Return(sut.err)

		// act
		err := sut.iterxmock.Select(arg)

		// assert
		sut.iterxmock.AssertExpectations(t)
		sut.iterxmock.AssertCalled(t, "Select", arg)
		sut.iterxmock.AssertNumberOfCalls(t, "Select", 1)
		assert.Error(t, err, sut.errMsg)
	})
}

func Test_Iterx_StructScan(t *testing.T) {
	t.Run("Should call StructScan with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeIterxSut()
		arg := makeArg("potato")
		sut.iterxmock.On("StructScan", arg).Return(sut.boolVar)

		// act
		result := sut.iterxmock.StructScan(arg)

		// assert
		sut.iterxmock.AssertExpectations(t)
		sut.iterxmock.AssertCalled(t, "StructScan", arg)
		sut.iterxmock.AssertNumberOfCalls(t, "StructScan", 1)
		assert.Equal(t, result, sut.boolVar)
	})
}

func Test_Iterx_Scan(t *testing.T) {
	t.Run("Should call Scan with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeIterxSut()
		arg := makeArg("potato")
		sut.iterxmock.On("Scan", []interface{}{arg}).Return(sut.boolVar)

		// act
		result := sut.iterxmock.Scan(arg)

		// assert
		sut.iterxmock.AssertExpectations(t)
		sut.iterxmock.AssertCalled(t, "Scan", []interface{}{arg})
		sut.iterxmock.AssertNumberOfCalls(t, "Scan", 1)
		assert.Equal(t, result, sut.boolVar)
	})
}

func Test_Iterx_Close(t *testing.T) {
	t.Run("Should call Close with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeIterxSut()
		sut.iterxmock.On("Close").Return(nil)

		// act
		err := sut.iterxmock.Close()

		// assert
		sut.iterxmock.AssertExpectations(t)
		sut.iterxmock.AssertCalled(t, "Close")
		sut.iterxmock.AssertNumberOfCalls(t, "Close", 1)
		assert.NoError(t, err)
	})

	t.Run("Should call Close with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeIterxSut()
		sut.iterxmock.On("Close").Return(sut.err)

		// act
		err := sut.iterxmock.Close()

		// assert
		sut.iterxmock.AssertExpectations(t)
		sut.iterxmock.AssertCalled(t, "Close")
		sut.iterxmock.AssertNumberOfCalls(t, "Close", 1)
		assert.Error(t, err, sut.errMsg)
	})
}

func Test_Iterx_MapScan(t *testing.T) {
	t.Run("Should call MapScan with proper parameters and return proper result", func(t *testing.T) {
		// arrange
		sut := makeIterxSut()
		arg := makeMapArg("potato")
		sut.iterxmock.On("MapScan", arg).Return(sut.boolVar)

		// act
		result := sut.iterxmock.MapScan(arg)

		// assert
		sut.iterxmock.AssertExpectations(t)
		sut.iterxmock.AssertCalled(t, "MapScan", arg)
		sut.iterxmock.AssertNumberOfCalls(t, "MapScan", 1)
		assert.Equal(t, result, sut.boolVar)
	})
}
