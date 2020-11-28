package conf

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Reader info struct of config file
type Reader struct {
	// Value data
	Value map[string]string

	// Path the config file path
	Path string

	// Delimiter the delimiter between groupname and keyname
	// default as ":"
	Delimiter string

	// Group the groupName list
	Group []string
}

// New to create a Reader
//
// will do UpdateContent() automatically
func New(path string) *Reader {
	var group []string
	value := make(map[string]string)
	info := Reader{Value: value, Group: group, Path: path, Delimiter: ":"}
	info.UpdateContent()
	return &info
}

// UpdateContent update data from path
// will provide path data for conf.Reader
func (r *Reader) UpdateContent() {

	path := r.Path
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("File open error: ", err)
		return
	}

	// create a reader
	reader := bufio.NewReader(f)

	// groupCache
	// make to storage the data of groupname
	groupCache := ""

	for {
		// bytes row data
		bytes, _, err := reader.ReadLine()

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("File open error: ", err)
		}

		// line data trim spaces
		line := strings.TrimSpace(string(bytes))

		// if is a comment line then goto the next row
		if strings.Index(line, "#") == 0 {
			continue
		}

		groupName := ""
		groupNameStart := strings.Index(line, "[")
		groupNameEnd := strings.LastIndex(line, "]")
		if groupNameStart > -1 && groupNameEnd > -1 && groupNameEnd > groupNameStart+1 {
			groupName = strings.TrimSpace(line[groupNameStart+1 : groupNameEnd])

		}

		// update cache
		if len(groupName) != 0 {
			groupCache = groupName
			// 保存组名
			r.Group = append(r.Group, groupCache)
		}

		// find the index of "="
		index := strings.Index(line, "=")

		// not found and skip
		if index < 0 {
			continue
		}

		// key name
		key := strings.TrimSpace(line[:index])

		if len(key) == 0 {
			continue
		}

		// value name
		value := strings.TrimSpace(line[index+1:])

		// comment found the position of a comment
		comment := strings.Index(value, "#")

		// if had comment then make a value slice
		if comment > -1 {
			value = value[0:comment]

		}

		// if start with "#" then skip
		if len(value) == 0 {
			continue
		}

		r.Value[groupCache+":"+key] = strings.TrimSpace(value)
	}

	// close file and check error
	if err := f.Close(); err != nil {
		fmt.Println("File close error: ", err)
		return
	}
}

// Get value by using groupname and keyname
func (r *Reader) Get(groupName, itemName string) (value string, key string, err error) {
	key = groupName + r.Delimiter + itemName
	value, canFound := r.Value[key]

	if !canFound {
		return "", key, fmt.Errorf("Can not get item as [%v%v%v]", groupName, r.Delimiter, itemName)
	}
	return value, key, nil
}

// GetValue Get value by using GroupName & KeyName
//
// Force return a value which is "" default
func (r *Reader) GetValue(groupName, itemName string) string {
	key := groupName + r.Delimiter + itemName
	value, canFound := r.Value[key]
	if !canFound {
		return ""
	}
	return value
}

// GetItems will return all suited itemnames from all groupnames
func (r *Reader) GetItems(itemName string) []string {
	var result []string
	for _, groupName := range r.Group {
		value, _, err := r.Get(groupName, itemName)
		if err != nil {
			continue
		}
		result = append(result, value)
	}
	return result
}
