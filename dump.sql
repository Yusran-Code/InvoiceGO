--
-- PostgreSQL database dump
--

-- Dumped from database version 16.8 (Ubuntu 16.8-0ubuntu0.24.04.1)
-- Dumped by pg_dump version 16.8 (Ubuntu 16.8-0ubuntu0.24.04.1)

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
-- Name: app_profile; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.app_profile (
    id integer NOT NULL,
    nama_pt text NOT NULL,
    nama_bank text NOT NULL,
    no_rekening text NOT NULL,
    penanggung_jawab text NOT NULL,
    alamat text,
    kabupaten text
);


ALTER TABLE public.app_profile OWNER TO postgres;

--
-- Name: app_profile_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.app_profile_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.app_profile_id_seq OWNER TO postgres;

--
-- Name: app_profile_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.app_profile_id_seq OWNED BY public.app_profile.id;


--
-- Name: lo_bulanan; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.lo_bulanan (
    id integer NOT NULL,
    tanggal date NOT NULL,
    no_lo character varying(20) NOT NULL,
    jumlah_tabung integer NOT NULL,
    jumlah_kg integer NOT NULL,
    tarif numeric(10,2) NOT NULL,
    biaya_angkut numeric(10,3) NOT NULL,
    no_so character varying(50)
);


ALTER TABLE public.lo_bulanan OWNER TO postgres;

--
-- Name: lo_bulanan_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.lo_bulanan_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.lo_bulanan_id_seq OWNER TO postgres;

--
-- Name: lo_bulanan_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.lo_bulanan_id_seq OWNED BY public.lo_bulanan.id;


--
-- Name: user_profile; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.user_profile (
    email text NOT NULL,
    nama_pt text,
    nama_bank text,
    no_rekening text,
    penanggung_jawab text,
    alamat text,
    kabupaten text
);


ALTER TABLE public.user_profile OWNER TO postgres;

--
-- Name: app_profile id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.app_profile ALTER COLUMN id SET DEFAULT nextval('public.app_profile_id_seq'::regclass);


--
-- Name: lo_bulanan id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lo_bulanan ALTER COLUMN id SET DEFAULT nextval('public.lo_bulanan_id_seq'::regclass);


--
-- Data for Name: app_profile; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.app_profile (id, nama_pt, nama_bank, no_rekening, penanggung_jawab, alamat, kabupaten) FROM stdin;
5	PT.ANUGERAH HAMZAH PUTRA	MANDIRI-KCP MAKASSAR SULAWESI	152-00-1037507-5	H.M YUSUF	Poros Makassar-Pare, Balombong	Pangkep
\.


--
-- Data for Name: lo_bulanan; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.lo_bulanan (id, tanggal, no_lo, jumlah_tabung, jumlah_kg, tarif, biaya_angkut, no_so) FROM stdin;
\.


--
-- Data for Name: user_profile; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.user_profile (email, nama_pt, nama_bank, no_rekening, penanggung_jawab, alamat, kabupaten) FROM stdin;
yusufhamzah805@gmail.com	PT CONTOH	MANDIRIiii	1234567	BOOSKU	JL MACCCINI RAYA	Makassar
yyuussrraann1992@gmail.com	PT khumaerah	bri	121212	aco	dfdfdfdfd	dfdfdfd
\.


--
-- Name: app_profile_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.app_profile_id_seq', 5, true);


--
-- Name: lo_bulanan_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.lo_bulanan_id_seq', 1, false);


--
-- Name: app_profile app_profile_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.app_profile
    ADD CONSTRAINT app_profile_pkey PRIMARY KEY (id);


--
-- Name: lo_bulanan lo_bulanan_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lo_bulanan
    ADD CONSTRAINT lo_bulanan_pkey PRIMARY KEY (id);


--
-- Name: user_profile user_profile_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.user_profile
    ADD CONSTRAINT user_profile_pkey PRIMARY KEY (email);


--
-- PostgreSQL database dump complete
--

