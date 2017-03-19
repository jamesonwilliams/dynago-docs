package db

/*
 * Note: there is a very useful example of using the AWS SDK for Go in
 * the github examples repository:
 *
 * https://github.com/aws/aws-sdk-go/tree/master/example/service/dynamodb/scanItems
 */
import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jamesonwilliams/dynago-docs/model"
	"log"
	"os"
)

type DynamoDatabase struct {
	Region    string
	TableName string
}

func (ddb *DynamoDatabase) RetrieveDocuments() ([]model.Document, error) {
	sess := session.New(&aws.Config{Region: aws.String(ddb.Region)})
	svc := dynamodb.New(sess)

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		TableName: aws.String(ddb.TableName),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	if err != nil {
		fmt.Println("failed to make Query API call, %v", err)
		os.Exit(1)
	}

	documents := []model.Document{}

	// Unmarshal the Items field in the result value to the Item Go type.
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &documents)
	if err != nil {
		fmt.Println("failed to unmarshal Query result documents, %v", err)
		os.Exit(1)
	}

	return documents, nil
}

func (ddb *DynamoDatabase) RetrieveDocument(id string) (model.Document, error) {
	sess := session.New(&aws.Config{Region: aws.String(ddb.Region)})
	svc := dynamodb.New(sess)

	params := &dynamodb.QueryInput{
		TableName: aws.String(ddb.TableName), // Required
		Select:    aws.String("ALL_ATTRIBUTES"),

		ConsistentRead: aws.Bool(true),
		KeyConditions: map[string]*dynamodb.Condition{
			"documentId": { // Required
				ComparisonOperator: aws.String("EQ"), // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					{ // Required
						S: aws.String(id),
					},
				},
			},
		},
	}

	resp, err := svc.Query(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return model.Document{}, err
	}

	// Pretty-print the response data.
	fmt.Println(resp)

	documents := []model.Document{}
	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &documents)
	fmt.Println(documents[0])

	return documents[0], err
}

func (ddb *DynamoDatabase) StoreDocument(document model.Document) (model.Document, error) {
	av, err := dynamodbattribute.MarshalMap(document)
	if err != nil {
		log.Println("Failed to marshal data", err)
		return document, err
	}

	input := &dynamodb.PutItemInput{
		Item:                av,
		TableName:           aws.String(ddb.TableName),
		ConditionExpression: aws.String("attribute_not_exists(documentId)"),
	}
	log.Println(input)

	sess := session.New(&aws.Config{Region: aws.String(ddb.Region)})
	svc := dynamodb.New(sess)
	result, err := svc.PutItem(input)
	if err != nil {
		log.Println("Failed to save item to table", err)
	}

	log.Println("Successfully added item to table.", result)

	return document, err
}

func (ddb *DynamoDatabase) DeleteDocument(id int) error {
	return nil
}
