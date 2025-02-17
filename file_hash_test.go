package checksum_test

import (
	"checksum"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func prepareFile() (string, error) {
	file, err := ioutil.TempFile("", "prefix")
	if err != nil {
		return "", err
	}
	if err := ioutil.WriteFile(file.Name(), []byte("some data"), 0600); err != nil {
		return "", err
	}
	return file.Name(), nil
}

func TestSHA256sumFile(t *testing.T) {
	file, err := prepareFile()
	if err != nil {
		t.Logf("could not create test file: %s", err)
		t.FailNow()
	}
	defer func() {
		err := os.Remove(file)
		if err != nil {
			t.Logf("could not remove test file: %s", err)
		}
	}()

	if sha256sum, err := checksum.SHA256sum(file); err != nil || sha256sum != "1307990e6ba5ca145eb35e99182a9bec46531bc54ddf656a602c780fa0240dee" {
		t.Error("SHA256sum(file) failed", sha256sum, err)
	}
}

func TestMd5sumFile(t *testing.T) {
	file, err := prepareFile()
	if err != nil {
		t.Logf("could not create test file: %s", err)
		t.FailNow()
	}
	defer func() {
		err := os.Remove(file)
		if err != nil {
			t.Logf("could not remove test file: %s", err)
		}
	}()

	if md5sum, err := checksum.MD5sum(file); err != nil || md5sum != "1e50210a0202497fb79bc38b6ade6c34" {
		t.Error("Md5sum(file) failed", md5sum, err)
	}
}

func TestMd5sumDir(t *testing.T) {
	homeDirectory, err := homedir.Dir()
	if err != nil {
		t.Logf("could not get home directory: %s", err)
		t.FailNow()
	}
	file := path.Join(homeDirectory, "Downloads")

	if md5sum, err := checksum.MD5sum(file); err != nil || md5sum != "" {
		t.Error("Md5sum(dir) failed", md5sum, err)
	}
}
