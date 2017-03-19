# Dynamo Document Server in Go

You'll firstly need to populate user credentials in
`~/.aws/credentials`, and otherwise do general AWS setup.

## Dependencies

### Create a Dynamo Table to Use
Note: this will be introduced to code shortly, but may continue to be
valuable for manual testing in the future.

    aws dynamodb create-table \
        --table-name Documents \
        --key-schema \
            AttributeName=documentId,KeyType=HASH \
            AttributeName=due,KeyType=RANGE \
        --attribute-definitions \
            AttributeName=documentId,AttributeType=S \
            AttributeName=due,AttributeType=S \
        --provisioned-throughput
            ReadCapacityUnits=1,WriteCapacityUnits=1

### Get Gorilla Mux
    go get "github.com/gorilla/mux"

## Build and Run
    go build
    ./dynago-docs &

## Use
### Create
    curl \
        -H "Content-Type: application/json" \
        -d '{ "name": "New Document" }' \
        http://localhost:8080/documents
### Get
    curl -X GET http://localhost:8080/documents/064ae763-3c0b-43ad-9054-6fe46f566dfe

### List
    curl -X GET http://localhost:8080/documents


## Credits / See Also
This was forked from [tns-restful-json-api][source], which goes along
with [this tutorial][tutorial].

The AWS SDK for Go includes an example of [how to scan and list
results][aws-example].

[source]: https://github.com/corylanou/tns-restful-json-api
[tutorial]: https://thenewstack.io/make-a-restful-json-api-go/
[aws-example]: https://github.com/aws/aws-sdk-go/blob/master/example/service/dynamodb/scanItems/scanItems.go

