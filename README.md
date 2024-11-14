
# TypeHub CLI

The CLI provides a simple binary written in go to push TypeAPI/TypeSchema specifications to the TypeHub platform.
The following example shows how you can use the binary.

## Build

To get the binary you can either use one of our pre-build binaries which you can download from our
[release page](https://github.com/apioo/typehub-cli/releases), or you can also simply build the binary
by yourself with:

> go build

## Push

The push command imports a document at TypeHub.

> typehub push [document_name] [schema_file] --client-id="[user]" --client-secret="[password]"

* __document_name__  
  The name of your document.
* __schema_file__  
  Contains the TypeAPI specification which you want to push to TypeHub.
* __--client-id__  
  The client id is either your username or an app key which you can create at our backend.
* __--client-secret__  
  This client secret is either your password or an app secret which you can create at our backend.

## Pushd

The pushd command imports all documents inside the provided directory to TypeHub.

> typehub pushd [directory] --client-id="[user]" --client-secret="[password]"

* __directory__  
  A folder containing TypeAPI/TypeSchema specifications. 
* __--client-id__  
  The client id is either your username or an app key which you can create at our backend.
* __--client-secret__  
  This client secret is either your password or an app secret which you can create at our backend.

## Docker

We provide also a [docker image](https://hub.docker.com/r/apiootech/typehub-push) which you can use to push documents to
the TypeHub platform. To use the image you can simply run:

> docker compose up

This mounts the `input/` folder and imports every specification inside the directory to the TypeHub platform.
Before you start this command you need to provide your credentials at the `.env` file.
