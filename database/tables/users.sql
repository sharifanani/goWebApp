DROP TABLE IF EXISTS public.users;
CREATE TABLE public.users
(
    id serial PRIMARY KEY ,
    username text NOT NULL,
    pwd bytea NOT NULL,
    email text NOT NULL,

    CONSTRAINT unique_email UNIQUE (email),
    CONSTRAINT unique_username UNIQUE (username)

);
ALTER TABLE public.users
    OWNER to anani;

