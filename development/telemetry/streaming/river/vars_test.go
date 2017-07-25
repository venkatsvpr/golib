package river

import (
	"encoding/json"
	"expvar"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/johnsiilver/boutique"
	"github.com/johnsiilver/golib/development/telemetry/streaming/river/state/data"
	"github.com/kylelemons/godebug/pretty"
)

// TestRiverVar makes sure that all our types conform to river.Var.
func TestRiverVar(t *testing.T) {
	_ = []Var{
		Int{}, Float{}, String{}, &Map{},
	}
}

func TestInt(t *testing.T) {
	t.Parallel()

	x := newInt("name", 1)

	// Subscribe to the changes.
	sub, cancel := x.Subscribe()
	defer cancel()
	var final atomic.Value // data.VarState

	// Have something listening to the changes.
	go func() {
		for i := range sub {
			final.Store(i.State.Data.(data.VarState))
			time.Sleep(1 * time.Second) // Pause to make sure we miss some changes.
		}
	}()

	// Add 1000 to our number, but do it with 1000 goroutines.
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			x.Add(1)
		}()
	}
	wg.Wait() // Wait for this to finish.

	// Make sure we get a final update of 1001, but wait no more than 10 seconds.
	start := time.Now()
	for {
		f := final.Load().(data.VarState)
		if time.Now().Sub(start) > 10*time.Second {
			t.Fatalf("TestInt: final subscribe int: got %v, want %v", f.Int, 1001)
		}
		if f.Int == 1001 {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	if x.String() != "1001" {
		t.Errorf("TestInt: Int.String(): got %v, want %v", x.String(), "1001")
	}

	if x.Value() != 1001 {
		t.Errorf("TestInt: Int.Value(): got %v, want %v", x.Value(), 1001)
	}

	if diff := pretty.Compare(data.VarState{Name: "name", Type: data.IntType, Int: 1001}, x.VarState()); diff != "" {
		t.Errorf("TestInt: Int.VarState(): -want/+got:\n%s", diff)
	}

	x.Set(10)
	if x.Value() != 10 {
		t.Errorf("TestInt: Int.Set(10): got %v, want %v", x.Value(), 10)
	}
}

func TestFloat(t *testing.T) {
	t.Parallel()

	x := newFloat("name", 1)

	// Subscribe to the changes.
	sub, cancel := x.Subscribe()
	defer cancel()
	var final atomic.Value // data.VarState

	// Have something listening to the changes.
	go func() {
		for i := range sub {
			final.Store(i.State.Data.(data.VarState))
			time.Sleep(1 * time.Second) // Pause to make sure we miss some changes.
		}
	}()

	// Add 1000 to our number, but do it with 1000 goroutines.
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			x.Add(1)
		}()
	}
	wg.Wait() // Wait for this to finish.

	// Make sure we get a final update of 1001, but wait no more than 10 seconds.
	start := time.Now()
	for {
		f := final.Load().(data.VarState)
		if time.Now().Sub(start) > 10*time.Second {
			t.Fatalf("TestFloat: final subscribe float: got %v, want %v", f.Float, 1001)
		}
		if f.Float == 1001 {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	if x.String() != "1001" {
		t.Errorf("TestFloat: Flaot.String(): got %v, want %v", x.String(), "1001")
	}

	if x.Value() != 1001 {
		t.Errorf("TestFloat: Float.Value(): got %v, want %v", x.Value(), 1001)
	}

	if diff := pretty.Compare(data.VarState{Name: "name", Type: data.FloatType, Float: 1001}, x.VarState()); diff != "" {
		t.Errorf("TestFloat: Float.VarState(): -want/+got:\n%s", diff)
	}

	x.Set(10)
	if x.Value() != 10 {
		t.Errorf("TestFloat: Float.Set(10): got %v, want %v", x.Value(), 10)
	}
}

func TestString(t *testing.T) {
	t.Parallel()

	x := newString("name", "b")

	// Subscribe to the changes.
	sub, cancel := x.Subscribe()
	defer cancel()
	var final atomic.Value // data.VarState

	// Have something listening to the changes.
	go func() {
		for i := range sub {
			final.Store(i.State.Data.(data.VarState))
			time.Sleep(1 * time.Second) // Pause to make sure we miss some changes.
		}
	}()

	str := []string{"b"}
	for i := 0; i < 1000; i++ {
		str = append(str, "a")
		x.Set(strings.Join(str, ""))
	}
	finalStr := strings.Join(str, "")

	start := time.Now()
	for {
		f := final.Load().(data.VarState)
		if time.Now().Sub(start) > 10*time.Second {
			t.Fatalf("TestString: final subscribe string: got %v, want %v", f.String, finalStr)
		}
		if f.String == finalStr {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	b, _ := json.Marshal(finalStr)
	if x.String() != string(b) {
		t.Errorf("TestString: String.String(): got %v, want %v", x.String(), string(b))
	}

	if x.Value() != finalStr {
		t.Errorf("TestString: String.Value(): got %v, want %v", x.Value(), finalStr)
	}

	if diff := pretty.Compare(data.VarState{Name: "name", Type: data.StringType, String: finalStr}, x.VarState()); diff != "" {
		t.Errorf("TestString: String.VarState(): -want/+got:\n%s", diff)
	}
}

func TestMap(t *testing.T) {
	t.Parallel()

	x := newMap("name", map[string]expvar.Var{})

	// Subscribe to the changes.
	sub, cancel := x.Subscribe()
	defer cancel()
	var final atomic.Value // data.VarState

	// Have something listening to the changes.
	go func() {
		for i := range sub {
			final.Store(i.State.Data.(data.VarState))
			time.Sleep(1 * time.Second) // Pause to make sure we miss some changes.
		}
	}()

	// Test Set(), Add(), AddFloat().
	wg := sync.WaitGroup{}
	x.Set("int", newInt("name", 0))
	for i := 0; i < 1000; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			x.Add("int", 1)
		}()
		go func() {
			defer wg.Done()
			x.AddFloat("float", 1)
		}()
	}
	wg.Wait()

	if x.Get("int").String() != "1000" {
		t.Errorf("TestMap: key 'int': got %s, want %s", x.Get("int"), "1000")
	}

	if x.Get("float").String() != "1000" {
		t.Errorf("TestMap: key 'float': got %v, want %v", x.Get("float"), "1000")
	}

	// Test that updating a var in a map gives us an update.
	i := newInt("great", 200)
	x.Set("great", i)
	up, cancel := x.Subscribe()
	defer cancel()
	var sig boutique.Signal
	wg.Add(1)
	ready := make(chan struct{})
	go func() {
		defer wg.Done()
		// Clear out any update we have waiting.
		select {
		case <-up:
		default:
		}
		close(ready)
		sig = <-up
	}()

	<-ready

	i.Add(1)
	wg.Wait()

	got := sig.State.Data.(data.VarState).Map["great"].String()
	if got != "201" {
		t.Errorf("TestMap: testing Map.Set() subscriptions: got %v, want %v", got, "201")
	}

	const wantStr = "{\"int\": 1000, \"float\": 1000, \"great\": 201}"
	gotStr := x.String()
	if gotStr != wantStr {
		t.Errorf("TestMap: testing Map.Strin(): got %v, want %v", gotStr, wantStr)
	}
}

type ByKey []expvar.KeyValue

func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }
