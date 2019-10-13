package backend

import (
	"reflect"
	"testing"
)

type FakeEnv struct{}

type FakeBackend struct {
	name string
	env  Environment
	err  error
}

func (fbk FakeBackend) Name() string {
	return fbk.name
}

func (fbk FakeBackend) Create() (Environment, error) {
	return fbk.env, fbk.err
}
func TestGet(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    Backend
		wantErr bool
	}{
		{
			name: "backend unregistered",
			args: args{
				name: "bk01",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "backend unregistered",
			args: args{
				name: "fbk1",
			},
			want:    FakeBackend{name: "fbk1"},
			wantErr: false,
		},
	}

	backendsMu.Lock()
	backends = map[string]Backend{
		"fbk1": FakeBackend{name: "fbk1"},
	}
	backendsMu.Unlock()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = '%v', wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	type args struct {
		backend Backend
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{
			name: "register a proper backend",
			args: args{
				backend: FakeBackend{name: "fbk1"},
			},
			wantPanic: false,
		},
		{
			name: `register backend with name ""`,
			args: args{
				backend: FakeBackend{name: ""},
			},
			wantPanic: true,
		},

		{
			name: "register <nil>",
			args: args{
				backend: nil,
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); (err != nil) != tt.wantPanic {
					t.Errorf("Register(), did panic=%v", tt.wantPanic)
				}
			}()

			Register(tt.args.backend)
		})
	}
}

func TestList(t *testing.T) {
	tests := []struct {
		name string
		bks  map[string]Backend
		want []string
	}{
		{
			name: "list with no backend",
			bks:  map[string]Backend{},
			want: []string{},
		},
		{
			name: "list one backend",
			bks: map[string]Backend{
				"fbk1": FakeBackend{name: "fbk1"},
			},
			want: []string{
				"fbk1",
			},
		},
		{
			name: "list many backends",
			bks: map[string]Backend{
				"fbk1":    FakeBackend{name: "fbk1"},
				"xlib":    FakeBackend{name: "xlib"},
				"fake123": FakeBackend{name: "fake123"},
				"back":    FakeBackend{name: "back"},
				"x123":    FakeBackend{name: "x123"},
			},
			want: []string{
				"back",
				"fake123",
				"fbk1",
				"x123",
				"xlib",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			backendsMu.Lock()
			backends = tt.bks
			backendsMu.Unlock()

			if got := List(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() = %v, want %v", got, tt.want)
			}
		})
	}
}
