-- public.album definition

-- Drop table

-- DROP TABLE public.album;

CREATE TABLE public.album (
	id int8 NOT NULL,
	title varchar NULL,
	artist varchar NULL,
	price float8 NULL,
	CONSTRAINT album_pk PRIMARY KEY (id)
);

INSERT INTO public.album (id, title, artist, price) VALUES(1, 'Hari Yang Cerah', 'Peterpan', 50000.0);
INSERT INTO public.album (id, title, artist, price) VALUES(2, 'Sebuah Nama Sebuah Cerita', 'Peterpan', 50000.0);
INSERT INTO public.album (id, title, artist, price) VALUES(3, 'Bintang Di surga', 'Peterpan', 50000.0);
