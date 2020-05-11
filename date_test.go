package fake

import (
	"testing"
	"time"
)

//func TestDate(t *testing.T) {
//t1,_ := time.Parse("2006-01-02 15:04:05", "2012-04-05 00:00:00")
//t2,_ := time.Parse("2006-01-02 15:04:05", "2012-04-05 00:03:00")
//for i := 0; i < testTryCount; i++ {
//	dt := DateRange(t1, t1)
//	if dt.Before(t1) || dt.After(t2) {
//		t.Errorf("DateRange() = %v, range = %v ~ %v", dt, t1, t2)
//	}
//}
//}

func TestDateRange(t *testing.T) {
	type args struct {
		t1 time.Time
		t2 time.Time
	}

	dt := time.Now()

	tests := []struct {
		name string
		args args
	}{
		{name: "equal", args: struct {
			t1 time.Time
			t2 time.Time
		}{t1: dt, t2: dt}},
		{name: "diff", args: struct {
			t1 time.Time
			t2 time.Time
		}{t1: dt, t2: dt.Add(3 * time.Second)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DateRange(tt.args.t1, tt.args.t2)
			if dt.Before(tt.args.t1) || dt.After(tt.args.t2) {
				t.Errorf("DateRange() = %v, range = %v ~ %v", got, tt.args.t1, tt.args.t2)
			}
		})
	}

}

func TestDate(t *testing.T) {
	type args struct {
		t time.Time
	}

	dt, _ := time.Parse("2006-01-02 15:04:05", "2012-04-05 00:00:00")
	tests := []struct {
		name string
		args args
	}{
		{name: "date", args: struct{ t time.Time }{t: dt}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Date(tt.args.t)
			if got.Before(tt.args.t.Add(commonJitter*(time.Duration(-1)))) || got.After(tt.args.t.Add(commonJitter)) {
				t.Errorf("Date() = %v, range: %v ~ %v", got, tt.args.t.Add(commonJitter*(time.Duration(-1))), tt.args.t.Add(commonJitter))
			}
		})
	}
}

func TestDateWithJitter(t *testing.T) {
	type args struct {
		t      time.Time
		jitter time.Duration
	}

	dt, _ := time.Parse("2006-01-02 15:04:05", "2012-04-05 00:00:00")

	tests := []struct {
		name string
		args args
	}{
		{name: "second", args: struct {
			t      time.Time
			jitter time.Duration
		}{t: dt, jitter: time.Second * 3}},
		{name: "equals", args: struct {
			t      time.Time
			jitter time.Duration
		}{t: dt, jitter: 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DateWithJitter(tt.args.t, tt.args.jitter)
			if got.Before(tt.args.t.Add(tt.args.jitter*(time.Duration(-1)))) || got.After(tt.args.t.Add(tt.args.jitter)) {
				t.Errorf("DateWithJitter() = %v, range: %v ~ %v", got, tt.args.t.Add(tt.args.jitter*-1), tt.args.t.Add(tt.args.jitter))
			}
		})
	}
}

func TestNow(t *testing.T) {
}
