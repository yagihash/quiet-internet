package config

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	EnvToken = "QI_TOKEN"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name        string
		input       map[string]string
		want        *Config
		expectError bool
	}{
		{
			name: "Success",
			input: map[string]string{
				EnvToken: "sk-xxxxxxxxxxxxxxxxxxxxx",
			},
			want: &Config{
				Token: "sk-xxxxxxxxxxxxxxxxxxxxx",
			},
			expectError: false,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			if err := SetEnv(c.input); err != nil {
				t.Fatal(err)
			}
			defer func() {
				if err := UnsetEnv(c.input); err != nil {
					t.Fatal(err)
				}
			}()

			got, err := Load()

			if c.expectError && err == nil {
				t.Error("expected error but got nil")
				return
			}

			if diff := cmp.Diff(got, c.want); diff != "" {
				t.Errorf("got an unexpected diff:\n%s", diff)
			}
		})
	}
}

func SetEnv(env map[string]string) error {
	for k, v := range env {
		err := os.Setenv(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func UnsetEnv(env map[string]string) error {
	for k, _ := range env {
		err := os.Unsetenv(k)
		if err != nil {
			return err
		}
	}

	return nil
}
