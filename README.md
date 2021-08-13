# Growatt You Like 🌱️

## 🚨️ Under Construction 🚨️

## Objectives
This is a hobby project with the purpose of getting and keeping my solar
panel/inverter data. It is heavily inspired by the very excellent
[grott](https://github.com/johanmeijer/grott) system by
[johanmeijer](https://github.com/johanmeijer).

How it differs is that the aim is to get the data into the cloud first,
then conduct any processing/decoding there once the data has been secured.
The idea is to keep it as simple as possible to reduce the risk of a bug
getting in the way of storing the data. With the data in the cloud
we can retry processing anytime.

As with [grott](https://github.com/johanmeijer/grott) the data is also sent
to the Growatt central server, though arguably that could also happen in
the cloud.

## Components
### Cloud Infrastructure
Cloud infrastructure is [AWS IOT](https://aws.amazon.com/iot/) and is built via Terraform.
Data is sent to AWS IOT via MQTT. Authentication by x.509 certificates. See Terraform
[README.md](./terraform/README.md).

### Mock Growatt
Mock Growatt is a small socket application that pretends to be the Growatt server.
For all intents and purposes it could pretend to be any socket application.
See Mock Growatt [README.md](./mock-growatt/README.md).

### Growatt To IOT
Growatt to IOT is the application that connects to AWS IOT MQTT and proxies messages
from the socket to an MQTT topic. By default, it sends the messages to the Growatt
Server as well.

See Growatt To IOT [README.md](./growatt-to-iot/README.md).