--
-- PostgreSQL database dump
--

-- Dumped from database version 17.0
-- Dumped by pg_dump version 17.0

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: analysis; Type: SCHEMA; Schema: -; Owner: general_user
--

CREATE SCHEMA analysis;


ALTER SCHEMA analysis OWNER TO general_user;

--
-- Name: contribution; Type: SCHEMA; Schema: -; Owner: general_user
--

CREATE SCHEMA contribution;


ALTER SCHEMA contribution OWNER TO general_user;

--
-- Name: developer; Type: SCHEMA; Schema: -; Owner: general_user
--

CREATE SCHEMA developer;


ALTER SCHEMA developer OWNER TO general_user;

--
-- Name: relation; Type: SCHEMA; Schema: -; Owner: general_user
--

CREATE SCHEMA relation;


ALTER SCHEMA relation OWNER TO general_user;

--
-- Name: repo; Type: SCHEMA; Schema: -; Owner: general_user
--

CREATE SCHEMA repo;


ALTER SCHEMA repo OWNER TO general_user;

--
-- Name: update_time(); Type: FUNCTION; Schema: analysis; Owner: general_user
--

CREATE FUNCTION analysis.update_time() RETURNS trigger
    LANGUAGE plpgsql
    AS $$begin    new.data_updated_at= current_timestamp;    return new;end$$;


ALTER FUNCTION analysis.update_time() OWNER TO general_user;

--
-- Name: update_time(); Type: FUNCTION; Schema: contribution; Owner: general_user
--

CREATE FUNCTION contribution.update_time() RETURNS trigger
    LANGUAGE plpgsql
    AS $$begin    new.data_updated_at= current_timestamp;    return new;end$$;


ALTER FUNCTION contribution.update_time() OWNER TO general_user;

--
-- Name: update_time(); Type: FUNCTION; Schema: developer; Owner: general_user
--

CREATE FUNCTION developer.update_time() RETURNS trigger
    LANGUAGE plpgsql
    AS $$begin    new.data_updated_at= current_timestamp;    return new;end$$;


ALTER FUNCTION developer.update_time() OWNER TO general_user;

--
-- Name: update_time(); Type: FUNCTION; Schema: repo; Owner: general_user
--

CREATE FUNCTION repo.update_time() RETURNS trigger
    LANGUAGE plpgsql
    AS $$begin    new.data_updated_at= current_timestamp;    return new;end$$;


ALTER FUNCTION repo.update_time() OWNER TO general_user;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: languages; Type: TABLE; Schema: analysis; Owner: general_user
--

CREATE TABLE analysis.languages (
    data_id bigint NOT NULL,
    data_created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    data_updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    developer_id bigint DEFAULT 0 NOT NULL,
    languages json DEFAULT '{}'::json NOT NULL
);


ALTER TABLE analysis.languages OWNER TO general_user;

--
-- Name: languages_data_id_seq; Type: SEQUENCE; Schema: analysis; Owner: general_user
--

ALTER TABLE analysis.languages ALTER COLUMN data_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME analysis.languages_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: region; Type: TABLE; Schema: analysis; Owner: general_user
--

CREATE TABLE analysis.region (
    data_id bigint NOT NULL,
    data_created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    data_updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    developer_id bigint DEFAULT 0 NOT NULL,
    region character varying(10) DEFAULT ''::character varying NOT NULL,
    confidence double precision DEFAULT 0 NOT NULL
);


ALTER TABLE analysis.region OWNER TO general_user;

--
-- Name: nation_data_id_seq; Type: SEQUENCE; Schema: analysis; Owner: general_user
--

ALTER TABLE analysis.region ALTER COLUMN data_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME analysis.nation_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: pulse_point; Type: TABLE; Schema: analysis; Owner: general_user
--

CREATE TABLE analysis.pulse_point (
    data_id bigint NOT NULL,
    data_created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    data_updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    developer_id bigint DEFAULT 0 NOT NULL,
    pulse_point double precision DEFAULT 0 NOT NULL
);


ALTER TABLE analysis.pulse_point OWNER TO general_user;

--
-- Name: pulse_point_data_id_seq; Type: SEQUENCE; Schema: analysis; Owner: general_user
--

ALTER TABLE analysis.pulse_point ALTER COLUMN data_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME analysis.pulse_point_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: comment_of_user_updated_at; Type: TABLE; Schema: contribution; Owner: general_user
--

CREATE TABLE contribution.comment_of_user_updated_at (
    data_id bigint NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    developer_id bigint DEFAULT 0 NOT NULL
);


ALTER TABLE contribution.comment_of_user_updated_at OWNER TO general_user;

--
-- Name: contribution; Type: TABLE; Schema: contribution; Owner: general_user
--

CREATE TABLE contribution.contribution (
    data_id bigint NOT NULL,
    data_created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    data_updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    user_id bigint DEFAULT 0 NOT NULL,
    repo_id bigint DEFAULT 0 NOT NULL,
    category character varying(20) DEFAULT 'Commit'::character varying NOT NULL,
    content text DEFAULT ''::text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    contribution_id bigint DEFAULT 0 NOT NULL,
    CONSTRAINT check_category CHECK ((((category)::text = 'OpenIssue'::text) OR ((category)::text = 'Comment'::text) OR ((category)::text = 'OpenPullRequest'::text) OR ((category)::text = 'Review'::text) OR ((category)::text = 'Merge'::text)))
);


ALTER TABLE contribution.contribution OWNER TO general_user;

--
-- Name: contribution_data_id_seq; Type: SEQUENCE; Schema: contribution; Owner: general_user
--

ALTER TABLE contribution.contribution ALTER COLUMN data_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME contribution.contribution_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: review_of_user_updated_at; Type: TABLE; Schema: contribution; Owner: general_user
--

CREATE TABLE contribution.review_of_user_updated_at (
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    data_id bigint NOT NULL,
    developer_id bigint DEFAULT 0 NOT NULL
);


ALTER TABLE contribution.review_of_user_updated_at OWNER TO general_user;

--
-- Name: contribution_update_at_2_data_id_seq; Type: SEQUENCE; Schema: contribution; Owner: general_user
--

ALTER TABLE contribution.review_of_user_updated_at ALTER COLUMN data_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME contribution.contribution_update_at_2_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: contribution_update_at_3_data_id_seq; Type: SEQUENCE; Schema: contribution; Owner: general_user
--

ALTER TABLE contribution.comment_of_user_updated_at ALTER COLUMN data_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME contribution.contribution_update_at_3_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: issue_pr_of_user_updated_at; Type: TABLE; Schema: contribution; Owner: general_user
--

CREATE TABLE contribution.issue_pr_of_user_updated_at (
    data_id bigint NOT NULL,
    developer_id bigint DEFAULT 0 NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE contribution.issue_pr_of_user_updated_at OWNER TO general_user;

--
-- Name: contribution_update_at_data_id_seq; Type: SEQUENCE; Schema: contribution; Owner: general_user
--

ALTER TABLE contribution.issue_pr_of_user_updated_at ALTER COLUMN data_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME contribution.contribution_update_at_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: developer; Type: TABLE; Schema: developer; Owner: general_user
--

CREATE TABLE developer.developer (
    data_id bigint NOT NULL,
    data_created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    data_updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    id bigint DEFAULT 0 NOT NULL,
    name character varying(255) DEFAULT ''::character varying NOT NULL,
    login character varying(255) DEFAULT ''::character varying NOT NULL,
    avatar_url character varying(255) DEFAULT ''::character varying NOT NULL,
    company character varying(255) DEFAULT ''::character varying NOT NULL,
    location character varying(255) DEFAULT ''::character varying NOT NULL,
    bio character varying(255) DEFAULT ''::character varying NOT NULL,
    blog character varying(255) DEFAULT ''::character varying NOT NULL,
    email character varying(255) DEFAULT ''::character varying NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    twitter_username character varying(255) DEFAULT ''::character varying NOT NULL,
    repos bigint DEFAULT 0 NOT NULL,
    following bigint DEFAULT 0 NOT NULL,
    followers bigint DEFAULT 0 NOT NULL,
    gists bigint DEFAULT 0 NOT NULL,
    stars bigint DEFAULT 0 NOT NULL
);


ALTER TABLE developer.developer OWNER TO general_user;

--
-- Name: developer_data_id_seq; Type: SEQUENCE; Schema: developer; Owner: general_user
--

ALTER TABLE developer.developer ALTER COLUMN data_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME developer.developer_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: create_repo; Type: TABLE; Schema: relation; Owner: general_user
--

CREATE TABLE relation.create_repo (
    data_id bigint NOT NULL,
    developer_id bigint DEFAULT 0 NOT NULL,
    repo_id bigint DEFAULT 0 NOT NULL
);


ALTER TABLE relation.create_repo OWNER TO general_user;

--
-- Name: create_repo_data_id_seq; Type: SEQUENCE; Schema: relation; Owner: general_user
--

ALTER TABLE relation.create_repo ALTER COLUMN data_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME relation.create_repo_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: created_repo_updated_at; Type: TABLE; Schema: relation; Owner: general_user
--

CREATE TABLE relation.created_repo_updated_at (
    data_id bigint NOT NULL,
    developer_id bigint DEFAULT 0 NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE relation.created_repo_updated_at OWNER TO general_user;

--
-- Name: create_repo_update_at_data_id_seq; Type: SEQUENCE; Schema: relation; Owner: general_user
--

ALTER TABLE relation.created_repo_updated_at ALTER COLUMN data_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME relation.create_repo_update_at_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: follow; Type: TABLE; Schema: relation; Owner: general_user
--

CREATE TABLE relation.follow (
    data_id bigint NOT NULL,
    follower_id bigint DEFAULT 0 NOT NULL,
    following_id bigint DEFAULT 0 NOT NULL
);


ALTER TABLE relation.follow OWNER TO general_user;

--
-- Name: follow_data_id_seq; Type: SEQUENCE; Schema: relation; Owner: general_user
--

ALTER TABLE relation.follow ALTER COLUMN data_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME relation.follow_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: following_updated_at; Type: TABLE; Schema: relation; Owner: general_user
--

CREATE TABLE relation.following_updated_at (
    data_id bigint NOT NULL,
    developer_id bigint DEFAULT 0 NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE relation.following_updated_at OWNER TO general_user;

--
-- Name: follow_update_at_data_id_seq; Type: SEQUENCE; Schema: relation; Owner: general_user
--

ALTER TABLE relation.following_updated_at ALTER COLUMN data_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME relation.follow_update_at_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: follower_updated_at; Type: TABLE; Schema: relation; Owner: general_user
--

CREATE TABLE relation.follower_updated_at (
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    data_id bigint NOT NULL,
    developer_id bigint DEFAULT 0 NOT NULL
);


ALTER TABLE relation.follower_updated_at OWNER TO general_user;

--
-- Name: follower_update_at_data_id_seq; Type: SEQUENCE; Schema: relation; Owner: general_user
--

ALTER TABLE relation.follower_updated_at ALTER COLUMN data_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME relation.follower_update_at_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: fork; Type: TABLE; Schema: relation; Owner: general_user
--

CREATE TABLE relation.fork (
    data_id bigint NOT NULL,
    original_repo_id bigint DEFAULT 0 NOT NULL,
    fork_repo_id bigint DEFAULT 0 NOT NULL
);


ALTER TABLE relation.fork OWNER TO general_user;

--
-- Name: fork_data_id_seq; Type: SEQUENCE; Schema: relation; Owner: general_user
--

ALTER TABLE relation.fork ALTER COLUMN data_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME relation.fork_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: fork_updated_at; Type: TABLE; Schema: relation; Owner: general_user
--

CREATE TABLE relation.fork_updated_at (
    data_id bigint NOT NULL,
    repo_id bigint DEFAULT 0 NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE relation.fork_updated_at OWNER TO general_user;

--
-- Name: fork_update_at_data_id_seq; Type: SEQUENCE; Schema: relation; Owner: general_user
--

ALTER TABLE relation.fork_updated_at ALTER COLUMN data_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME relation.fork_update_at_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: star; Type: TABLE; Schema: relation; Owner: general_user
--

CREATE TABLE relation.star (
    data_id bigint NOT NULL,
    developer_id bigint DEFAULT 0 NOT NULL,
    repo_id bigint DEFAULT 0 NOT NULL
);


ALTER TABLE relation.star OWNER TO general_user;

--
-- Name: star_data_id_seq; Type: SEQUENCE; Schema: relation; Owner: general_user
--

ALTER TABLE relation.star ALTER COLUMN data_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME relation.star_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: starred_repo_updated_at; Type: TABLE; Schema: relation; Owner: general_user
--

CREATE TABLE relation.starred_repo_updated_at (
    data_id bigint NOT NULL,
    developer_id bigint DEFAULT 0 NOT NULL,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE relation.starred_repo_updated_at OWNER TO general_user;

--
-- Name: star_update_at_data_id_seq; Type: SEQUENCE; Schema: relation; Owner: general_user
--

ALTER TABLE relation.starred_repo_updated_at ALTER COLUMN data_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME relation.star_update_at_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: repo; Type: TABLE; Schema: repo; Owner: general_user
--

CREATE TABLE repo.repo (
    data_id bigint NOT NULL,
    data_created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    data_updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    id bigint DEFAULT 0 NOT NULL,
    name character varying(255) DEFAULT ''::character varying NOT NULL,
    star_count bigint DEFAULT 0 NOT NULL,
    fork_count bigint DEFAULT 0 NOT NULL,
    issue_count bigint DEFAULT 0 NOT NULL,
    commit_count bigint DEFAULT 0 NOT NULL,
    language json DEFAULT '{}'::json NOT NULL,
    description character varying(255) DEFAULT ''::character varying NOT NULL,
    last_fetch_fork_at timestamp with time zone DEFAULT to_timestamp((0)::double precision) NOT NULL,
    last_fetch_contribution_at timestamp with time zone DEFAULT to_timestamp((0)::double precision) NOT NULL,
    merged_pr_count bigint DEFAULT 0 NOT NULL,
    open_pr_count bigint DEFAULT 0 NOT NULL,
    comment_count bigint DEFAULT 0 NOT NULL,
    review_count bigint DEFAULT 0 NOT NULL,
    pr_count bigint DEFAULT 0 NOT NULL
);


ALTER TABLE repo.repo OWNER TO general_user;

--
-- Name: repo_data_id_seq; Type: SEQUENCE; Schema: repo; Owner: general_user
--

ALTER TABLE repo.repo ALTER COLUMN data_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME repo.repo_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Data for Name: languages; Type: TABLE DATA; Schema: analysis; Owner: general_user
--

COPY analysis.languages (data_id, data_created_at, data_updated_at, developer_id, languages) FROM stdin;
\.


--
-- Data for Name: pulse_point; Type: TABLE DATA; Schema: analysis; Owner: general_user
--

COPY analysis.pulse_point (data_id, data_created_at, data_updated_at, developer_id, pulse_point) FROM stdin;
\.


--
-- Data for Name: region; Type: TABLE DATA; Schema: analysis; Owner: general_user
--

COPY analysis.region (data_id, data_created_at, data_updated_at, developer_id, region, confidence) FROM stdin;
\.


--
-- Data for Name: comment_of_user_updated_at; Type: TABLE DATA; Schema: contribution; Owner: general_user
--

COPY contribution.comment_of_user_updated_at (data_id, updated_at, developer_id) FROM stdin;
\.


--
-- Data for Name: contribution; Type: TABLE DATA; Schema: contribution; Owner: general_user
--

COPY contribution.contribution (data_id, data_created_at, data_updated_at, user_id, repo_id, category, content, created_at, updated_at, contribution_id) FROM stdin;
\.


--
-- Data for Name: issue_pr_of_user_updated_at; Type: TABLE DATA; Schema: contribution; Owner: general_user
--

COPY contribution.issue_pr_of_user_updated_at (data_id, developer_id, updated_at) FROM stdin;
\.


--
-- Data for Name: review_of_user_updated_at; Type: TABLE DATA; Schema: contribution; Owner: general_user
--

COPY contribution.review_of_user_updated_at (updated_at, data_id, developer_id) FROM stdin;
\.


--
-- Data for Name: developer; Type: TABLE DATA; Schema: developer; Owner: general_user
--

COPY developer.developer (data_id, data_created_at, data_updated_at, id, name, login, avatar_url, company, location, bio, blog, email, created_at, updated_at, twitter_username, repos, following, followers, gists, stars) FROM stdin;
\.


--
-- Data for Name: create_repo; Type: TABLE DATA; Schema: relation; Owner: general_user
--

COPY relation.create_repo (data_id, developer_id, repo_id) FROM stdin;
\.


--
-- Data for Name: created_repo_updated_at; Type: TABLE DATA; Schema: relation; Owner: general_user
--

COPY relation.created_repo_updated_at (data_id, developer_id, updated_at) FROM stdin;
\.


--
-- Data for Name: follow; Type: TABLE DATA; Schema: relation; Owner: general_user
--

COPY relation.follow (data_id, follower_id, following_id) FROM stdin;
\.


--
-- Data for Name: follower_updated_at; Type: TABLE DATA; Schema: relation; Owner: general_user
--

COPY relation.follower_updated_at (updated_at, data_id, developer_id) FROM stdin;
\.


--
-- Data for Name: following_updated_at; Type: TABLE DATA; Schema: relation; Owner: general_user
--

COPY relation.following_updated_at (data_id, developer_id, updated_at) FROM stdin;
\.


--
-- Data for Name: fork; Type: TABLE DATA; Schema: relation; Owner: general_user
--

COPY relation.fork (data_id, original_repo_id, fork_repo_id) FROM stdin;
\.


--
-- Data for Name: fork_updated_at; Type: TABLE DATA; Schema: relation; Owner: general_user
--

COPY relation.fork_updated_at (data_id, repo_id, updated_at) FROM stdin;
\.


--
-- Data for Name: star; Type: TABLE DATA; Schema: relation; Owner: general_user
--

COPY relation.star (data_id, developer_id, repo_id) FROM stdin;
\.


--
-- Data for Name: starred_repo_updated_at; Type: TABLE DATA; Schema: relation; Owner: general_user
--

COPY relation.starred_repo_updated_at (data_id, developer_id, updated_at) FROM stdin;
\.


--
-- Data for Name: repo; Type: TABLE DATA; Schema: repo; Owner: general_user
--

COPY repo.repo (data_id, data_created_at, data_updated_at, id, name, star_count, fork_count, issue_count, commit_count, language, description, last_fetch_fork_at, last_fetch_contribution_at, merged_pr_count, open_pr_count, comment_count, review_count, pr_count) FROM stdin;
\.


--
-- Name: languages_data_id_seq; Type: SEQUENCE SET; Schema: analysis; Owner: general_user
--

SELECT pg_catalog.setval('analysis.languages_data_id_seq', 15, true);


--
-- Name: nation_data_id_seq; Type: SEQUENCE SET; Schema: analysis; Owner: general_user
--

SELECT pg_catalog.setval('analysis.nation_data_id_seq', 59, true);


--
-- Name: pulse_point_data_id_seq; Type: SEQUENCE SET; Schema: analysis; Owner: general_user
--

SELECT pg_catalog.setval('analysis.pulse_point_data_id_seq', 27, true);


--
-- Name: contribution_data_id_seq; Type: SEQUENCE SET; Schema: contribution; Owner: general_user
--

SELECT pg_catalog.setval('contribution.contribution_data_id_seq', 2113, true);


--
-- Name: contribution_update_at_2_data_id_seq; Type: SEQUENCE SET; Schema: contribution; Owner: general_user
--

SELECT pg_catalog.setval('contribution.contribution_update_at_2_data_id_seq', 2, true);


--
-- Name: contribution_update_at_3_data_id_seq; Type: SEQUENCE SET; Schema: contribution; Owner: general_user
--

SELECT pg_catalog.setval('contribution.contribution_update_at_3_data_id_seq', 2, true);


--
-- Name: contribution_update_at_data_id_seq; Type: SEQUENCE SET; Schema: contribution; Owner: general_user
--

SELECT pg_catalog.setval('contribution.contribution_update_at_data_id_seq', 2, true);


--
-- Name: developer_data_id_seq; Type: SEQUENCE SET; Schema: developer; Owner: general_user
--

SELECT pg_catalog.setval('developer.developer_data_id_seq', 134, true);


--
-- Name: create_repo_data_id_seq; Type: SEQUENCE SET; Schema: relation; Owner: general_user
--

SELECT pg_catalog.setval('relation.create_repo_data_id_seq', 270, true);


--
-- Name: create_repo_update_at_data_id_seq; Type: SEQUENCE SET; Schema: relation; Owner: general_user
--

SELECT pg_catalog.setval('relation.create_repo_update_at_data_id_seq', 4, true);


--
-- Name: follow_data_id_seq; Type: SEQUENCE SET; Schema: relation; Owner: general_user
--

SELECT pg_catalog.setval('relation.follow_data_id_seq', 63, true);


--
-- Name: follow_update_at_data_id_seq; Type: SEQUENCE SET; Schema: relation; Owner: general_user
--

SELECT pg_catalog.setval('relation.follow_update_at_data_id_seq', 1, false);


--
-- Name: follower_update_at_data_id_seq; Type: SEQUENCE SET; Schema: relation; Owner: general_user
--

SELECT pg_catalog.setval('relation.follower_update_at_data_id_seq', 1, false);


--
-- Name: fork_data_id_seq; Type: SEQUENCE SET; Schema: relation; Owner: general_user
--

SELECT pg_catalog.setval('relation.fork_data_id_seq', 13, true);


--
-- Name: fork_update_at_data_id_seq; Type: SEQUENCE SET; Schema: relation; Owner: general_user
--

SELECT pg_catalog.setval('relation.fork_update_at_data_id_seq', 1, false);


--
-- Name: star_data_id_seq; Type: SEQUENCE SET; Schema: relation; Owner: general_user
--

SELECT pg_catalog.setval('relation.star_data_id_seq', 55, true);


--
-- Name: star_update_at_data_id_seq; Type: SEQUENCE SET; Schema: relation; Owner: general_user
--

SELECT pg_catalog.setval('relation.star_update_at_data_id_seq', 1, false);


--
-- Name: repo_data_id_seq; Type: SEQUENCE SET; Schema: repo; Owner: general_user
--

SELECT pg_catalog.setval('repo.repo_data_id_seq', 411, true);


--
-- Name: languages languages_pk; Type: CONSTRAINT; Schema: analysis; Owner: general_user
--

ALTER TABLE ONLY analysis.languages
    ADD CONSTRAINT languages_pk PRIMARY KEY (data_id);


--
-- Name: languages languages_pk_2; Type: CONSTRAINT; Schema: analysis; Owner: general_user
--

ALTER TABLE ONLY analysis.languages
    ADD CONSTRAINT languages_pk_2 UNIQUE (developer_id);


--
-- Name: pulse_point pulse_point_pk; Type: CONSTRAINT; Schema: analysis; Owner: general_user
--

ALTER TABLE ONLY analysis.pulse_point
    ADD CONSTRAINT pulse_point_pk PRIMARY KEY (data_id);


--
-- Name: pulse_point pulse_point_pk_2; Type: CONSTRAINT; Schema: analysis; Owner: general_user
--

ALTER TABLE ONLY analysis.pulse_point
    ADD CONSTRAINT pulse_point_pk_2 UNIQUE (developer_id);


--
-- Name: region region_pk; Type: CONSTRAINT; Schema: analysis; Owner: general_user
--

ALTER TABLE ONLY analysis.region
    ADD CONSTRAINT region_pk PRIMARY KEY (data_id);


--
-- Name: region region_pk_2; Type: CONSTRAINT; Schema: analysis; Owner: general_user
--

ALTER TABLE ONLY analysis.region
    ADD CONSTRAINT region_pk_2 UNIQUE (developer_id);


--
-- Name: comment_of_user_updated_at comment_of_user_update_at_pk; Type: CONSTRAINT; Schema: contribution; Owner: general_user
--

ALTER TABLE ONLY contribution.comment_of_user_updated_at
    ADD CONSTRAINT comment_of_user_update_at_pk PRIMARY KEY (data_id);


--
-- Name: comment_of_user_updated_at comment_of_user_update_at_pk_2; Type: CONSTRAINT; Schema: contribution; Owner: general_user
--

ALTER TABLE ONLY contribution.comment_of_user_updated_at
    ADD CONSTRAINT comment_of_user_update_at_pk_2 UNIQUE (developer_id);


--
-- Name: contribution contribution_pk; Type: CONSTRAINT; Schema: contribution; Owner: general_user
--

ALTER TABLE ONLY contribution.contribution
    ADD CONSTRAINT contribution_pk PRIMARY KEY (data_id);


--
-- Name: contribution contribution_pk_2; Type: CONSTRAINT; Schema: contribution; Owner: general_user
--

ALTER TABLE ONLY contribution.contribution
    ADD CONSTRAINT contribution_pk_2 UNIQUE (category, repo_id, contribution_id);


--
-- Name: issue_pr_of_user_updated_at issue_pr_of_user_update_at_pk; Type: CONSTRAINT; Schema: contribution; Owner: general_user
--

ALTER TABLE ONLY contribution.issue_pr_of_user_updated_at
    ADD CONSTRAINT issue_pr_of_user_update_at_pk PRIMARY KEY (data_id);


--
-- Name: issue_pr_of_user_updated_at issue_pr_of_user_update_at_pk_2; Type: CONSTRAINT; Schema: contribution; Owner: general_user
--

ALTER TABLE ONLY contribution.issue_pr_of_user_updated_at
    ADD CONSTRAINT issue_pr_of_user_update_at_pk_2 UNIQUE (developer_id);


--
-- Name: review_of_user_updated_at review_of_user_update_at_pk; Type: CONSTRAINT; Schema: contribution; Owner: general_user
--

ALTER TABLE ONLY contribution.review_of_user_updated_at
    ADD CONSTRAINT review_of_user_update_at_pk PRIMARY KEY (data_id);


--
-- Name: review_of_user_updated_at review_of_user_update_at_pk_2; Type: CONSTRAINT; Schema: contribution; Owner: general_user
--

ALTER TABLE ONLY contribution.review_of_user_updated_at
    ADD CONSTRAINT review_of_user_update_at_pk_2 UNIQUE (developer_id);


--
-- Name: developer developer_pk; Type: CONSTRAINT; Schema: developer; Owner: general_user
--

ALTER TABLE ONLY developer.developer
    ADD CONSTRAINT developer_pk PRIMARY KEY (data_id);


--
-- Name: developer developer_pk_2; Type: CONSTRAINT; Schema: developer; Owner: general_user
--

ALTER TABLE ONLY developer.developer
    ADD CONSTRAINT developer_pk_2 UNIQUE (id);


--
-- Name: developer developer_pk_3; Type: CONSTRAINT; Schema: developer; Owner: general_user
--

ALTER TABLE ONLY developer.developer
    ADD CONSTRAINT developer_pk_3 UNIQUE (login);


--
-- Name: create_repo create_repo_pk; Type: CONSTRAINT; Schema: relation; Owner: general_user
--

ALTER TABLE ONLY relation.create_repo
    ADD CONSTRAINT create_repo_pk PRIMARY KEY (data_id);


--
-- Name: create_repo create_repo_pk_2; Type: CONSTRAINT; Schema: relation; Owner: general_user
--

ALTER TABLE ONLY relation.create_repo
    ADD CONSTRAINT create_repo_pk_2 UNIQUE (repo_id);


--
-- Name: created_repo_updated_at create_repo_update_at_pk; Type: CONSTRAINT; Schema: relation; Owner: general_user
--

ALTER TABLE ONLY relation.created_repo_updated_at
    ADD CONSTRAINT create_repo_update_at_pk PRIMARY KEY (data_id);


--
-- Name: created_repo_updated_at create_repo_update_at_pk_2; Type: CONSTRAINT; Schema: relation; Owner: general_user
--

ALTER TABLE ONLY relation.created_repo_updated_at
    ADD CONSTRAINT create_repo_update_at_pk_2 UNIQUE (developer_id);


--
-- Name: follow follow_pk; Type: CONSTRAINT; Schema: relation; Owner: general_user
--

ALTER TABLE ONLY relation.follow
    ADD CONSTRAINT follow_pk PRIMARY KEY (data_id);


--
-- Name: follow follow_pk_2; Type: CONSTRAINT; Schema: relation; Owner: general_user
--

ALTER TABLE ONLY relation.follow
    ADD CONSTRAINT follow_pk_2 UNIQUE (follower_id, following_id);


--
-- Name: following_updated_at follow_update_at_pk; Type: CONSTRAINT; Schema: relation; Owner: general_user
--

ALTER TABLE ONLY relation.following_updated_at
    ADD CONSTRAINT follow_update_at_pk PRIMARY KEY (data_id);


--
-- Name: following_updated_at follow_update_at_pk_2; Type: CONSTRAINT; Schema: relation; Owner: general_user
--

ALTER TABLE ONLY relation.following_updated_at
    ADD CONSTRAINT follow_update_at_pk_2 UNIQUE (developer_id);


--
-- Name: follower_updated_at follower_update_at_pk; Type: CONSTRAINT; Schema: relation; Owner: general_user
--

ALTER TABLE ONLY relation.follower_updated_at
    ADD CONSTRAINT follower_update_at_pk PRIMARY KEY (data_id);


--
-- Name: follower_updated_at follower_update_at_pk_2; Type: CONSTRAINT; Schema: relation; Owner: general_user
--

ALTER TABLE ONLY relation.follower_updated_at
    ADD CONSTRAINT follower_update_at_pk_2 UNIQUE (developer_id);


--
-- Name: fork fork_pk; Type: CONSTRAINT; Schema: relation; Owner: general_user
--

ALTER TABLE ONLY relation.fork
    ADD CONSTRAINT fork_pk PRIMARY KEY (data_id);


--
-- Name: fork fork_pk_2; Type: CONSTRAINT; Schema: relation; Owner: general_user
--

ALTER TABLE ONLY relation.fork
    ADD CONSTRAINT fork_pk_2 UNIQUE (fork_repo_id);


--
-- Name: fork_updated_at fork_update_at_pk; Type: CONSTRAINT; Schema: relation; Owner: general_user
--

ALTER TABLE ONLY relation.fork_updated_at
    ADD CONSTRAINT fork_update_at_pk PRIMARY KEY (data_id);


--
-- Name: fork_updated_at fork_update_at_pk_2; Type: CONSTRAINT; Schema: relation; Owner: general_user
--

ALTER TABLE ONLY relation.fork_updated_at
    ADD CONSTRAINT fork_update_at_pk_2 UNIQUE (repo_id);


--
-- Name: star star_pk; Type: CONSTRAINT; Schema: relation; Owner: general_user
--

ALTER TABLE ONLY relation.star
    ADD CONSTRAINT star_pk PRIMARY KEY (data_id);


--
-- Name: star star_pk_2; Type: CONSTRAINT; Schema: relation; Owner: general_user
--

ALTER TABLE ONLY relation.star
    ADD CONSTRAINT star_pk_2 UNIQUE (developer_id, repo_id);


--
-- Name: starred_repo_updated_at star_update_at_pk; Type: CONSTRAINT; Schema: relation; Owner: general_user
--

ALTER TABLE ONLY relation.starred_repo_updated_at
    ADD CONSTRAINT star_update_at_pk PRIMARY KEY (data_id);


--
-- Name: starred_repo_updated_at star_update_at_pk_2; Type: CONSTRAINT; Schema: relation; Owner: general_user
--

ALTER TABLE ONLY relation.starred_repo_updated_at
    ADD CONSTRAINT star_update_at_pk_2 UNIQUE (developer_id);


--
-- Name: repo repo_pk; Type: CONSTRAINT; Schema: repo; Owner: general_user
--

ALTER TABLE ONLY repo.repo
    ADD CONSTRAINT repo_pk PRIMARY KEY (data_id);


--
-- Name: repo repo_pk_2; Type: CONSTRAINT; Schema: repo; Owner: general_user
--

ALTER TABLE ONLY repo.repo
    ADD CONSTRAINT repo_pk_2 UNIQUE (id);


--
-- Name: languages update_time; Type: TRIGGER; Schema: analysis; Owner: general_user
--

CREATE TRIGGER update_time BEFORE UPDATE ON analysis.languages FOR EACH ROW EXECUTE FUNCTION analysis.update_time();


--
-- Name: pulse_point update_time; Type: TRIGGER; Schema: analysis; Owner: general_user
--

CREATE TRIGGER update_time BEFORE UPDATE ON analysis.pulse_point FOR EACH ROW EXECUTE FUNCTION analysis.update_time();


--
-- Name: region update_time; Type: TRIGGER; Schema: analysis; Owner: general_user
--

CREATE TRIGGER update_time BEFORE UPDATE ON analysis.region FOR EACH ROW EXECUTE FUNCTION analysis.update_time();


--
-- Name: contribution update_time; Type: TRIGGER; Schema: contribution; Owner: general_user
--

CREATE TRIGGER update_time BEFORE UPDATE ON contribution.contribution FOR EACH ROW EXECUTE FUNCTION contribution.update_time();


--
-- Name: developer update_time; Type: TRIGGER; Schema: developer; Owner: general_user
--

CREATE TRIGGER update_time BEFORE UPDATE ON developer.developer FOR EACH ROW EXECUTE FUNCTION developer.update_time();


--
-- Name: repo update_time; Type: TRIGGER; Schema: repo; Owner: general_user
--

CREATE TRIGGER update_time BEFORE UPDATE ON repo.repo FOR EACH ROW EXECUTE FUNCTION repo.update_time();


--
-- PostgreSQL database dump complete
--

