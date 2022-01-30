package main

// import "github.com/caarlos0/env"

//////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

type envConfig struct {
	Addr     string `env:"MS_ADDR" envDefault:"localhost:7070"` // Server address
	Compress bool   `env:"MS_COMPRESS"`                         // Enable transparent response compression
	File     string `env:"MS_FILE" envDefault:"mock.json"`      // Mock json file location
	Verbose  int    `env:"MS_VERBOSE" envDefault:"1"`           // Verbose mode (higher numbers increase the verbosity)
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "mockserver"
//          version   = "0.1.0"
//          date = "2022-01-30"

//          e envConfig
//  )
