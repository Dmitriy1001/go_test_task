package event

type Event struct {
	Id          int64    `db:"id,omitempty"`
	Name        string   `db:"name"`
	Image       string   `db:"image"`
	Description string   `db:"description"`
	Gallery     []string `db:"gallery"`
	Latitude    float64  `db:"latitude"`
	Longitude   float64  `db:"longitude"`
	DateTime    string   `db:"date_time"`
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
