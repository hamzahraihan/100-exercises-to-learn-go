// Package cliflags teaches the flag package for command-line argument parsing.
package cliflags

type Config struct {
	Port     int
	Debug    bool
	DataFile string
}

// ParseConfig parses args using a FlagSet and returns a Config.
// Flags: -port int (default 8080), -debug bool (default false),
//        -data string (default "data.json").
// TODO: flag.NewFlagSet("server", flag.ContinueOnError); register the three
// flags; fs.Parse(args); return Config{Port:*port, Debug:*debug, DataFile:*data}, nil.
func ParseConfig(args []string) (Config, error) {
	return Config{Port: 8080, DataFile: "data.json"}, nil // TODO: implement
}
