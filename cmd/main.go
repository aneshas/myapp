package main

import "github.com/aneshas/myapp"

func main() {

	srvr := NewServer(
		myapp.NewProvisionAccount(nil /* inject the dependency */),
		// Additional use cases would follow the same pattern. For eg:
		// myapp.NewAddFunds(...),
		// myapp.NewCloseAccount(...),
	)

	_ = srvr

	// Start your server
}
