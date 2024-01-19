-- Database export via SQLPro (https://www.sqlprostudio.com/allapps.html)
-- Exported by e180 at 19-01-2024 19:40.
-- WARNING: This file may contain descructive statements such as DROPs.
-- Please ensure that you are running the script at the proper location.


-- BEGIN TABLE public.phone_numbers
DROP TABLE IF EXISTS public.phone_numbers CASCADE;
BEGIN;

CREATE TABLE IF NOT EXISTS public.phone_numbers (
	id bigint DEFAULT nextval('phone_numbers_id_seq'::regclass) NOT NULL,
	phone_number text,
	provider text,
	"type" bigint,
	created_at timestamp with time zone,
	updated_at timestamp with time zone,
	deleted_at timestamp with time zone,
	PRIMARY KEY(id)
);

COMMIT;

-- END TABLE public.phone_numbers

