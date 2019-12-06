DROP TABLE IF EXISTS public.groups;
CREATE TABLE public.groups
(
    id serial PRIMARY KEY ,
    name text NOT NULL,
    data jsonb
);
ALTER TABLE public.groups
    OWNER to anani;
