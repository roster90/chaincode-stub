package main

/**
 * tokenv5
 * Shows the use of ChaincodeStub API for getting & setting key/value pairs
 **/
import (
	"fmt"
	"os"

	// April 2020, Updated for Fabric 2.0
	// Video may have shim package import for Fabric 1.4 - please ignore

	"github.com/hyperledger/fabric-chaincode-go/shim"

	peer "github.com/hyperledger/fabric-protos-go/peer"

	// Conversion functions - needed to convert MyToken to Integer
	"strconv"
)

// TokenChaincode Represents our chaincode object
type TokenChaincode struct {
}

// Init Implements the Init method
func (token *TokenChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Simply print a message
	fmt.Println("Init executed")

	// Lets create a Key-Value pair
	stub.PutState("MyToken", []byte("2000"))

	// Return success
	return shim.Success([]byte("true"))
}

// Invoke method
func (token *TokenChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	funcName, _ := stub.GetFunctionAndParameters()
	fmt.Println("Function=", funcName)

	// V5
	if funcName == "set" {
		// Sets the value
		return SetToken(stub)

	} else if funcName == "get" {

		// Gets the value
		return GetToken(stub)

	}

	// This is not good
	return shim.Error(("Bad Function Name = " + funcName + "!!!"))
}

// SetToken inrements the value of the token by 10
// V5
// Returns true if successful
func SetToken(stub shim.ChaincodeStubInterface) peer.Response {

	// Get the current value
	value, err := stub.GetState("MyToken")

	// If there is error in retrieve send back an error response
	if err != nil {
		return shim.Error(err.Error())
	}

	// Convert value to integer
	intValue, err := strconv.Atoi(string(value))

	// If there is an error in conversion - return false
	if err != nil {
		// May also return sh.Error
		return shim.Success([]byte("false"))
	}

	// Increment the value by 10
	intValue += 10

	// Execute PutState - overwrites the current value
	stub.PutState("MyToken", []byte(strconv.Itoa(intValue)))

	return shim.Success([]byte("true"))
}

// GetToken reads the value of the token from the database
// V5
// Reurns the value or -1 in case MyToken doesn't exist
func GetToken(stub shim.ChaincodeStubInterface) peer.Response {
	// Holds a string for the response
	var myToken string

	// Local variables for value & error
	var value []byte
	var err error

	if value, err = stub.GetState("MyToken"); err != nil {

		fmt.Println("Get Failed!!! ", err.Error())

		return shim.Error(("Get Failed!! " + err.Error() + "!!!"))

	}

	// nil indicates non existent key
	if value == nil {
		// Return value -1 is to indicate to caller that MyToken
		// Does NOT exist in state data
		myToken = "-1"

	} else {

		myToken = "MyToken=" + string(value)

	}

	return shim.Success([]byte(myToken))
}

// type serverConfig struct {
// 	CCID    string
// 	Address string
// }

// Chaincode registers with the Shim on startup
func main() {
	ccid := os.Getenv("CORE_CHAINCODE_ID_NAME")
	address := os.Getenv("CHAINCODE_SERVER_ADDRESS")

	fmt.Printf("Started Chaincode. token/v5\n")
	fmt.Printf("Chaincode server address: %s\n", address)
	fmt.Printf("Chaincode ID: %s\n", ccid)

	server := &shim.ChaincodeServer{
		CCID:    ccid,
		Address: address,
		CC:      new(TokenChaincode),
		TLSProps: shim.TLSProperties{
			Disabled: true, // nếu bạn chưa bật TLS (để debug đơn giản trước)
		},
	}

	if err := server.Start(); err != nil {
		fmt.Printf("Error starting chaincode: %s\n", err)
		os.Exit(1)
	}
}
