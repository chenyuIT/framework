package console

import (
	"testing"

	"github.com/stretchr/testify/assert"

	consolemocks "github.com/chenyuIT/framework/contracts/console/mocks"
	"github.com/chenyuIT/framework/support/file"
)

func TestEventMakeCommand(t *testing.T) {
	eventMakeCommand := &EventMakeCommand{}
	mockContext := &consolemocks.Context{}
	mockContext.On("Argument", 0).Return("").Once()
	err := eventMakeCommand.Handle(mockContext)
	assert.EqualError(t, err, "Not enough arguments (missing: name) ")

	mockContext.On("Argument", 0).Return("GoravelEvent").Once()
	err = eventMakeCommand.Handle(mockContext)
	assert.Nil(t, err)
	assert.True(t, file.Exists("app/events/goravel_event.go"))
	assert.True(t, file.Remove("app"))
}
