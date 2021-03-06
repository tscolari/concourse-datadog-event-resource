// This file was generated by counterfeiter
package resourcefakes

import (
	"sync"

	"github.com/tscolari/concourse-datadog-event-resource/resource"
	go_datadog_api "github.com/zorkian/go-datadog-api"
)

type FakeDatadogClient struct {
	GetEventsStub        func(start, end int, priority, sources, tags string) ([]go_datadog_api.Event, error)
	getEventsMutex       sync.RWMutex
	getEventsArgsForCall []struct {
		start    int
		end      int
		priority string
		sources  string
		tags     string
	}
	getEventsReturns struct {
		result1 []go_datadog_api.Event
		result2 error
	}
	GetEventStub        func(id int) (*go_datadog_api.Event, error)
	getEventMutex       sync.RWMutex
	getEventArgsForCall []struct {
		id int
	}
	getEventReturns struct {
		result1 *go_datadog_api.Event
		result2 error
	}
	PostEventStub        func(event *go_datadog_api.Event) (*go_datadog_api.Event, error)
	postEventMutex       sync.RWMutex
	postEventArgsForCall []struct {
		event *go_datadog_api.Event
	}
	postEventReturns struct {
		result1 *go_datadog_api.Event
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDatadogClient) GetEvents(start int, end int, priority string, sources string, tags string) ([]go_datadog_api.Event, error) {
	fake.getEventsMutex.Lock()
	fake.getEventsArgsForCall = append(fake.getEventsArgsForCall, struct {
		start    int
		end      int
		priority string
		sources  string
		tags     string
	}{start, end, priority, sources, tags})
	fake.recordInvocation("GetEvents", []interface{}{start, end, priority, sources, tags})
	fake.getEventsMutex.Unlock()
	if fake.GetEventsStub != nil {
		return fake.GetEventsStub(start, end, priority, sources, tags)
	} else {
		return fake.getEventsReturns.result1, fake.getEventsReturns.result2
	}
}

func (fake *FakeDatadogClient) GetEventsCallCount() int {
	fake.getEventsMutex.RLock()
	defer fake.getEventsMutex.RUnlock()
	return len(fake.getEventsArgsForCall)
}

func (fake *FakeDatadogClient) GetEventsArgsForCall(i int) (int, int, string, string, string) {
	fake.getEventsMutex.RLock()
	defer fake.getEventsMutex.RUnlock()
	return fake.getEventsArgsForCall[i].start, fake.getEventsArgsForCall[i].end, fake.getEventsArgsForCall[i].priority, fake.getEventsArgsForCall[i].sources, fake.getEventsArgsForCall[i].tags
}

func (fake *FakeDatadogClient) GetEventsReturns(result1 []go_datadog_api.Event, result2 error) {
	fake.GetEventsStub = nil
	fake.getEventsReturns = struct {
		result1 []go_datadog_api.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeDatadogClient) GetEvent(id int) (*go_datadog_api.Event, error) {
	fake.getEventMutex.Lock()
	fake.getEventArgsForCall = append(fake.getEventArgsForCall, struct {
		id int
	}{id})
	fake.recordInvocation("GetEvent", []interface{}{id})
	fake.getEventMutex.Unlock()
	if fake.GetEventStub != nil {
		return fake.GetEventStub(id)
	} else {
		return fake.getEventReturns.result1, fake.getEventReturns.result2
	}
}

func (fake *FakeDatadogClient) GetEventCallCount() int {
	fake.getEventMutex.RLock()
	defer fake.getEventMutex.RUnlock()
	return len(fake.getEventArgsForCall)
}

func (fake *FakeDatadogClient) GetEventArgsForCall(i int) int {
	fake.getEventMutex.RLock()
	defer fake.getEventMutex.RUnlock()
	return fake.getEventArgsForCall[i].id
}

func (fake *FakeDatadogClient) GetEventReturns(result1 *go_datadog_api.Event, result2 error) {
	fake.GetEventStub = nil
	fake.getEventReturns = struct {
		result1 *go_datadog_api.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeDatadogClient) PostEvent(event *go_datadog_api.Event) (*go_datadog_api.Event, error) {
	fake.postEventMutex.Lock()
	fake.postEventArgsForCall = append(fake.postEventArgsForCall, struct {
		event *go_datadog_api.Event
	}{event})
	fake.recordInvocation("PostEvent", []interface{}{event})
	fake.postEventMutex.Unlock()
	if fake.PostEventStub != nil {
		return fake.PostEventStub(event)
	} else {
		return fake.postEventReturns.result1, fake.postEventReturns.result2
	}
}

func (fake *FakeDatadogClient) PostEventCallCount() int {
	fake.postEventMutex.RLock()
	defer fake.postEventMutex.RUnlock()
	return len(fake.postEventArgsForCall)
}

func (fake *FakeDatadogClient) PostEventArgsForCall(i int) *go_datadog_api.Event {
	fake.postEventMutex.RLock()
	defer fake.postEventMutex.RUnlock()
	return fake.postEventArgsForCall[i].event
}

func (fake *FakeDatadogClient) PostEventReturns(result1 *go_datadog_api.Event, result2 error) {
	fake.PostEventStub = nil
	fake.postEventReturns = struct {
		result1 *go_datadog_api.Event
		result2 error
	}{result1, result2}
}

func (fake *FakeDatadogClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getEventsMutex.RLock()
	defer fake.getEventsMutex.RUnlock()
	fake.getEventMutex.RLock()
	defer fake.getEventMutex.RUnlock()
	fake.postEventMutex.RLock()
	defer fake.postEventMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeDatadogClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ resource.DatadogClient = new(FakeDatadogClient)
