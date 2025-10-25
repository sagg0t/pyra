package nutrition

import (
	"testing"
)

func Test_Measurement(t *testing.T) {
	t.Run("NewMeasurement", func(t *testing.T) {
		testCases := []struct {
			in  float64
			out int32
		}{
			{in: 0, out: 0},
			{in: 0.0001, out: 0},
			{in: 0.001, out: 0},
			{in: 0.01, out: 1},
			{in: 0.1, out: 10},
			{in: 1, out: 100},
			{in: 1.2, out: 120},
			{in: 1.23, out: 123},
			{in: 1.234, out: 123},
		}

		for _, tc := range testCases {
			m := NewMeasurement(tc.in)

			if m != Measurement(tc.out) {
				t.Errorf("expected NewMeasurement(%f) to equal Measurement(%d), get %d", tc.in, tc.out, m)
			}
		}
	})

	t.Run("ParseMeasurement", func(t *testing.T) {
		t.Run("valid input", func(t *testing.T) {
			testCases := []struct {
				in  string
				out int32
			}{
				{in: "", out: 0},
				{in: "0.0001", out: 0},
				{in: "0.01", out: 1},
				{in: "0.1", out: 10},
				{in: "1", out: 100},
				{in: "1.2", out: 120},
				{in: "1.23", out: 123},
				{in: "1.234", out: 123},
				// Edge cases
				{in: "8631.71", out: 863171},
			}

			for _, tc := range testCases {
				m, err := ParseMeasurement(tc.in)
				if err != nil {
					t.Errorf("expected ParseMeasurement(%q) to return no errors, got %v", tc.in, err)
				}

				if m != Measurement(tc.out) {
					t.Errorf("expected ParseMeasurement(%q) to equal Measurement(%d), get %d", tc.in, tc.out, m)
				}
			}
		})

		t.Run("invalid input", func(t *testing.T) {
			testCases := []string{
				" ",
				"alsdkfj",
				// Parsable, but not allowed values
				"NaN",
				"nan",
				"Inf",
				"+Inf",
				"-Inf",
			}

			for _, input := range testCases {
				_, err := ParseMeasurement(input)
				if err == nil {
					t.Errorf("expected ParseMeasurement(%q) to return error", input)
				} else if _, ok := err.(*MeasurementError); !ok {
					t.Errorf("expected ParseMeasurement(%q) to return %T, got: (%T) %v",
						input, &MeasurementError{}, err, err)
				}
			}
		})
	})

	t.Run("Measurement.String", func(t *testing.T) {
		t.Run("with 2 decimal digits", func(t *testing.T) {
			m := Measurement(479)
			mS := "4.79"

			if m.String() != mS {
				t.Errorf("expected Measurement(%d).String() to be %q, got %q", m, mS, m.String())
			}
		})
	})

	t.Run("Measurement.Float", func(t *testing.T) {
		testCases := []struct {
			in  int32
			out float64
		}{
			{in: 1, out: 0.01},
			{in: 12, out: 0.12},
			{in: 123, out: 1.23},
			{in: 1234, out: 12.34},
			// Edge cases
			{in: 863171, out: 8631.71},
		}

		for _, tc := range testCases {
			m := Measurement(tc.in)
			mFloat := m.Float()
			if mFloat != tc.out {
				t.Errorf("expected Measurement(%d).Float() to equal %f, get %f", tc.in, tc.out, mFloat)
			}
		}
	})

	t.Run("Measurement.Scale", func(t *testing.T) {
		testCases := []struct {
			in     Measurement
			out    Measurement
			factor float64
		}{
			{in: 1, out: 3, factor: 3},
		}

		for _, tc := range testCases {
			scaled := tc.in.Scale(tc.factor)
			if scaled != tc.out {
				t.Errorf("expected Measurement(%d).Scale(%f) to equal %d, get %d",
					tc.in, tc.factor, tc.out, scaled)
			}
		}
	})
}
