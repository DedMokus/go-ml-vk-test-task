CREATE TABLE IF NOT EXISTS docs
(
    id SERIAL PRIMARY KEY,
    url TEXT,
    pubdate BIGINT,
    fetchtime BIGINT,
    text TEXT,
    firstfetchtime BIGINT
);
