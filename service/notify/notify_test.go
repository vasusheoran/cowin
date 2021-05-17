package notify

import (
	"cowin/pkg/constants"
	"cowin/pkg/contracts"
	"os"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
)

func TestNotifyService_Load(t *testing.T) {
	logger := log.NewLogfmtLogger(os.Stderr)

	testCases := []struct {
		name     string
		add      contracts.Filter
		expected error
	}{
		{
			name: "Success",
			add: contracts.Filter{
				ID:       "testid",
				District: 150,
				DoseType: 1,
				Age:      46,
				Vaccine:  constants.COVAX,
				Phone:    "",
				Email:    "test@gmail.com",
			},
			expected: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ns := New(logger)
			_, err := ns.Load()

			if testCase.expected != nil {
				assert.Equal(t, testCase.expected, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
