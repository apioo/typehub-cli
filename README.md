
# TypeHub CLI

The CLI provides a simple binary written in go to access the https://typehub.cloud/ API. The following example
shows how you can use the binary.

## Build

To get the binary you can either use one of our pre-build binaries which you can download from our
[release page](https://github.com/apioo/typehub-cli/releases), or you can also simply build the binary
by yourself with:

> go build

## Push

The push command creates or updates a document at TypeHub.

> typehub push [document_name] [schema_file] --client-id="[user]" --client-secret="[password]"

* __document_name__  
  The name of your document
* __schema_file__  
  Optional 
* __--client-id__  
  The client id is either your username or an app key which you can create at our backend.
* __--client-secret__  
  This client secret is either your password or an app secret which you can create at our backend.

