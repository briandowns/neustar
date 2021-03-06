package neustar

import (
	"os"
	"reflect"
	"testing"
)

func setUp() *Neustar {
	return NewNeustar(os.Getenv("NEUSTAR_KEY"), os.Getenv("NEUSTAR_SECRET"))
}

// TestNewNeustar
func TestNewNeustar(t *testing.T) {
	n := NewNeustar(os.Getenv("NEUSTAR_KEY"), os.Getenv("NEUSTAR_SECRET"))

	if reflect.TypeOf(n).String() != "*neustar.Neustar" {
		t.Error("Incorrect data type pointer returned from NewNeustar function")
	}
}

// TestNewMonitor
func TestNewMonitor(t *testing.T) {
	t.Parallel()

	m := NewMonitor(setUp())

	if reflect.TypeOf(m).String() != "*neustar.Monitoring" {
		t.Error("Incorrect data type pointer returned from NewMonitor function")
	}
}

// TestCreate
func TestCreate(t *testing.T) {
	t.Parallel()

}

// TestList
func TestList(t *testing.T) {
	t.Parallel()

	monitor := NewMonitor(setUp())

	_, err := monitor.List()
	if err != nil {
		t.Error(err)
	}
}

// TestGet
func TestGet(t *testing.T) {
	t.Parallel()

	monitor := NewMonitor(setUp())

	monitorIDs := []string{"5a1c2ab6588c11e592489848e1660ab3", "21e85880588d11e5b2309848e1660ab3"}

	for _, monitorID := range monitorIDs {
		_, err := monitor.Get(monitorID)
		if err != nil {
			t.Error(err)
		}
	}
}

// TestUpdate
func TestUpdate(t *testing.T) {
	t.Parallel()
}

// TestDelete
func TestDelete(t *testing.T) {
	t.Parallel()
}

// TestRawSampleData
func TestRawSampleData(t *testing.T) {
	t.Parallel()
}

// TestSamples
func TestSamples(t *testing.T) {
	t.Parallel()
}

// TestAggregateSampleData
func TestAggregateSampleData(t *testing.T) {
	t.Parallel()
}

// TestSummary
func TestSummary(t *testing.T) {
	t.Parallel()
}

// TestLocations
func TestLocations(t *testing.T) {
	t.Parallel()
}

// TestValidMonitorType
func TestValidMonitorType(t *testing.T) {
	t.Parallel()
}

// TestValidBrowserType
func TestValidBrowserType(t *testing.T) {
	t.Parallel()

	testTypes := []string{"FF", "CHROME", "IE", "Opera"}

	for _, bt := range BrowserTypes {
		for _, tt := range testTypes {
			if tt == bt {
				continue
			}
		}
	}
}

// TestValidUpdateInterval
func TestValidUpdateInterval(t *testing.T) {
	t.Parallel()

	testIntervals := []int{1, 2, 3, 4, 5, 10, 15, 20, 30, 60, 77}

	for _, bt := range UpdateIntervals {
		for _, ti := range testIntervals {
			if ti == bt {
				continue
			}
		}
	}
}

// TestValidAggregateSampleDataFrequency
func TestValidAggregateSampleDataFrequency(t *testing.T) {
	t.Parallel()
}

// TestValidAggregateSampleGroupBy
func TestValidAggregateSampleGroupBy(t *testing.T) {
	t.Parallel()
}
