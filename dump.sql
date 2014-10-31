PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE forums(id INTEGER PRIMARY KEY, title varchar(255), description varchar(255));
INSERT INTO "forums" VALUES(1,'test','tester forum');
INSERT INTO "forums" VALUES(2,'forum zwei','eine Prüfung');
CREATE TABLE topics(id INTEGER PRIMARY KEY, title varchar(255), description varchar(255), forum_id integer, FOREIGN KEY(forum_id) REFERENCES forum(id));
INSERT INTO "topics" VALUES(1,'test topic','asdf asdf asdf',1);
INSERT INTO "topics" VALUES(2,'test topic','for forum 2: asdf asdf asdf',2);
CREATE TABLE posts(id INTEGER PRIMARY KEY, text TEXT, published TIMESTAMP, topic_id INTEGER, FOREIGN KEY(topic_id) REFERENCES topic(id));
INSERT INTO "posts" VALUES(1,'test','2014-10-31 07:50:55.810912273',1);
INSERT INTO "posts" VALUES(2,'test2','2014-10-31 07:52:32.129118657',1);
INSERT INTO "posts" VALUES(3,'test3','2014-10-31 07:52:51.073031409',1);
INSERT INTO "posts" VALUES(4,'test4','2014-10-31 07:52:55.094815942',1);
INSERT INTO "posts" VALUES(5,'Hello, this is a simple test!','2014-10-31 07:54:45.416706454',2);
COMMIT;
