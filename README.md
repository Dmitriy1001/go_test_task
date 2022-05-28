# Test Routes

## Ping

`http://localhost:8081/ping`

## Create Event

`http://localhost:8081/v1/events`

## Events List

`http://localhost:8081/v1/events`

## Event (by ID)

`http://localhost:8081/v1/events/{id}`

## Update Event (by ID)

`http://localhost:8081/v1/events{id}`

## Delete Event (by ID)

`http://localhost:8081/v1/events{id}`

---


### Create table query(postgres)
#
	CREATE TABLE IF NOT EXISTS event
	(
		id integer NOT NULL DEFAULT nextval('event_id_seq'::regclass),
		name character varying(255) NOT NULL,
		image character varying,
		description text NOT NULL,
		gallery character varying[],
		latitude double precision NOT NULL,
		longitude double precision NOT NULL,
		date_time timestamp with time zone NOT NULL,
		CONSTRAINT event_pkey PRIMARY KEY (id)
	)
#

### It is necessary to fill in the fields with the settings variable in main.go
#
    settings := postgresql.ConnectionURL{
		Database: ``,
		Host:     ``,
		User:     ``,
		Password: ``,
	}
#

### POST request example(JSON)
#
    {
        "name": "go conference",
        "image": "http://main_image_link",
        "description": "Big European  conference on Golang",
        "gallery": "http://img_link1, http://img_link2, http://img_link3",
        "latitude": "50.4547",
        "longitude": "30.5238",
        "date_time": "2022-11-28 22:00"
    }
#