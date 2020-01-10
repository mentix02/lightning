# Lightning Structures

Even though Lightning was written with the thought process of being **simple** yet fast, some time there are instances
where the standard way of doing things just isn't enough and thus one has to resort to building another layer of
abstraction over what's already been provided. That is why the [`structs.go`](../structs.go) file exists. It will include
a number of useful data structures that help in parsing or storing data retrieved from the MySQL database. An example - 

In the [`getAuthorBookmarkedArticlesPks`](../handlers.go#L10) handler, we iterate over the results fetched from the
database incrementally, i.e. one by one. If we were to use a standard slice that would have an article primary key
appended with every iteration, it'd have caused Go to go through the 
[long and arduous process of resizing a slice](https://dev.to/andyhaskell/a-closer-look-at-go-s-slice-append-function-3bhb).

So what data structure should we use when we are just going to access the elements in a collections only once but are
probably going to append new elements to it a lot? A list. So should we use Go's `container/list`? Well, that's a doubly
linked list - every node contains pointers to the previous as well as the next one. We don't really care about traversing
the list backwards. Another drawback? It stores data as an `interface` - that's not really required since we know that
primary keys are only going to be an unsigned integer. Even if we ignore these two big issues, we still have to go through
the process of converting the list to a slice as is required by Go's `encoding/json` package. 

Thus in the end, we have to have a data structure specific to our use case and thus we have a custom [List](../structs.go#L31)
implemented in [`structs.go`](../structs.go) along with appropriate helper methods to convert it to a slice. Problem
solved!

Another problem that the [`structs.go`](../structs.go) file solves is that of custom JSON responses. As we see a lot in
RESTful APIs, there's a lot of common responses that we have to send out - 400 errors (validation, not found, 
unauthorized, etc), 200 successes (created, ok, deleted, etc), and since we're trying to imitate a (faster) rest_framework
server, we have to send some body in JSON format. Instead of building an anonymous `map[string]string` every time we have
to send a response, we can just use the [`DetailResponse`](../structs.go#L12) structure since it can be converted into a
JSON valid string.

## Data Structures

As explained above, there's a number of reasons to host a home brewed library of structures to hold data from the database.
Any future models describing the Tables will also live in [`structs.go`](../structs.go). For now, there's only a few data
types that have been implemented - 

| # | Name                      | Description                                     | Fields                                    |
|---|---------------------------|-------------------------------------------------|-------------------------------------------|
| 1 | [Node](../structs.go#L23) | The base element type used in building List.    | `value interface{}, next *Node`           |
| 2 | [List](../structs.go#L31) | A singly linked list used to hold primary keys. | `len uint32, head *Node, tail *Node`      |

## JSON Responses

These are **simple** data types, usually containing only a single field. They might or might not have any methods.

| # | Name                               | Description                                           | Fields          |
|---|------------------------------------|-------------------------------------------------------|-----------------|
| 1 | [DetailResponse](../structs.go#L12) | The default JSON response holder to return data with. | `Detail string` |

## Further Reading

As can be expected from a project so early in its development stages, more custom definitions can be expected with proper
methods and documentation.
