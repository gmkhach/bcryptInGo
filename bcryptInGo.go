package main

// The x/crypto/bcrypt package has to be brought in to your local dev environment before go compiler can recognize it
// Run "go get golang.org/x/crypto/bycrpt" in your terminal to get that package
import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	x := `password123`
	xb, err := bcrypt.GenerateFromPassword([]byte(x), bcrypt.MinCost)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println(x)
	fmt.Println(xb)

	psswdAttempt := `password123`
	err = bcrypt.CompareHashAndPassword(xb, []byte(psswdAttempt))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Successful login!")

}
