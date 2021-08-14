## Growatt To IOT

## TLS Files
To connect to AWS IOT we need the TLS files and the AWS ATS Root CA file. Follow the steps in 
the Terraform [README.md](../terraform/README.md) to get the TLS files. By default the application
assumes the files are in the same folder it is running from, so just copy them over.

## MQTT Topic
The growat MQTT topic and 

## TODO
* A fatal error in the chain should not stop the rest of the chain
* An EOF in the listen should not stop the program - instead
restart (could be supervisord/docker-compose or the app?)
* Try and capture program closure and let each handler close their connections
nicely
* Let user configure via ENV vars and a config file
* Commenting / tests / idiomatic Go == quality
* Correct use of pointers/references == ensure we're not over allocating (re-using structs)
* Compile and dockerfile

## Questions for myself!
* Is this really a growatt service? With some doc changes it could be just a
socket packet to AWS IOT program... which might have a very wide range of
applications. The experiment is to refactor into something generic that you
can plug in your own processing.