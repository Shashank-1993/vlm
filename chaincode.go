/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a asset
type SmartContract struct {
	contractapi.Contract
}

// asset describes basic details of what makes up a asset
type asset struct {
	VehicleDetails      VehicleDetails 		`json:"vehicleDetails"`
	CitizenDetails      CitizenDetails 		`json:"citizenDetails"`
	InsuranceDetails    InsuranceDetails    `json:"insuranceDetails"`
	DocumnetHash        DocumnetHash        `json:"documnetHash"`
}

type VehicleDetails struct {
	EngineNumber 	string `json:"engineNumber"`
	ChassisNumber   string `json:"chassisNumber"`
	InvoicedAmount  string `json:"invoicedAmount"`
	Date   			string `json:"date"`
}

type CitizenDetails struct {
	Name        	string `json:"name"`
	DateOfBirth 	string `json:"dateOfBirth"`
	AadharNumber  	string `json:"aadharNumber"`
	AddressProof  	string `json:"addressProof"`
}

type InsuranceDetails struct {
	InsurerCompany     string `json:"insurerCompany"`
	InsuredAmount 	   string `json:"insuredAmount"`
	ValidTill          string `json:"validTill"`
}

type DocumnetHash struct{
	Form20     				string `json:"form20"`
	Form21 	   				string `json:"form21"`
	TemporaryRegistration   string `json:"temporaryRegistration"`
	InsuranceDoc            string `json:"insuranceDoc"`
	
}

// QueryResult structure used for handling result of query
type QueryResult struct {
	Record *asset
}

// CreateRecord adds a new asset to the world state with given details
func (s *SmartContract) createRecord(ctx contractapi.TransactionContextInterface,
	string engineNumber,
	string chassisNumber, 
	string invoicedAmount,
	string date,
	string name,
	string dateOfBirth,
	string aadharNumber,
	string addressProof,
	string insurerCompany,
	string insuredAmount,
	string validTill,
	string form20,
	string form21,
	string temporaryRegistration,
	string insuranceDoc
	) {

	vehicleDetails := VehicleDetails{
		EngineNumber 	: engineNumber,
		ChassisNumber 	: chassisNumber,
		InvoicedAmount  : invoicedAmount,
		Date			: date
	}

	citizenDetails := CitizenDetails{
		Name 			: name,
		DateOfBirth 	: dateOfBirth,
		AadharNumber  	: aadharNumber,
		AddressProof	: addressProof
	}

	insuranceDetails := InsuranceDetails{
		InsurerCompany 	: insurerCompany,
		insuredAmount 	: insuredAmount,
		ValidTill  		: validTill
	}

	documnetHash := DocumnetHash{
		Form20 					: form20,
		Form21 					: form21,
		TemporaryRegistration   : temporaryRegistration,
		InsuranceDoc			: insuranceDoc
	}
	
	asset := asset{
		VehicleDetails 		vehicleDetails,
		CitizenDetails 		citizenDetails,
		InsuranceDetails 	insuranceDetails,
		DocumnetHash 		documnetHash
	}

	assetAsBytes, _ := json.Marshal(asset)

	return ctx.GetStub().PutState(chassisNumber, assetAsBytes)
}

// Queryasset returns the asset stored in the world state with given id
func (s *SmartContract) QueryRecordByChassisNumber(ctx contractapi.TransactionContextInterface, chassisNumber string) (*asset, error) {
	assetAsBytes, err := ctx.GetStub().GetState(chassisNumber)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if assetAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", id)
	}

	RetrivedAsset := new(asset)
	_ = json.Unmarshal(assetAsBytes, RetrivedAsset)

	return RetrivedAsset, nil
}

// // QueryAllassets returns all assets found in world state
// func (s *SmartContract) QueryAllRecords(ctx contractapi.TransactionContextInterface) ([]asset, error) {
// 	startKey := ""
// 	endKey := ""

// 	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resultsIterator.Close()

// 	results := []asset{}

// 	for resultsIterator.HasNext() {
// 		queryResponse, err := resultsIterator.Next()

// 		if err != nil {
// 			return nil, err
// 		}

// 		asset := new(asset)
// 		_ = json.Unmarshal(queryResponse.Value, asset)

// 		// queryResult := asset
// 		results = append(results, *asset)
// 	}

// 	return results, nil
// }

// // QueryRecordHistory returns the asset stored in the world state with given id
// func (s *SmartContract) QueryRecordHistory(ctx contractapi.TransactionContextInterface, id string) (string, error) {

// 	resultsIterator, err := ctx.GetStub().GetHistoryForKey(id)
// 	if err != nil {
// 		return "nil", err
// 	}
// 	defer resultsIterator.Close()

// 	// buffer is a JSON array containing historic values for the record
// 	var buffer bytes.Buffer
// 	buffer.WriteString("[")

// 	bArrayMemberAlreadyWritten := false
// 	for resultsIterator.HasNext() {
// 		response, err := resultsIterator.Next()
// 		if err != nil {
// 			return "nil", err
// 		}
// 		// Add a comma before array members, suppress it for the first array member
// 		if bArrayMemberAlreadyWritten == true {
// 			buffer.WriteString(",")
// 		}
// 		buffer.WriteString("{\"TxId\":")
// 		buffer.WriteString("\"")
// 		buffer.WriteString(response.TxId)
// 		buffer.WriteString("\"")

// 		buffer.WriteString(", \"Value\":")

// 		if response.IsDelete {
// 			buffer.WriteString("null")
// 		} else {
// 			buffer.WriteString(string(response.Value))
// 		}

// 		buffer.WriteString(", \"Timestamp\":")
// 		buffer.WriteString("\"")
// 		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
// 		buffer.WriteString("\"")

// 		buffer.WriteString(", \"IsDelete\":")
// 		buffer.WriteString("\"")
// 		buffer.WriteString(strconv.FormatBool(response.IsDelete))
// 		buffer.WriteString("\"")

// 		buffer.WriteString("}")
// 		bArrayMemberAlreadyWritten = true
// 	}
// 	buffer.WriteString("]")

// 	return buffer.String(), nil
// }

// // QueryAllRecordsByProject returns all assets found in world state
// func (s *SmartContract) QueryAllRecordsByProject(ctx contractapi.TransactionContextInterface, projectName string) ([]asset, error) {
// 	startKey := ""
// 	endKey := ""

// 	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resultsIterator.Close()

// 	results := []asset{}

// 	for resultsIterator.HasNext() {
// 		queryResponse, err := resultsIterator.Next()

// 		if err != nil {
// 			return nil, err
// 		}

// 		asset := new(asset)
// 		_ = json.Unmarshal(queryResponse.Value, asset)

// 		// queryResult := asset
// 		if asset.ProjectDetails.Name == projectName {
// 			results = append(results, *asset)
// 		}

// 	}

// 	return results, nil
// }

// // UpdateRecordStatus updates the owner field of asset with given id in world state
// func (s *SmartContract) UpdateRecordStatus(ctx contractapi.TransactionContextInterface,
// 	ID string,
// 	status string,
// 	updater string,
// 	user string,
// 	updatedDate string) error {
// 	asset, err := s.QueryRecordByRegNumber(ctx, ID)

// 	if err != nil {
// 		return err
// 	}

// 	if updater == "BANK" {
// 		asset.BankStatus = status
// 		asset.BankReviewer = user
// 		asset.BankReviewdAt = updatedDate
// 	} else if updater == "BENEFICIARY" {
// 		asset.BeneficiaryStatus = status
// 		asset.ReviewedBy = user
// 		asset.ReviewedAt = updatedDate
// 	} else if updater == "BGREQ" {
// 		asset.BGRequirements.Status = status
// 		asset.BGRequirements.StatusUpdatedBy = user
// 		asset.BGRequirements.StatusUpdatedAt = updatedDate
// 	}

// 	assetAsBytes, _ := json.Marshal(asset)

// 	return ctx.GetStub().PutState(ID, assetAsBytes)
// }

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create VehicleManagement chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting BGasset chaincode: %s", err.Error())
	}
}
