package main

import (
	"fmt"
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
