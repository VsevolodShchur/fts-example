CREATE TABLE menu_items (
    id serial NOT NULL,
    title varchar(64) NOT NULL,
    description varchar(500) NOT NULL,
	weight     int          NULL,
    weight_measure varchar(16) NOT NULL,

    CONSTRAINT form_pkey PRIMARY KEY (id)
);