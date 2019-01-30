##### Motiviation go-rdf

Generate golang structs, validation, en/decoding from various rdf flavours like owl, jsonld, schema.org and friends.


##### structure


```
.
 ├── cache // cache for all rdf definitions, contains .html, .jsonld, .xml (linked data, yo)
 ├── docs
 ├── pkg // tools for traversing the linked data graph, normalization and code generation, repo root needs to be working directory
 ├── gen
    └── schema // contains structs generated from `schema.org`
```
