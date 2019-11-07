DROP TABLE IF EXISTS public.samples;

CREATE TABLE public.samples
(
    id serial PRIMARY KEY ,
    sample_id text NOT NULL,
    CONSTRAINT unique_sample UNIQUE (sample_id)

);

ALTER TABLE public.samples
    OWNER to anani;