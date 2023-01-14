# Commands

To run the API:

```
go run ./cmd/api/
```

To run tests:

```
go test ./cmd/api/
```

# Requests

Examples of GET-requests and their responses:
<br></br>

#### localhost:8080/

```
{
  "error": false,
  "message": "Flight API"
}
```

<br></br>

#### localhost:8080/calculate

The initial endpoint for the Calc functionality which returns default response:

```
{
  "error": false,
  "message": "Calc API"
}
```

<br></br>

#### localhost:8080/calculate?path=IND,EWR,SFO,ATL,GSO,IND,ATL,GSO&output=count

The requests return count of transfers:

```
{
  "error": false,
  "message": "Count",
  "data": 4
}
```

<br></br>

#### localhost:8080/calculate?path=IND,EWR,SFO,ATL,GSO,IND,ATL,GSO&output=destination

The endpoint with such parameters returns starting and ending points of the travel:

```
{
  "error": false,
  "message": "Destinations",
  "data": [
    "EWR",
    "SFO"
  ]
}
```

<br></br>

#### localhost:8080/calculate?path=IND,EWR,SFO,ATL,GSO,IND,ATL,GSO&output=itinerary

The response contains entire itinerary in proper order:

```
{
  "error": false,
  "message": "Itinerary",
  "data": [
    [
      "SFO",
      "ATL"
    ],
    [
      "ATL",
      "GSO"
    ],
    [
      "GSO",
      "IND"
    ],
    [
      "IND",
      "EWR"
    ]
  ]
}
```

<br></br>

#### localhost:8080/calculate?path=IND,EWR,SFO,ATL,GSO,IND,ATL,GSO

By default, the calculation-endpoint, without the output-parameter specification, returns everything above:

* count of transfers;
* starting and ending points of the travel, i.e. destination;
* and entire itinerary in proper order;

```
{
  "error": false,
  "message": "Info",
  "data": {
    "destination": [
      "EWR",
      "SFO"
    ],
    "itinerary": [
      [
        "SFO",
        "ATL"
      ],
      [
        "ATL",
        "GSO"
      ],
      [
        "GSO",
        "IND"
      ],
      [
        "IND",
        "EWR"
      ]
    ],
    "count": 4
  }
}
```

<br></br>

# Improvements:

* Making API more robust by adding additional validation for requests parameters, e.g.:
    * if there is proper amount of stops — i.e. it is supposed to be even;
    * if all of the stops are the same data type, i.e. strings;
    * if the stops are properly paired, i.e. no gaps among transfers;
    * etc.
* Adding additional tests in order to cover edge cases;
* Adding tests for the calculate endpoint;
* Adding proper logging (I would add a separated microservice for logging in order to keep the separation of concerns
  intact)
* Improving naming, e.g. utils, helpers, handlers are just umbrella terms which could hide actual purposes of the code. Naming is supposed to be more descriptive and expressive;
* Adding virtualization layer — i.e. Docker — for developing purposes;
* Replacing all possible magic numbers and magic strings with descriptive constants;
* Properly renaming variables and functions — presently their names are inconsistent;
* Adding handlers for POST-requests in order to record requests to a DB (e.g. PostgreSQL) — it would make the API more interesting. Furthermore, it will give an opportunity to work with data-layer, to implement HEX-architecture, to create models, etc.;