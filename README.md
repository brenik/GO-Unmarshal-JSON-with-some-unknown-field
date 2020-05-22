# GO-Unmarshal-JSON-with-some-unknown-field

If you need to get some fields from JSON in Go, but you don’t know the whole structure.

Small transformations will help us.

File data.json contains the classifier of economic activities in JSON format,

and has many different fields, but need only one. 

Not having the described structure is hard to get the right field. 

We use an array of empty interfaces to obtain the necessary data.
