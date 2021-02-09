CREATE TABLE usercameras(
    id INTEGER PRIMARY KEY,
    ip VARCHAR(120),
    password TEXT NOTNULL,
    username TEXT NOTNULL,
    last_time_login INTEGER NOTNULL
);
CREATE TABLE userclients(
    cookie TEXT
);

