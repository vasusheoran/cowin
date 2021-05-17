package cowin

import (
	err2 "cowin/err"
	"cowin/pkg/api/mocks"
	"cowin/pkg/constants"
	"os"
	"testing"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCowinService_GetCalendarByDistrict(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)

	testCases := []struct {
		name     string
		client   func() *mocks.HTTPClient
		notify   func() *mocks.Notify
		district int
		date     time.Time
		expected error
	}{
		{
			name: "",
			client: func() *mocks.HTTPClient {
				mockClient := &mocks.HTTPClient{}
				mockClient.On("Do", mock.Anything).Return(nil, err2.FailedToMakeHTTPRequest)
				return mockClient
			},
			notify: func() *mocks.Notify {
				mockNotify := &mocks.Notify{}
				mockNotify.On("Add", mock.Anything).Return()
				mockNotify.On("Notify", mock.Anything, mock.Anything).Return()
				return mockNotify
			},
			district: constants.SWD,
			date:     time.Now(),
			expected: err2.FailedToMakeHTTPRequest,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			cs := New(logger, testCase.notify(), testCase.client())
			err := cs.GetCalendarByDistrict(testCase.district, testCase.date)

			if testCase.expected != nil {
				assert.Equal(t, testCase.expected, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
