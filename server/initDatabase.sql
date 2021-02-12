CREATE TABLE usercameras(
    id INTEGER PRIMARY KEY,
    ip TEXT NOT NULL,
    password TEXT NOT NULL,
    token TEXT ,
    username TEXT NOT NULL,
    last_time_login INTEGER NOT NULL
);
-- if we need to add a new colum 
-- ALTER TABLE <table> ADD <column> <type>;
CREATE TABLE userclients(
    cookie TEXT
);

