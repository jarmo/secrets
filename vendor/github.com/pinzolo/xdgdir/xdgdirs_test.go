package xdgdir

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConfigDir(t *testing.T) {
	table := []struct {
		xdgHome     string
		home        string
		userProfile string
		expected    string
		err         bool
	}{
		{"x", "y", "z", "x", false},
		{"", "y", "z", path("y", ".config"), false},
		{"", "", "z", path("z", ".config"), false},
		{"", "", "", "", true},
	}

	for _, tbl := range table {
		os.Setenv("XDG_CONFIG_HOME", tbl.xdgHome)
		os.Setenv("HOME", tbl.home)
		os.Setenv("USERPROFILE", tbl.userProfile)
		dir, err := ConfigDir()
		if tbl.err {
			if err == nil {
				t.Error("should raise error, but not raised")
			}
		} else {
			if err != nil {
				t.Error(err)
			}
			if dir != tbl.expected {
				t.Errorf("expected %s, but got %s", tbl.expected, dir)
			}
		}
	}
}

func TestDataDir(t *testing.T) {
	table := []struct {
		xdgHome     string
		home        string
		userProfile string
		expected    string
		err         bool
	}{
		{"x", "y", "z", "x", false},
		{"", "y", "z", path("y", ".local", "share"), false},
		{"", "", "z", path("z", ".local", "share"), false},
		{"", "", "", "", true},
	}

	for _, tbl := range table {
		os.Setenv("XDG_DATA_HOME", tbl.xdgHome)
		os.Setenv("HOME", tbl.home)
		os.Setenv("USERPROFILE", tbl.userProfile)
		dir, err := DataDir()
		if tbl.err {
			if err == nil {
				t.Error("should raise error, but not raised")
			}
		} else {
			if err != nil {
				t.Error(err)
			}
			if dir != tbl.expected {
				t.Errorf("expected %s, but got %s", tbl.expected, dir)
			}
		}
	}
}

func TestCacheDir(t *testing.T) {
	table := []struct {
		xdgHome     string
		home        string
		userProfile string
		expected    string
		err         bool
	}{
		{"x", "y", "z", "x", false},
		{"", "y", "z", path("y", ".cache"), false},
		{"", "", "z", path("z", ".cache"), false},
		{"", "", "", "", true},
	}

	for _, tbl := range table {
		os.Setenv("XDG_CACHE_HOME", tbl.xdgHome)
		os.Setenv("HOME", tbl.home)
		os.Setenv("USERPROFILE", tbl.userProfile)
		dir, err := CacheDir()
		if tbl.err {
			if err == nil {
				t.Error("should raise error, but not raised")
			}
		} else {
			if err != nil {
				t.Error(err)
			}
			if dir != tbl.expected {
				t.Errorf("expected %s, but got %s", tbl.expected, dir)
			}
		}
	}
}

func TestRuntimeDir(t *testing.T) {
	if RuntimeDir() == "" {
		t.Error("runtime dir should be not empty")
	}

	os.Setenv("XDG_RUNTIME_DIR", "x")
	if dir := RuntimeDir(); dir != "x" {
		t.Errorf("expected x, but got %s", dir)
	}
}

func path(elm ...string) string {
	return filepath.Join(elm...)
}
