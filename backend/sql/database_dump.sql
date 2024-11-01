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
-- Name: analysis; Type: TABLE; Schema: analysis; Owner: general_user
--

CREATE TABLE analysis.analysis (
    data_id bigint NOT NULL,
    data_created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    data_updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    developer_id bigint DEFAULT 0 NOT NULL,
    languages json DEFAULT '{}'::json NOT NULL,
    talent_rank double precision DEFAULT 0 NOT NULL,
    nation character varying(255) DEFAULT ''::character varying NOT NULL,
    pulse_point double precision DEFAULT 0 NOT NULL
);


ALTER TABLE analysis.analysis OWNER TO general_user;

--
-- Name: analysis_data_id_seq; Type: SEQUENCE; Schema: analysis; Owner: general_user
--

ALTER TABLE analysis.analysis ALTER COLUMN data_id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME analysis.analysis_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


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
    stars bigint DEFAULT 0 NOT NULL,
    last_fetch_create_repo_at timestamp with time zone DEFAULT to_timestamp((0)::double precision) NOT NULL,
    last_fetch_follow_at timestamp with time zone DEFAULT to_timestamp((0)::double precision) NOT NULL,
    last_fetch_star_at timestamp with time zone DEFAULT to_timestamp((0)::double precision) NOT NULL,
    last_fetch_contribution_at timestamp with time zone DEFAULT to_timestamp((0)::double precision) NOT NULL
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
    pr_count bigint DEFAULT 0 NOT NULL,
    language json DEFAULT '{}'::json NOT NULL,
    description character varying(255) DEFAULT ''::character varying NOT NULL,
    last_fetch_fork_at timestamp with time zone DEFAULT to_timestamp((0)::double precision) NOT NULL,
    last_fetch_contribution_at timestamp with time zone DEFAULT to_timestamp((0)::double precision) NOT NULL
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
-- Data for Name: analysis; Type: TABLE DATA; Schema: analysis; Owner: general_user
--

COPY analysis.analysis (data_id, data_created_at, data_updated_at, developer_id, languages, talent_rank, nation, pulse_point) FROM stdin;
1	2024-10-27 23:05:03.138255+08	2024-10-27 23:05:03.138257+08	-5859999829173043	{}	90	adipisicing proident quis esse elit	0
3	2024-10-27 23:05:49.482043+08	2024-10-27 23:08:29.323881+08	7479059541506461	{}	4	aute	0
\.


--
-- Data for Name: contribution; Type: TABLE DATA; Schema: contribution; Owner: general_user
--

COPY contribution.contribution (data_id, data_created_at, data_updated_at, user_id, repo_id, category, content, created_at, updated_at, contribution_id) FROM stdin;
5	2024-10-27 17:47:03.197869+08	2024-10-27 17:47:03.197873+08	-1228728364880275	2933813597660321	OpenIssue	in	2024-10-27 17:47:03.198096+08	2024-10-27 17:47:03.198096+08	89132188865681
6	2024-10-27 17:47:08.030081+08	2024-10-27 17:47:08.030083+08	2181321206809377	-2851151441856547	OpenIssue	ea laboris	2024-10-27 17:47:08.030295+08	2024-10-27 17:47:08.030295+08	7886607944603321
7	2024-10-27 17:47:56.179865+08	2024-10-27 17:47:56.179868+08	2181321206809377	-2851151441856547	OpenPullRequest	ea laboris	2024-10-27 17:47:56.183734+08	2024-10-27 17:47:56.183734+08	7886607944603321
9	2024-10-27 17:48:04.462636+08	2024-10-27 17:48:04.462638+08	2424484547017825	240073372326149	OpenPullRequest	sunt cillum	2024-10-27 17:48:04.462873+08	2024-10-27 17:48:04.462873+08	-1005542944539355
10	2024-10-27 17:48:07.809511+08	2024-10-27 17:48:07.809514+08	2130166492783173	3703182763939297	OpenPullRequest	in nulla labore sed reprehenderit	2024-10-27 17:48:07.80974+08	2024-10-27 17:48:07.80974+08	3121030929766249
11	2024-10-27 17:48:12.66339+08	2024-10-27 17:52:35.350832+08	-6397966326814175	-4208338970191259	OpenPullRequest	utsadsadasdasdasdasdasdaidunt	2024-10-27 17:48:12.663633+08	2024-10-27 17:48:12.663633+08	-11410223958999
2	2024-10-27 17:46:41.436558+08	2024-11-01 12:55:02.665284+08	7243146715757265	-5982590624140767	OpenIssue	laboris	2024-10-27 17:46:41.436819+08	2024-10-27 17:46:41.436819+08	5411042840125525
4	2024-10-27 17:46:55.877717+08	2024-11-01 12:55:02.665284+08	-6925608478524259	200942203100505	OpenIssue	cillum	2024-10-27 17:46:55.877953+08	2024-10-27 17:46:55.877953+08	-2959791549189611
3	2024-10-27 17:46:52.205534+08	2024-11-01 12:55:02.665284+08	7456789445831205	1613420073909173	OpenIssue	laboris tempor	2024-10-27 17:46:52.205762+08	2024-10-27 17:46:52.205762+08	7736778005277345
12	2024-10-27 17:53:01.380784+08	2024-11-01 12:55:02.665284+08	123123123	0	OpenIssue	ut sint incididunt	2024-10-27 17:53:01.381249+08	2024-10-27 17:53:01.381249+08	0
\.


--
-- Data for Name: developer; Type: TABLE DATA; Schema: developer; Owner: general_user
--

COPY developer.developer (data_id, data_created_at, data_updated_at, id, name, login, avatar_url, company, location, bio, blog, email, created_at, updated_at, twitter_username, repos, following, followers, gists, stars, last_fetch_create_repo_at, last_fetch_follow_at, last_fetch_star_at, last_fetch_contribution_at) FROM stdin;
14	2024-10-27 12:42:46.091326+08	2024-10-27 14:19:04.443238+08	0	dsada	12							2024-10-27 12:42:46.091326+08	2024-10-27 12:42:46.091326+08		0	0	0	0	0	2024-10-31 10:52:01.101542+08	2024-10-31 10:52:01.112956+08	2024-10-31 10:52:01.119971+08	1970-01-01 08:00:00+08
16	2024-10-27 14:24:12.236316+08	2024-10-27 14:24:12.236318+08	3510714200528412	茹乙萍	籍婷方	https://avatars.githubusercontent.com/u/12477906	et ad tempor	dolor eu sed sit esse	顾问爱好者，工程师😆	laborum aliqua	rn1hdu_px57@yahoo.com.cn	2024-10-27 14:24:12.236575+08	2024-10-27 14:24:12.236575+08		0	0	0	0	0	2024-10-31 10:52:01.101542+08	2024-10-31 10:52:01.112956+08	2024-10-31 10:52:01.119971+08	1970-01-01 08:00:00+08
17	2024-10-27 14:24:14.276458+08	2024-10-27 14:24:14.27646+08	6506382737603632	冯思佳	揭艺涵	https://avatars.githubusercontent.com/u/36778824	nulla ullamco	quis dolor proident labore	极客，摄影爱好者，领导者	Duis amet cillum et velit	i1csbt93@sohu.com	2024-10-27 14:24:14.276702+08	2024-10-27 14:24:14.276702+08		0	0	0	0	0	2024-10-31 10:52:01.101542+08	2024-10-31 10:52:01.112956+08	2024-10-31 10:52:01.119971+08	1970-01-01 08:00:00+08
18	2024-10-27 14:24:16.159932+08	2024-10-27 14:24:16.159935+08	2321693143446954	介雨欣	赖静	https://avatars.githubusercontent.com/u/85038730	ut cupidatat dolor Lorem	dolor	公众演说家，哲学家	Duis enim proident dolor	k7nggq89@126.com	2024-10-27 14:24:16.160167+08	2024-10-27 14:24:16.160167+08		0	0	0	0	0	2024-10-31 10:52:01.101542+08	2024-10-31 10:52:01.112956+08	2024-10-31 10:52:01.119971+08	1970-01-01 08:00:00+08
19	2024-10-27 14:24:17.902788+08	2024-10-27 14:24:17.90279+08	5999171471706830	其宇航	瞿国香	https://avatars.githubusercontent.com/u/4692930	eu occaecat in sed	deserunt quis tempor	附属贡献者🚩	dolor amet	fmnrmi50@sina.com	2024-10-27 14:24:17.903022+08	2024-10-27 14:24:17.903022+08		0	0	0	0	0	2024-10-31 10:52:01.101542+08	2024-10-31 10:52:01.112956+08	2024-10-31 10:52:01.119971+08	1970-01-01 08:00:00+08
20	2024-10-30 13:32:45.864509+08	2024-10-30 13:32:45.864511+08	175381207478377	宫奕辰	李敏	https://avatars.githubusercontent.com/u/12945711	minim non laborum amet Ut	tempor exercitation Lorem ullamco	鳄鱼贡献者	enim commodo nostrud Excepteur	hm1p8q.ggn39@foxmail.com	2024-10-30 13:32:45.87325+08	2024-10-30 13:32:45.87325+08		0	0	0	0	0	2024-10-31 10:52:01.101542+08	2024-10-31 10:52:01.112956+08	2024-10-31 10:52:01.119971+08	1970-01-01 08:00:00+08
23	2024-10-30 13:33:38.631482+08	2024-10-30 13:34:45.873688+08	-3155039655783567	韦敬阳	邶依诺	https://avatars.githubusercontent.com/u/29889666	Excepteur occaecat amet	ut sed	熟人倡导者，模特	ipsum dolor eiusmod sunt reprehenderit	gs2m56_k4g63@126.com	2024-10-30 13:33:38.631749+08	2024-10-30 13:33:38.631749+08		0	0	0	0	0	2024-10-31 10:52:01.101542+08	2024-10-31 10:52:01.112956+08	2024-10-31 10:52:01.119971+08	1970-01-01 08:00:00+08
26	2024-10-30 13:54:11.255281+08	2024-10-30 13:54:11.255284+08	-5977226435828183	宣浩辰	锺丽芬	https://avatars.githubusercontent.com/u/84022104	in consectetur pariatur labore	officia	类比倡导者	cupidatat aliqua aliquip	tdcsbt29@sina.com	2024-10-30 13:54:11.259969+08	2024-10-30 13:54:11.259969+08		0	0	0	0	0	2024-10-31 10:52:01.101542+08	2024-10-31 10:52:01.112956+08	2024-10-31 10:52:01.119971+08	1970-01-01 08:00:00+08
57	2024-11-02 00:21:27.482863+08	2024-11-02 00:21:27.482866+08	48657826	Kawe Mazidjatari	Mauler125	https://avatars.githubusercontent.com/u/48657826?v=4		Netherlands	Engine Tooling & Reverse Engineering.			2024-11-02 00:21:28.920423+08	2024-11-02 00:21:28.920423+08		12	7	119	0	0	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43
58	2024-11-02 00:22:28.386876+08	2024-11-02 00:22:28.38688+08	962416	Kong	Kong	https://avatars.githubusercontent.com/u/962416?v=4		San Francisco	The Cloud Connectivity Company. Community Driven & Enterprise Adopted.	https://konghq.com		2024-11-02 00:22:29.689378+08	2024-11-02 00:22:29.689378+08		484	0	1228	0	0	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43
59	2024-11-02 00:22:48.445841+08	2024-11-02 00:22:48.445844+08	1035487	Perker	perklet	https://avatars.githubusercontent.com/u/1035487?v=4			Backend developer\r\n			2024-11-02 00:22:49.924233+08	2024-11-02 00:22:49.924233+08		41	25	448	2	6	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43
60	2024-11-02 00:26:33.712026+08	2024-11-02 00:26:33.712029+08	5857640		feihua	https://avatars.githubusercontent.com/u/5857640?v=4					1002219331@qq.com	2024-11-02 00:26:34.789086+08	2024-11-02 00:26:34.789086+08		47	3	67	0	257	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43
61	2024-11-02 00:27:05.819689+08	2024-11-02 00:27:05.819691+08	11828206	Numberwolf-Yanlong	numberwolf	https://avatars.githubusercontent.com/u/11828206?v=4	@Bytedance EX: @Baidu @Tencent	HangZhou, ZheJiang	Multimedia R&D: ChangYanlong(常炎隆)\r\nQQ Group:925466059; Discord:numberwolf#8694; 	开源技术支持QQ群:925466059	porschegt23@foxmail.com	2024-11-02 00:27:07.012592+08	2024-11-02 00:27:07.012592+08		44	14	275	0	108	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43
49	2024-11-01 23:47:44.043566+08	2024-11-01 23:47:44.043569+08	38996248	ShellWen | 颉文	ShellWen	https://avatars.githubusercontent.com/u/38996248?v=4			Another Furry/🌈/Coder/Student	https://shellwen.com	me@shellwen.com	2024-11-01 23:47:45.761583+08	2024-11-01 23:47:45.761583+08	realShellWen	85	321	267	6	1342	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43
50	2024-11-01 23:48:54.044843+08	2024-11-01 23:48:54.044935+08	105362324	Chengkun Chen	seri037	https://avatars.githubusercontent.com/u/105362324?v=4	JLU	China	JLU undergraduate, majoring in Software Engineering.			2024-11-01 23:48:55.514191+08	2024-11-01 23:48:55.514191+08	serix2004	1	34	13	0	149	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43
51	2024-11-01 23:52:47.718951+08	2024-11-01 23:53:05.071867+08	38599937	hanbings	hanbings	https://avatars.githubusercontent.com/u/38599937?v=4		Your Heart	🍀 Nice to meet you!	blog.hanbings.io	hanbings@hanbings.io	2024-11-01 23:52:55.329171+08	2024-11-01 23:52:55.329171+08	IceCatHanbings	17	106	232	0	292	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43
52	2024-11-01 23:53:40.296782+08	2024-11-01 23:53:40.296784+08	74225106	go-zero team	zeromicro	https://avatars.githubusercontent.com/u/74225106?v=4			Make development easy!	https://go-zero.dev		2024-11-01 23:53:41.795173+08	2024-11-01 23:53:41.795173+08	kevwanzero	40	0	664	0	0	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43
53	2024-11-01 23:54:33.85838+08	2024-11-01 23:54:33.858383+08	1991296	Georgi Gerganov	ggerganov	https://avatars.githubusercontent.com/u/1991296?v=4	@ggml-org 	Sofia, Bulgaria	I like big .vimrc and I cannot lie	https://ggerganov.com	ggerganov@gmail.com	2024-11-01 23:54:42.503957+08	2024-11-01 23:54:42.503957+08	ggerganov	71	13	15250	10	337	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43
54	2024-11-01 23:56:52.627687+08	2024-11-01 23:56:52.627689+08	26871028	Sparrow He	sparrowhe	https://avatars.githubusercontent.com/u/26871028?v=4	Jining No.1 High School, @clipteam	Jining, Shandong, P.R.C	Fur / FullStack / React / Golang / JavaScript		sparrowhe@gmail.com	2024-11-01 23:56:53.662053+08	2024-11-01 23:56:53.662053+08		114	62	83	2	252	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43
55	2024-11-01 23:57:51.103802+08	2024-11-01 23:57:51.103804+08	11919660	Jeff Mony	JeffMony	https://avatars.githubusercontent.com/u/11919660?v=4	Happy	Shanghai	Coding after thinking\r\nE-mail: jeffmony@163.com\r\nWeChat: LOVE_BigLi	公众号：音视频平凡之路	2339762046@qq.com	2024-11-01 23:57:52.126058+08	2024-11-01 23:57:52.126058+08		101	37	244	0	565	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43
56	2024-11-02 00:21:03.537728+08	2024-11-02 00:21:03.537731+08	11687509	misumi	mismith0227	https://avatars.githubusercontent.com/u/11687509?v=4		japan		https://mismith.me/		2024-11-02 00:21:04.938335+08	2024-11-02 00:21:04.938335+08		20	34	42	0	297	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43
62	2024-11-02 00:30:08.417753+08	2024-11-02 00:30:08.417757+08	16970045	Hiroki Chen	hiroki-chen	https://avatars.githubusercontent.com/u/16970045?v=4	Unemployed	Saratoga, CA	CS Ph.D. @ IUB\r\n\r\n"Algorithms are the computational content of proofs." (Robert Harper)	hiroki-chen.github.io	haobchen@iu.edu	2024-11-02 00:30:09.588196+08	2024-11-02 00:30:09.588196+08		101	183	83	0	654	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43	0001-01-01 08:05:43+08:05:43
\.


--
-- Data for Name: create_repo; Type: TABLE DATA; Schema: relation; Owner: general_user
--

COPY relation.create_repo (data_id, developer_id, repo_id) FROM stdin;
3	0	6107760326870916
2	0	7512849632332436
4	0	7531205886403258
5	0	929876781030314
7	4404762683078221	1289668908638565
9	-7283717019284255	-4760036343357399
\.


--
-- Data for Name: follow; Type: TABLE DATA; Schema: relation; Owner: general_user
--

COPY relation.follow (data_id, follower_id, following_id) FROM stdin;
1	8379010844589220	5918517660987592
2	2585625582319098	6436475565794394
3	1949004959769462	6790165847142646
4	0	6790165847142646
5	0	7881747681974942
6	1683779426959164	0
7	1593948799959836	0
8	4189376139550958	0
9	0	5328269762254538
12	6747354333297525	2457899291576957
14	-4266425914545183	8722768799896505
15	2579195537594089	3008251821688085
\.


--
-- Data for Name: fork; Type: TABLE DATA; Schema: relation; Owner: general_user
--

COPY relation.fork (data_id, original_repo_id, fork_repo_id) FROM stdin;
1	1793833785616072	421649995112468
2	8406190102424874	1933983963942788
3	1409793094274684	5959779089443492
4	8091811042474432	3573706620660924
5	1593785440865188	4598063677940488
6	7386012362969422	6269857114891454
9	0	8801593539434776
10	0	882255865807296
11	0	8553635381376412
12	-3603338754682847	8651176517523881
13	-1281283971074775	-6600081079653567
\.


--
-- Data for Name: star; Type: TABLE DATA; Schema: relation; Owner: general_user
--

COPY relation.star (data_id, developer_id, repo_id) FROM stdin;
1	6873778198853754	3571579110738132
2	8274824970900256	7249993278181624
3	4236904963941184	3670204281797044
4	6774760386864752	8606063719622078
5	2846680938215578	3150195317120386
6	2846680938215578	0
7	3353565227513694	0
8	3113390314831850	0
9	0	2449194866770226
10	0	1010631992094434
11	0	2649646080465742
\.


--
-- Data for Name: repo; Type: TABLE DATA; Schema: repo; Owner: general_user
--

COPY repo.repo (data_id, data_created_at, data_updated_at, id, name, star_count, fork_count, issue_count, commit_count, pr_count, language, description, last_fetch_fork_at, last_fetch_contribution_at) FROM stdin;
5	2024-10-27 15:37:02.212148+08	2024-10-27 15:37:02.21215+08	3523632283650764	纵国琴	-1289752196028791	-1210269439784199	-8598683384517927	-2293580815825747	1289819019356761	{}	人广都精没。满新反报带理。感信用表非所反感路。员间意民好。经效他还华还水带长。强安前。	1970-01-01 08:00:00+08	1970-01-01 08:00:00+08
7	2024-10-27 15:37:12.567671+08	2024-10-27 15:37:12.567673+08	3067809241010708	钦诚	1821866848122965	-6071622107913391	-2014352923821635	164733007140365	5819365439908217	{}	需她格新立少。再何支如建话么再王。准离组证示资始片。计生料证北。特形口。文率亲南将。	1970-01-01 08:00:00+08	1970-01-01 08:00:00+08
9	2024-10-27 15:38:38.809677+08	2024-10-27 15:38:59.39244+08	6249417936809650	湛磊	-3250057335001851	6393638176137649	-4375831238569531	-7513812163244219	-6137185102164223	{}	去数红又装商劳数组者。部万进做学称。住接大看济马理除。	1970-01-01 08:00:00+08	1970-01-01 08:00:00+08
4	2024-10-27 12:58:06.857896+08	2024-10-28 22:01:17.674521+08	123132	腾万佳	-8530808456505775	1331975846439665	7816573501139929	-740602614358191	-3478036959738907	{}	性应看号候行表发几没。型上却价后华。电农百据六。引什非。们外如角认据见外手。切外以酸解说位次听。却这以过。切情该些了万学用道最。	1970-01-01 08:00:00+08	1970-01-01 08:00:00+08
8	2024-10-27 15:37:16.269004+08	2024-10-30 16:25:05.59473+08	8545238099456102	铎文韬	2586743306437733	5523854433059365	5926289815197205	-2076770298075035	2125876007232625	{"java": 123}	动还专越精观如位消。包业者克除八不但。被强低需提照连。自影听。何民达造特求反于斗。完分行。非面非约新效明上向先。	1970-01-01 08:00:00+08	1970-01-01 08:00:00+08
\.


--
-- Name: analysis_data_id_seq; Type: SEQUENCE SET; Schema: analysis; Owner: general_user
--

SELECT pg_catalog.setval('analysis.analysis_data_id_seq', 4, true);


--
-- Name: contribution_data_id_seq; Type: SEQUENCE SET; Schema: contribution; Owner: general_user
--

SELECT pg_catalog.setval('contribution.contribution_data_id_seq', 12, true);


--
-- Name: developer_data_id_seq; Type: SEQUENCE SET; Schema: developer; Owner: general_user
--

SELECT pg_catalog.setval('developer.developer_data_id_seq', 62, true);


--
-- Name: create_repo_data_id_seq; Type: SEQUENCE SET; Schema: relation; Owner: general_user
--

SELECT pg_catalog.setval('relation.create_repo_data_id_seq', 10, true);


--
-- Name: follow_data_id_seq; Type: SEQUENCE SET; Schema: relation; Owner: general_user
--

SELECT pg_catalog.setval('relation.follow_data_id_seq', 15, true);


--
-- Name: fork_data_id_seq; Type: SEQUENCE SET; Schema: relation; Owner: general_user
--

SELECT pg_catalog.setval('relation.fork_data_id_seq', 13, true);


--
-- Name: star_data_id_seq; Type: SEQUENCE SET; Schema: relation; Owner: general_user
--

SELECT pg_catalog.setval('relation.star_data_id_seq', 12, true);


--
-- Name: repo_data_id_seq; Type: SEQUENCE SET; Schema: repo; Owner: general_user
--

SELECT pg_catalog.setval('repo.repo_data_id_seq', 10, true);


--
-- Name: analysis analysis_pk; Type: CONSTRAINT; Schema: analysis; Owner: general_user
--

ALTER TABLE ONLY analysis.analysis
    ADD CONSTRAINT analysis_pk PRIMARY KEY (data_id);


--
-- Name: analysis analysis_pk_2; Type: CONSTRAINT; Schema: analysis; Owner: general_user
--

ALTER TABLE ONLY analysis.analysis
    ADD CONSTRAINT analysis_pk_2 UNIQUE (developer_id);


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
-- Name: analysis update_time; Type: TRIGGER; Schema: analysis; Owner: general_user
--

CREATE TRIGGER update_time BEFORE UPDATE ON analysis.analysis FOR EACH ROW EXECUTE FUNCTION analysis.update_time();


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

