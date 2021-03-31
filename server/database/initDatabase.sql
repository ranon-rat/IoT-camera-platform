-- if we need to add a new colum 
-- ALTER TABLE <table> ADD <column> <type>;
CREATE TABLE usercameras(
    id          INTEGER,
    ip          VARCHAR(64) PRIMARY KEY NOT NULL UNIQUE,
    password    VARCHAR(64) NOT NULL UNIQUE,
    token       VARCHAR(64) UNIQUE,
    username    TEXT NOT NULL UNIQUE,

    last_time_login INTEGER NOT NULL   
    
    
);
CREATE TABLE userclients(
    id_camera_client TEXT NOT  NULL,
     id INTEGER PRIMARY KEY,
    -- this is for be related with the table usercameras
    cookie TEXT 
    
);

