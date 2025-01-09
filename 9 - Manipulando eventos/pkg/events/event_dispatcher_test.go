package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type EventHandlerTest struct {
	Id string
}

func (h *EventHandlerTest) Handle(event EventInterface, wg *sync.WaitGroup) {
	wg.Done()
}

func (h *EventHandlerTest) GetId() string {
	return h.Id

}

type EventTest struct {
	Name    string
	Payload interface{}
}

func (e *EventTest) GetName() string {
	return e.Name
}

func (e *EventTest) GetPayload() interface{} {
	return e.Payload
}

func (e *EventTest) GetDateTime() time.Time {
	return time.Now()
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           EventTest
	event2          EventTest
	handler         EventHandlerTest
	handler2        EventHandlerTest
	handler3        EventHandlerTest
	eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.event = EventTest{Name: "teste-1", Payload: "teste 1"}
	suite.event2 = EventTest{Name: "teste-2", Payload: "teste 2"}
	suite.handler = EventHandlerTest{Id: "1"}
	suite.handler2 = EventHandlerTest{Id: "2"}
	suite.handler3 = EventHandlerTest{Id: "3"}
	suite.eventDispatcher = NewEventDispatcher()

}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	assert.Equal(suite.T(), &suite.handler, suite.eventDispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)

	suite.Equal(ErrorHandlerAlreadyRegister, err)

	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_Clear() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)

	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	suite.eventDispatcher.Clear()
	suite.Equal(0, len(suite.eventDispatcher.handlers))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_Has() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)

	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler))
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler2))
	assert.False(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler3))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (m *MockHandler) GetId() string {
	return "x"
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {

	eh := &MockHandler{}
	eh.On("Handle", &suite.event)

	suite.eventDispatcher.Register(suite.event.GetName(), eh)
	suite.eventDispatcher.Dispatch(&suite.event)

	eh.AssertExpectations(suite.T())

	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)

}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)

	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)

	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	err = suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
	suite.Equal(&suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][0])

	err = suite.eventDispatcher.Remove(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	err = suite.eventDispatcher.Remove(suite.event2.GetName(), &suite.handler3)
	assert.Equal(suite.T(), err, ErrorHandlerHasntRegister)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
