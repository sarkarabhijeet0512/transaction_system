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

# Key Operations and Their Asymptotic Behavior
## Insertion of a Transaction (PUT /transactionservice/transaction/$transaction_id):

## Insertion:
The insertion of a transaction itself is typically O(1) in terms of the operation's time complexity, but it can be affected by indexing overhead, especially if multiple indexes exist.
Time Complexity: O(1) (amortized), assuming the database handles the actual insert efficiently.
Space Complexity: O(1) for the operation itself, but the total space used by the database grows linearly with the number of transactions.
**Retrieving Transactions by Type (GET /transactionservice/types/$type):**
## Query:
The query involves selecting all transactions that match a given type. If the type column is indexed, this operation can be performed efficiently.
### Time Complexity:
O(M log N) where M is the number of transactions matching the type and N is the total number of transactions.
### Space Complexity:
O(M) for storing the result set (the list of matching transaction IDs).
**Summing Transactions Linked by parent_id (GET /transactionservice/sum/$transaction_id):**
## Recursive Query:
This operation involves a recursive query that traverses all transactions linked by parent_id to the specified transaction. The complexity depends on the depth of the tree formed by these links and the number of linked transactions.
### Time Complexity:
In the worst case, if every transaction is linked (forming a chain), the complexity is O(D * log N), where D is the depth of the tree and N is the number of transactions. If the tree is balanced, this can be closer to O(D log N) with D = log N.
### Space Complexity:
O(D) for the call stack during recursion and O(M) for storing the sums of M transactions.
## Foreign Key Constraint Check:

### Query to Verify Parent Exists:
The system checks if a parent_id exists before insertion. This is a simple indexed lookup.
### Time Complexity: O(log N), as mentioned before, assuming an index on transaction_id.
### Space Complexity: O(1).
Database Indexing and Its Impact
Indexes improve query performance, particularly for lookups like the parent_id check and the retrieval of transactions by type. However, maintaining indexes comes with a cost:

## Insertion Time:
Insertion operations may become slightly slower due to the overhead of maintaining the index.
## Space Complexity:
Indexes increase the space complexity as they require additional storage.
## Worst-Case Scenarios
Deeply Nested Transactions: If the transaction tree (formed by parent_id) is very deep (e.g., a long chain), recursive queries for summing transaction amounts can become expensive. In the worst case, they could approach O(N log N) if every transaction in the database is linked in a single chain.

Large Volume of Transactions of a Single Type: If a single type dominates the dataset, retrieving all transactions by that type could be costly, approaching O(N log N) for large N.

## Trade-offs and Optimization
### Indexing: 
Carefully choosing which fields to index (e.g., transaction_id, type, and parent_id) can optimize query performance, but at the cost of slower inserts and more storage.

### Batching and Caching: 
For high-load scenarios, consider batching inserts and caching frequently accessed sums or transaction lists to reduce the load on the database.




