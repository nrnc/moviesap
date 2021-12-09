--
-- PostgreSQL database dump
--

-- Dumped from database version 14.1
-- Dumped by pg_dump version 14.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: movies; Type: TABLE; Schema: public; Owner: unbxd
--

CREATE TABLE public.movies (
    id integer NOT NULL,
    title character varying(50) NOT NULL,
    description character varying(255) DEFAULT ''::character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.movies OWNER TO unbxd;

--
-- Name: movies_id_seq; Type: SEQUENCE; Schema: public; Owner: unbxd
--

CREATE SEQUENCE public.movies_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.movies_id_seq OWNER TO unbxd;

--
-- Name: movies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: unbxd
--

ALTER SEQUENCE public.movies_id_seq OWNED BY public.movies.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: unbxd
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO unbxd;

--
-- Name: movies id; Type: DEFAULT; Schema: public; Owner: unbxd
--

ALTER TABLE ONLY public.movies ALTER COLUMN id SET DEFAULT nextval('public.movies_id_seq'::regclass);


--
-- Name: movies movies_pkey; Type: CONSTRAINT; Schema: public; Owner: unbxd
--

ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: unbxd
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

