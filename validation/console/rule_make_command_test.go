package console

import (
	"testing"

	consolemocks "github.com/chenyuIT/framework/contracts/console/mocks"
	"github.com/chenyuIT/framework/support/file"

	"github.com/stretchr/testify/assert"
)

func TestRuleMakeCommand(t *testing.T) {
	requestMakeCommand := &RuleMakeCommand{}
	mockContext := &consolemocks.Context{}
	mockContext.On("Argument", 0).Return("").Once()
	err := requestMakeCommand.Handle(mockContext)
	assert.EqualError(t, err, "Not enough arguments (missing: name) ")

	mockContext.On("Argument", 0).Return("Uppercase").Once()
	err = requestMakeCommand.Handle(mockContext)
	assert.Nil(t, err)
	assert.True(t, file.Exists("app/rules/uppercase.go"))
	assert.True(t, file.Remove("app"))
}
