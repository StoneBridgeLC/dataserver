create table topic (
    id  INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    topic TEXT,
    positive INT,
    negative INT
);

create table news (
    id    INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    title   TEXT NOT NULL,
    body  TEXT NOT NULL,
    hash    VARCHAR(64) NOT NULL,
    create_time  DATETIME,
    update_time  DATETIME
);

create table news_topic (
    nid   INT NOT NULL,
    tid   INT NOT NULL,
    PRIMARY KEY(nid, tid),
    FOREIGN KEY (nid) REFERENCES news (id) ON DELETE CASCADE,
    FOREIGN KEY (tid) REFERENCES topic (id) ON DELETE CASCADE
);

create table comment (
    id  INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    nid INT NOT NULL,
    body    TEXT NOT NULL,
    pid INT,
    is_pos   TINYINT,
    create_time  DATETIME,
    update_time  DATETIME,
    FOREIGN KEY (pid) REFERENCES comment (id) ON DELETE CASCADE
);