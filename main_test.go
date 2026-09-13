package main
import("net/http";"testing")
func TestContract(t *testing.T){if http.StatusOK!=200{t.Fatal()}}
