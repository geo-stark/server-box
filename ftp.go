// Copyright 2018 The goftp Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// This is a very simple ftpd server using this library as an example
// and as something to run tests against.

package main

import (
	"log"
	"strconv"

	filedriver "gitea.com/goftp/file-driver"
	"goftp.io/server"
)

func ftpServer(options map[string]string) {

	root := getOpt(options, "root", "")                                  // "Root directory to serve"
	user := getOpt(options, "user", "/")                                 // "Username for login"
	pass := getOpt(options, "password", "")                              // "Password for login"
	port, _ := strconv.ParseInt(getOpt(options, "port", "2121"), 10, 16) // Port
	host := getOpt(options, "host", "localhost")                         // Host

	if root == "" {
		log.Fatalf("root directory is not specified")
	}
	if user == "" || password == "" {
		log.Fatalf("user or/and password is not specified")
	}

	factory := &filedriver.FileDriverFactory{
		RootPath: root,
		Perm:     server.NewSimplePerm("", ""),
	}

	opts := &server.ServerOpts{
		Factory:  factory,
		Port:     int(port),
		Hostname: host,
		Auth:     &server.SimpleAuth{Name: user, Password: pass},
	}

	log.Printf("Starting ftp server on %v:%v", opts.Hostname, opts.Port)
	log.Printf("Username %v, Password %v", user, pass)
	server := server.NewServer(opts)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
