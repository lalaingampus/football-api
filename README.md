
# Super Simple REST API Golang Without Database

## Table of contents 👀
* [General info](#general-info)
* [Technologies](#technologies)
* [Setup](#setup)


### General info
 Football-API is a Football REST API made to sho some love for the football.

#### The Football-API Object 🍵
| Properties | Description | Type  |
|:----------- |:---------------|:--------|
|name| the player name & team name | String| 
|description| the description from team | String |
|location| location team | String | 
|nickname| nickname from player |String | 
|position| position player in field | String| 


#### Routes ⚡
| Routes | HTTP Methods| Description
|:------- |:---------------|:--------------
| /Team     | GET                  | Displays all team & player
| /Team      | POST               | Creates a team & player
| /Team      | DELETE            | Deletes all team & player
|/Team/:id| GET     | Displays a specific team & player
|/Team/:id| PUT  | Update player & team value
|/Team/:id| DELETE | Deletes a specific player & team
	
### Technologies
Project is created with:

* Golang
* Gomux

### Setup
To run this project locally, clone repo and add an file in the root:

Then execute in command prompt:
```
$ cd football-api
$ go mod tidy
$ go run main.go
```
## API Reference

These are the endpoints available from the app

### `GET /Team`

Returns list of teams 

#### Response

<details><summary>Show example response</summary>
<p>

```json
{
  "meta": {
    "code": 200
  },
  "data": [
    {
      "id": "5f6a5d6129b2289c40b7444b",
      "name": "AC Milan 2",
      "description": "some-description",
      "location": "Italy",
      "players": [
        {
          "id": "5f6a5d6129b2289c40b74448",
          "name": "John Doe 1",
          "nickname": "Lolo",
          "position": "forward",
          "created_at": "2020-09-22T20:24:01.872Z"
        }
      ],
      "created_at": "2020-09-22T20:24:01.846Z"
    }
  ]
}
```

</p>
</details>

---

### `GET /Team/:id`

Returns a team by id

#### Response

<details><summary>Show example response</summary>
<p>

```json
{
  "meta": {
    "code": 200
  },
  "data": {
    "id": "5f6a5d6129b2289c40b7444b",
    "name": "AC Milan 2",
    "description": "some-description",
    "location": "Italy",
    "players": [
      {
        "id": "5f6a5d6129b2289c40b74448",
        "name": "John Doe 1",
        "nickname": "Lolo",
        "position": "forward",
        "created_at": "2020-09-22T20:24:01.872Z"
      }
    ],
    "created_at": "2020-09-22T20:24:01.846Z"
  }
}
```

</p>
</details>

---

### `POST /Team`

Creates a new team and it's players

#### Request 

This request requires body payload, you can find the example below.

<details><summary>Show example payload</summary>
<p>

```json
{
  "id":"4",
  "name": "AC Milan 2",
  "description": "some-description",
  "location": "Italy",
  "players": [
    {
      "id":"3",
      "name": "John Doe 1",
      "nickname": "Lolo",
      "position": "forward"
    }
  ]
}
```
</p>
</details>

#### Response

The request will return the created data in JSON response like this:

<details><summary>Show example response</summary>
<p>

```json
{
  "meta": {
    "code": 200
  },
  "data": {
    "id": "5f6a5d6129b2289c40b7444b",
    "name": "AC Milan 2",
    "description": "some-description",
    "location": "Italy",
    "players": [
      {
        "id": "5f6a5d6129b2289c40b74448",
        "name": "John Doe 1",
        "nickname": "Lolo",
        "position": "forward",
        "created_at": "2020-09-22T20:24:01.872Z"
      }
    ],
    "created_at": "2020-09-22T20:24:01.846Z"
  }
}
```

</p>
</details>

---

### `GET /players`

Returns list of players 

#### Response

<details><summary>Show example response</summary>
<p>

```json
{
  "meta": {
    "code": 200
  },
  "data": [
    {
      "id": "5f6a5c31d7c451c369802c02",
      "name": "John Doe 1",
      "nickname": "Lolo",
      "position": "forward",
      "created_at": "2020-09-22T20:18:57.957Z"
    }
  ]
}
```

</p>
</details>

---


### `GET /players/:id`

Returns a player by id

#### Response

<details><summary>Show example response</summary>
<p>

```json
{
  "meta": {
    "code": 200
  },
  "data": {
    "id": "5f6a5c31d7c451c369802c02",
    "name": "John Doe 1",
    "nickname": "Lolo",
    "position": "forward",
    "created_at": "2020-09-22T20:18:57.957Z"
  }
}
```

</p>
</details>

---

	
