package main

// define ur levels here

var Info = LogLevel{
	Name:  "INFO",
	style: InfoStl,
}

var Debug = LogLevel{
	Name:  "DEBUG",
	style: DebugStl,
}

var Warning = LogLevel{
	Name:  "WARN",
	style: WarningStl,
}

var Error = LogLevel{
	Name:  "ERROR",
	style: ErrorStl,
}

var Fatal = LogLevel{
	Name:  "FATAL",
	style: ErrorStl,
}
