package xdgdir

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestNewApp(t *testing.T) {
	app := NewApp("test")
	if app.Name != "test" {
		t.Errorf("expected test, but got %s", app.Name)
	}
}

func TestAppConfigDir(t *testing.T) {
	app := NewApp("test")
	table := []struct {
		xdgHome     string
		home        string
		userProfile string
		expected    string
		err         bool
	}{
		{"a", "b", "c", path("a", "test"), false},
		{"", "b", "c", path("b", ".config", "test"), false},
		{"", "", "c", path("c", ".config", "test"), false},
		{"", "", "", "", true},
	}

	for _, tbl := range table {
		os.Setenv("XDG_CONFIG_HOME", tbl.xdgHome)
		os.Setenv("HOME", tbl.home)
		os.Setenv("USERPROFILE", tbl.userProfile)
		dir, err := app.ConfigDir()
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

func TestAppDataDir(t *testing.T) {
	app := NewApp("test")
	table := []struct {
		xdgHome     string
		home        string
		userProfile string
		expected    string
		err         bool
	}{
		{"a", "b", "c", path("a", "test"), false},
		{"", "b", "c", path("b", ".local", "share", "test"), false},
		{"", "", "c", path("c", ".local", "share", "test"), false},
		{"", "", "", "", true},
	}

	for _, tbl := range table {
		os.Setenv("XDG_DATA_HOME", tbl.xdgHome)
		os.Setenv("HOME", tbl.home)
		os.Setenv("USERPROFILE", tbl.userProfile)
		dir, err := app.DataDir()
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

func TestAppCacheDir(t *testing.T) {
	app := NewApp("test")
	table := []struct {
		xdgHome     string
		home        string
		userProfile string
		expected    string
		err         bool
	}{
		{"a", "b", "c", path("a", "test"), false},
		{"", "b", "c", path("b", ".cache", "test"), false},
		{"", "", "c", path("c", ".cache", "test"), false},
		{"", "", "", "", true},
	}

	for _, tbl := range table {
		os.Setenv("XDG_CACHE_HOME", tbl.xdgHome)
		os.Setenv("HOME", tbl.home)
		os.Setenv("USERPROFILE", tbl.userProfile)
		dir, err := app.CacheDir()
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

func TestAppRuntimeDir(t *testing.T) {
	app := NewApp("test")
	if app.RuntimeDir() == "" {
		t.Error("runtime dir should be not empty")
	}

	os.Setenv("XDG_RUNTIME_DIR", "/x")
	if dir := app.RuntimeDir(); dir != "/x/test" {
		t.Errorf("expected /x/test, but got %s", dir)
	}
}

func TestAppConfigFile(t *testing.T) {
	app := NewApp("test")
	name := "config.json"
	table := []struct {
		xdgHome     string
		home        string
		userProfile string
		expected    string
		err         bool
	}{
		{"a", "b", "c", path("a", "test", name), false},
		{"", "b", "c", path("b", ".config", "test", name), false},
		{"", "", "c", path("c", ".config", "test", name), false},
		{"", "", "", "", true},
	}

	for _, tbl := range table {
		os.Setenv("XDG_CONFIG_HOME", tbl.xdgHome)
		os.Setenv("HOME", tbl.home)
		os.Setenv("USERPROFILE", tbl.userProfile)
		dir, err := app.ConfigFile(name)
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

func TestAppConfigFileWithSubdirectory(t *testing.T) {
	app := NewApp("test")
	subdir := "sub"
	name := "config.json"
	table := []struct {
		xdgHome     string
		home        string
		userProfile string
		expected    string
		err         bool
	}{
		{"a", "b", "c", path("a", "test", subdir, name), false},
		{"", "b", "c", path("b", ".config", "test", subdir, name), false},
		{"", "", "c", path("c", ".config", "test", subdir, name), false},
		{"", "", "", "", true},
	}

	for _, tbl := range table {
		os.Setenv("XDG_CONFIG_HOME", tbl.xdgHome)
		os.Setenv("HOME", tbl.home)
		os.Setenv("USERPROFILE", tbl.userProfile)
		dir, err := app.ConfigFile(subdir, name)
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

func TestAppDataFile(t *testing.T) {
	app := NewApp("test")
	name := "data.json"
	table := []struct {
		xdgHome     string
		home        string
		userProfile string
		expected    string
		err         bool
	}{
		{"a", "b", "c", path("a", "test", name), false},
		{"", "b", "c", path("b", ".local", "share", "test", name), false},
		{"", "", "c", path("c", ".local", "share", "test", name), false},
		{"", "", "", "", true},
	}

	for _, tbl := range table {
		os.Setenv("XDG_DATA_HOME", tbl.xdgHome)
		os.Setenv("HOME", tbl.home)
		os.Setenv("USERPROFILE", tbl.userProfile)
		dir, err := app.DataFile(name)
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

func TestAppDataFileWithSubdirectory(t *testing.T) {
	app := NewApp("test")
	subdir := "sub"
	name := "data.json"
	table := []struct {
		xdgHome     string
		home        string
		userProfile string
		expected    string
		err         bool
	}{
		{"a", "b", "c", path("a", "test", subdir, name), false},
		{"", "b", "c", path("b", ".local", "share", "test", subdir, name), false},
		{"", "", "c", path("c", ".local", "share", "test", subdir, name), false},
		{"", "", "", "", true},
	}

	for _, tbl := range table {
		os.Setenv("XDG_DATA_HOME", tbl.xdgHome)
		os.Setenv("HOME", tbl.home)
		os.Setenv("USERPROFILE", tbl.userProfile)
		dir, err := app.DataFile("sub", "data.json")
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

func TestAppCacheFile(t *testing.T) {
	app := NewApp("test")
	name := "cache.json"
	table := []struct {
		xdgHome     string
		home        string
		userProfile string
		expected    string
		err         bool
	}{
		{"a", "b", "c", path("a", "test", name), false},
		{"", "b", "c", path("b", ".cache", "test", name), false},
		{"", "", "c", path("c", ".cache", "test", name), false},
		{"", "", "", "", true},
	}

	for _, tbl := range table {
		os.Setenv("XDG_CACHE_HOME", tbl.xdgHome)
		os.Setenv("HOME", tbl.home)
		os.Setenv("USERPROFILE", tbl.userProfile)
		dir, err := app.CacheFile(name)
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

func TestAppCacheFileWithSubdirectory(t *testing.T) {
	app := NewApp("test")
	subdir := "sub"
	name := "cache.json"
	table := []struct {
		xdgHome     string
		home        string
		userProfile string
		expected    string
		err         bool
	}{
		{"a", "b", "c", path("a", "test", subdir, name), false},
		{"", "b", "c", path("b", ".cache", "test", subdir, name), false},
		{"", "", "c", path("c", ".cache", "test", subdir, name), false},
		{"", "", "", "", true},
	}

	for _, tbl := range table {
		os.Setenv("XDG_CACHE_HOME", tbl.xdgHome)
		os.Setenv("HOME", tbl.home)
		os.Setenv("USERPROFILE", tbl.userProfile)
		dir, err := app.CacheFile(subdir, name)
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

func TestAppRuntimeFile(t *testing.T) {
	app := NewApp("test")
	name := "runtime.pid"
	if app.RuntimeFile(name) == "" {
		t.Error("runtime dir should be not empty")
	}

	os.Setenv("XDG_RUNTIME_DIR", "a")
	expected := path("a", "test", "runtime.pid")
	if dir := app.RuntimeFile(name); dir != expected {
		t.Errorf("expected %s but got %s", expected, dir)
	}
}

func TestAppRuntimeFileWithSubdirectory(t *testing.T) {
	app := NewApp("test")
	subdir := "sub"
	name := "runtime.pid"
	if app.RuntimeFile(name) == "" {
		t.Error("runtime dir should be not empty")
	}

	os.Setenv("XDG_RUNTIME_DIR", "a")
	expected := path("a", "test", subdir, name)
	if dir := app.RuntimeFile(subdir, name); dir != expected {
		t.Errorf("expected %s, but got %s", expected, dir)
	}
}

func TestAppFindConfigFile(t *testing.T) {
	app := NewApp("test")
	os.Setenv("XDG_CONFIG_HOME", path("testdata", "a"))
	os.Setenv("XDG_CONFIG_DIRS", join(path("testdata", "b"), path("testdata", "c")))
	table := []struct {
		name    string
		content string
		err     bool
	}{
		{"zzz.txt", "testdata/z/test/zzz.txt", true},
		{"aaa.txt", "testdata/a/test/aaa.txt", false},
		{"bbb.txt", "testdata/b/test/bbb.txt", false},
		{"ccc.txt", "testdata/c/test/ccc.txt", false},
	}
	for _, test := range table {
		f, err := app.FindConfigFile(test.name)
		if test.err {
			if err == nil {
				t.Error("should raise error, but not raised")
			}
			continue
		}

		if err != nil {
			t.Error(err)
			continue
		}
		s, err := openFile(f)
		if err != nil {
			t.Error(err)
			continue
		}
		if s != test.content {
			t.Error("find invalid file")
		}
	}
}

func TestAppFindConfigFileWithSubdirecotry(t *testing.T) {
	app := NewApp("test")
	os.Setenv("XDG_CONFIG_HOME", path("testdata", "a"))
	os.Setenv("XDG_CONFIG_DIRS", join(path("testdata", "b"), path("testdata", "c")))

	f, err := app.FindConfigFile("d", "ddd.txt")
	if err != nil {
		t.Error(err)
	}
	s, err := openFile(f)
	if err != nil {
		t.Fatal(err)
	}
	if s != "testdata/c/test/d/ddd.txt" {
		t.Error("find invalid file")
	}
}

func TestAppFindDataFile(t *testing.T) {
	app := NewApp("test")
	os.Setenv("XDG_DATA_HOME", path("testdata", "a"))
	os.Setenv("XDG_DATA_DIRS", join(path("testdata", "b"), path("testdata", "c")))
	table := []struct {
		name    string
		content string
		err     bool
	}{
		{"zzz.txt", "testdata/z/test/zzz.txt", true},
		{"aaa.txt", "testdata/a/test/aaa.txt", false},
		{"bbb.txt", "testdata/b/test/bbb.txt", false},
		{"ccc.txt", "testdata/c/test/ccc.txt", false},
	}
	for _, test := range table {
		f, err := app.FindDataFile(test.name)
		if test.err {
			if err == nil {
				t.Error("should raise error, but not raised")
			}
			continue
		}

		if err != nil {
			t.Error(err)
			continue
		}
		s, err := openFile(f)
		if err != nil {
			t.Error(err)
			continue
		}
		if s != test.content {
			t.Error("find invalid file")
		}
	}
}

func TestAppFindDataFileWithSubdirectory(t *testing.T) {
	app := NewApp("test")
	os.Setenv("XDG_DATA_HOME", path("testdata", "a"))
	os.Setenv("XDG_DATA_DIRS", join(path("testdata", "b"), path("testdata", "c")))

	f, err := app.FindDataFile("d", "ddd.txt")
	if err != nil {
		t.Error(err)
	}
	s, err := openFile(f)
	if err != nil {
		t.Fatal(err)
	}
	if s != "testdata/c/test/d/ddd.txt" {
		t.Error("find invalid file")
	}
}

func openFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	p, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(p)), nil
}

func join(elm ...string) string {
	return strings.Join(elm, string(os.PathListSeparator))
}
