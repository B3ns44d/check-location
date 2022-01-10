this is a golang service to check user location using their IP address.

The purpose of this service is to provide a simple way to check if a user is located in a particular country. This is part of an app I'm working on and this is a simple service from it.


### Usage

Make a request to `http://127.0.0.1:3005/geo/139.130.4.5`, for example. You can omit an IP address to use your current IP address, or replace to use another. If the IP address is invalid, a HTTP 400 is returned.

Examples: 
```bash

curl -X GET http://127.0.0.1:3005/geo/139.130.4.5
```

[![Publish Docker Image](https://github.com/B3ns44d/check-location/actions/workflows/docker-image.yml/badge.svg)](https://github.com/B3ns44d/check-location/actions/workflows/docker-image.yml)