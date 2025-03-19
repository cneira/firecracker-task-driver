package libtime

// Many common formats can be found in go's time package, such as ANSIC and RFC,
// so check before adding any format to this file.
//
// Additional formats are described here for convenience.
const (
	// ISO8601Micro is the default joda datetime print format, which is useful
	// when interacting with services written in Java.
	ISO8601Micro = `2006-01-02T15:04:05.000-07:00`
)
