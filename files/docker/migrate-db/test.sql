-- this will create table clients if not exist
-- and insert 1 sample row

CREATE TABLE if NOT EXISTS public.clients (
	id varchar(36) UNIQUE NOT NULL,
	secret varchar(36) NOT NULL,
	CONSTRAINT clients_pkey PRIMARY KEY (id)
);

INSERT INTO public.clients (id, secret)
VALUES ('sample_id','sample_secret');