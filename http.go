package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// type Reader interface {
//   Read(p []byte) (n int, err error)
// }

// ^ In the above, the Reader Interface's 'Read' method will take an empty byte array.

func makeGetRequest() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// fmt.Println(resp)
	// bs := []byte{}
	bs := make([]byte, 99999)

	// Make is a built in function that takes a type of a slice, and than as a second argument, initializes the slice with a number of empty spaces.

	// The reason we initialized the byte slice with such a giant amount of capacity is because the Read function is not set up to automatically resie the
	// slice if the slice is already full.

	// [, , , , , , , , , , , , ,  ... n-1, n] v
	//                                          ---->  Read
	// Real source of data, e.g. HTML response ^.

	resp.Body.Read(bs)
	fmt.Println(string(bs))
	// 1. Response Body will take empty byte slice passed into Read function as it's parameter.
	// 2. Response Body will pump it's HTML response into that empty byte slice into the 'Read' function.

	// Go does have a lot of built-in helper functions for automatically working with Reader interfaces, and getting data out of Reader type, and into the terminal.

	//======= QUICKER WAY OF DOING THE ABOVE =========//
	io.Copy(os.Stdout, resp.Body)
	// Copy just pipes data from Reader to Writer

	// 1. Take byte slice, and pass it to some value that implements a "Writer" interface.
	// 2. Writer interface describes something that can take info and send it outside of our program.
	// 3. Writer is used to write outgoing request or write some information to a text file on your HDrive, or to make an outgoing HTTP request, or log to a terminal, etc.
	// 4. For us to use the "Writer" interface we have to go through the Go standard library and find something that implements the Writer interface, and use that to log info
	// 5. In order for something to implement Writer interface, it needs to use Write method on "Writer" inferface
	// e.g. => Write(p []byte) (n int, err error)
	// 6. io.Copy function takes in a Writer type, and Reader type as parameters.
	// 7.  io.Copy(os.Stdout => "*File" type implements Writer interface, resp.Body => implements Reader type)
}

// HTTP has a property of BODY which is of type io.ReadCloser
// ReadCloser is an Interface with two properties, Reader, and Closer
// Reader, and Closer are ALSO Interfaces
// Reader has a method called Read([]byte) (int, err)
// Closer has a method called Close() (error)

// If value in struct is an interface, it means you can use any value for that type, as long as it is a part of that interface.

// That being said, we can create our own type, as long as it has the Read and Close methods needed for it to inherit the ReadCloser Interface
// You can take multiple Go interfaces, and create another interface.
// ReadCloser interface
//   --Reader interface
//   --Closer interface
