package event

type Event struct {
	Id          int64    `json:"id,omitempty" db:"id,omitempty"`
	Name        string   `json:"name" db:"name"`
	Image       string   `json:"image" db:"image"`
	Description string   `json:"description" db:"description"`
	Gallery     []string `json:"gallery" db:"gallery"`
	Latitude    float64  `json:"latitude" db:"latitude"`
	Longitude   float64  `json:"longitude" db:"longitude"`
	DateTime    string   `json:"date_time" db:"date_time"`
}

/*
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
*/
