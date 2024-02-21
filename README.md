# Go Backend Project

This Go project offers backend services including user management, place recommendations, and trip planning.

## Quick Start

### Prerequisites

- Go (1.x+)

### Setup

Clone the repo and install dependencies:

```
git clone git@github.com:xwd0418/TripPlanningBackend.git
cd go/src/tripPlanning
go mod tidy
```

### Running

Start the server with `go run main.go`

### APIs Overview
#### User APIs
Signup: POST /signup - Register a new user.

Login: POST /login - Authenticate a user.

#### Place APIs

Show Places: GET /showDefaultPlaces?max_num_display=<number> - Get default places.

Search Places: GET /searchPlaces - Search for places.

#### Trip Planning
Generate Trip: POST /generateTripPlan - Create a new trip plan.

Modify Trip: POST /modifyTrip/{tripID} - Update an existing trip.

Delete Trip: DELETE /deleteTrip/{tripID} - Remove a trip.

#### Trip Advising
Generate Trip Advice: GET /recommendation - Use ChatGPT to provide advice to users 



