CREATE TABLE IF NOT EXISTS parameters
(
    id SERIAL NOT NULL ,
    name VARCHAR(256) NOT NULL ,
    value json ,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP ,
    CONSTRAINT parameters_id_pk PRIMARY KEY (id),
    CONSTRAINT parameters_name_uk UNIQUE (name)
);

COMMENT ON TABLE parameters IS 'Serve as a configuration file';