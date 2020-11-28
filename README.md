# conf

A golang package for reading config file

# Import package

`import "github.com/minph/conf"`

- [Overview](#pkg-overview)
- [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>

## <a name="pkg-index">Index</a>

- [type Reader](#Reader)
  - [func New(path string) \*Reader](#New)
  - [func (r \*Reader) Get(groupName, itemName string) (value string, key string, err error)](#Reader.Get)
  - [func (r \*Reader) GetItems(itemName string) []string](#Reader.GetItems)
  - [func (r \*Reader) GetValue(groupName, itemName string) string](#Reader.GetValue)
  - [func (r \*Reader) UpdateContent()](#Reader.UpdateContent)

## <a name="Reader">type</a> Reader

```go
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

```

Reader info struct of config file

### <a name="New">func</a> New

```go
func New(path string) *Reader
```

New to create a Reader

will do UpdateContent() automatically

### <a name="Reader.Get">func</a> (\*Reader) Get

```go
func (r *Reader) Get(groupName, itemName string) (value string, key string, err error)
```

Get value by using groupname and keyname

### <a name="Reader.GetItems">func</a> (\*Reader) GetItems

```go
func (r *Reader) GetItems(itemName string) []string
```

GetItems will return all suited itemnames from all groupnames

### <a name="Reader.GetValue">func</a> (\*Reader) GetValue

```go
func (r *Reader) GetValue(groupName, itemName string) string
```

GetValue Get value by using GroupName & KeyName

Force return a value which is "" default

### <a name="Reader.UpdateContent">func</a> (\*Reader) UpdateContent

```go
func (r *Reader) UpdateContent()
```

UpdateContent update data from path
will provide path data for conf.Reader
