# gocqlxmock [![Go Report Card](https://goreportcard.com/badge/github.com/Guilospanck/gocqlxmock)](https://goreportcard.com/report/github.com/Guilospanck/gocqlxmock) [![CircleCI](https://circleci.com/gh/Guilospanck/gocqlxmock/tree/main.svg?style=svg)](https://circleci.com/gh/Guilospanck/gocqlxmock/tree/main)

Simple mock for [`gocqlx`](https://github.com/scylladb/gocqlx). As `gocqlx` doesn't provide interfaces for its methods and functions in order to make the mock possible, we're using `igocqlx`, which is a wrapper around `gocqlx` that provides interfaces. 

You can find more about `igocqlx` in [this repo](https://github.com/Guilospanck/igocqlx).

## Installation
```bash
go get github.com/Guilospanck/gocqlxmock
```

## How to use
In an application that is using `igocqlx` as a wrapper for `gocqlx`, you can do something like:

```go
// query_builder.go
type queryBuilder struct {
  model   igocqlxtable.ITable
  session igocqlx.ISessionx
  logger  interfaces.ILogger
}

func NewQueryBuider(model igocqlxtable.ITable, session igocqlx.ISessionx, logger interfaces.ILogger) *queryBuilder {
  return &queryBuilder{
    model,
    session,
    logger,
  }
}

func (queryBuilder *queryBuilder) Insert(ctx context.Context, insertData *entities.TrackingData) error {
  insertStatement, insertNames := queryBuilder.model.Insert()
  insertQuery := queryBuilder.session.Query(insertStatement, insertNames).WithContext(ctx)

  err := insertQuery.BindStruct(insertData).ExecRelease()
  if err != nil {
    queryBuilder.logger.Error(fmt.Sprintf("Insert() error %s", err.Error()))
    return err
  }

  return nil
}
```

```go
// query_builder_test.go
func Test_Insert(t *testing.T) {
  t.Run("Should insert data and have no error", func(t *testing.T) {
    // arrange
    stmt := `INSERT INTO tracking_data (first_name,last_name,timestamp,heat,location,speed,telepathy_powers) VALUES (?,?,?,?,?,?,?) `
    names := []string{"first_name", "last_name", "timestamp", "heat", "location", "speed", "telepathy_powers"}
    ctx := context.Background()


    trackingModel := models.NewTrackingDataTable().Table
    loggerSpy := logger.LoggerSpy{}
    sessionMock := &gocqlxmock.SessionxMock{}
    queryMock := &gocqlxmock.QueryxMock{
      Ctx:   ctx,
      Stmt:  stmt,
      Names: names,
    }

    queryBuilder := NewQueryBuider(trackingModel, sessionMock, loggerSpy)

    sessionMock.On("Query", stmt, names).Return(queryMock)
    queryMock.On("WithContext", context.Background()).Return(queryMock)
    queryMock.On("BindStruct", &mocks.CompleteDataEntity).Return(queryMock)
    queryMock.On("ExecRelease").Return(nil)

    // act
    err := queryBuilder.Insert(ctx, &mocks.CompleteDataEntity)

    // assert
    assert.NoError(t, err)
    sessionMock.AssertExpectations(t)
    sessionMock.AssertNumberOfCalls(t, "Query", 1)
    queryMock.AssertNumberOfCalls(t, "WithContext", 1)
    queryMock.AssertNumberOfCalls(t, "BindStruct", 1)
    queryMock.AssertNumberOfCalls(t, "ExecRelease", 1)
    sessionMock.AssertCalled(t, "Query", stmt, names)
    queryMock.AssertCalled(t, "WithContext", context.Background())
    queryMock.AssertCalled(t, "BindStruct", &mocks.CompleteDataEntity)
    queryMock.AssertCalled(t, "ExecRelease")
  })
}
```




