// Code generated for package migrations by go-bindata DO NOT EDIT. (@generated)
// sources:
// migrations/1688189195_init.down.sql
// migrations/1688189195_init.up.sql
package migrations

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _migrations1688189195_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x2f\xcd\x8c\x4f\xce\xcf\x4b\xcb\x4c\xb7\x06\x04\x00\x00\xff\xff\x49\xa7\x32\xcb\x20\x00\x00\x00")

func migrations1688189195_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations1688189195_initDownSql,
		"migrations/1688189195_init.down.sql",
	)
}

func migrations1688189195_initDownSql() (*asset, error) {
	bytes, err := migrations1688189195_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/1688189195_init.down.sql", size: 32, mode: os.FileMode(436), modTime: time.Unix(1665214714, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations1688189195_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xd2\xb1\x4e\xc3\x30\x10\x06\xe0\xdd\x4f\xf1\x8f\x89\xc4\x06\x62\xe9\xe4\x96\x2b\xb5\x70\xec\xca\xb9\x50\xca\x52\x19\xe2\x56\x91\x20\x45\x49\x2a\x78\x7c\xe4\x84\x2e\x61\x20\xa2\x37\x9e\x75\x9f\xee\x97\x6f\xe1\x48\x32\x81\xe5\x5c\x13\xd4\x12\xc6\x32\xe8\x49\xe5\x9c\xe3\xd4\x86\xa6\x15\x89\x00\x80\xaa\xc4\xb9\xe6\xea\x5e\x19\xc6\x84\x8a\x96\x29\xb4\xc6\xda\xa9\x4c\xba\x2d\x1e\x68\x7b\xd5\x73\xfb\xaa\x69\xbb\x5d\xed\xdf\x03\x1e\xa5\x5b\xac\xa4\x4b\x6e\x6f\xd2\x89\xdc\x40\xbc\xf9\xb3\xf0\x6f\x22\x06\x1c\x84\x0b\xb6\xa8\x0f\x27\x7f\x88\x44\x3f\x7f\xfd\xd7\xf8\x6f\xe2\xb5\x09\xbe\x0b\xe5\xce\x77\x60\x95\x51\xce\x32\x5b\x63\xa3\x78\x65\x0b\xee\x3b\x78\xb6\x86\x70\x47\x4b\x59\x68\x46\x7d\xfc\x4c\xd2\x71\x90\x8f\x72\x0a\x31\xda\xa2\xd0\x5a\xa4\x33\x21\xa4\x66\x72\x3f\x17\x30\xfc\x79\x7c\xb6\x1b\x13\xbb\x16\xe1\xab\xda\xbf\x1c\xbb\x99\xf8\x0e\x00\x00\xff\xff\x32\x2d\x53\xa1\x2b\x02\x00\x00")

func migrations1688189195_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations1688189195_initUpSql,
		"migrations/1688189195_init.up.sql",
	)
}

func migrations1688189195_initUpSql() (*asset, error) {
	bytes, err := migrations1688189195_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/1688189195_init.up.sql", size: 555, mode: os.FileMode(436), modTime: time.Unix(1709718595, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"migrations/1688189195_init.down.sql": migrations1688189195_initDownSql,
	"migrations/1688189195_init.up.sql":   migrations1688189195_initUpSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"migrations": &bintree{nil, map[string]*bintree{
		"1688189195_init.down.sql": &bintree{migrations1688189195_initDownSql, map[string]*bintree{}},
		"1688189195_init.up.sql":   &bintree{migrations1688189195_initUpSql, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
