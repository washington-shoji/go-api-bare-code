package server

import (
	"testing"
)

func TestServer(t *testing.T) {
	// Replace with your desired listen address
	listenAddress := ":7777"

	// Call the Server function to create an APIServer instance
	server := Server(listenAddress)

	// Check that the returned APIServer instance is not nil
	if server == nil {
		t.Errorf("Expected a non-nil APIServer instance, but got nil")
	}

	// Check that the ListenAddress property matches the input
	if server.ListenAddress != listenAddress {
		t.Errorf("Expected ListenAddress to be %s, but got %s", listenAddress, server.ListenAddress)
	}

	// You can add more assertions if your Server function has additional logic.
}
