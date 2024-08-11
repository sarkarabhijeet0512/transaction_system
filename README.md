# Transaction Backend
Transaction system backend

Note:Tables Get auto Created when project is started
need to add the resource rows.

Note :Table Structure Changed here as id is a primary key and it should be auto increment therefore 
adding transaction_id with uinque index for making workaround according to task's requirement.

## Run the project in local 
cd cmd
go run .

## To generate new errors using stringer
cd er
go generate

# transactionservice
## transaction Registration 
### PUT http://localhost:8765/v1/transactionservice/transaction/:transaction_id

```
{
    "amount": 5000,
    "type": "cars",
}
```

## transaction Registration 
### PUT http://localhost:8765/v1/transactionservice/transaction/:transaction_id

```
{
    "amount": 10000,
    "type": "shopping",
    "parent_id": 10
}
```
 
## SUM of Transaction 
### GET http://localhost:8765/v1/transactionservice/sum/:transaction_id

### OUTPUT
```
{
    "success": true,
    "message": "Success",
    "data": {
        "sum": 15000
    },
    "meta": {}
}
```

## Get transction by type
### GET http://localhost:8765/v1/transactionservice/types/:type

### OUTPUT

```
{
    "success": true,
    "message": "Success",
    "data": [
        10
    ],
    "meta": {}
}
```

