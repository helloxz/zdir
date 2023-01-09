-- 创建options表
CREATE TABLE z_options (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"key" TEXT(512) NOT NULL,
	value TEXT NOT NULL,
	note TEXT, created_at int(13) NOT NULL, updated_at int(13),
	CONSTRAINT z_options_UN UNIQUE ("key")
);

-- 导入默认数据
INSERT INTO z_options
("key", value, note, created_at, updated_at)
VALUES('site_data', '{"custom_header":"","description":"Zdir3是一款轻量级目录列表程序，使用Golang + Vue3开发。","footer":"","keywords":"zdir,目录列表","logo":"","title":"Zdir"}', '', 1672293223, 1672583054);

-- 创建z_db_logs表，用于记录SQL更新日志
CREATE TABLE z_db_logs (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	sql_name TEXT(32) NOT NULL,
	created_at INTEGER NOT NULL,
	updated_at INTEGER,
	state INTEGER DEFAULT 1 NOT NULL,
	note TEXT(512),
	CONSTRAINT z_db_logs_UN UNIQUE (sql_name)
);

-- 创建z_login_logs，用来记录用户登录日志
CREATE TABLE z_login_logs (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	cid TEXT(6) NOT NULL,
	token TEXT(32) NOT NULL,
	behavior TEXT(16) DEFAULT login NOT NULL,
	created_at INTEGER NOT NULL,
	updated_at INTEGER,
	expired_at INTEGER NOT NULL,
	ip TEXT(50) NOT NULL,
	ua TEXT(1024) NOT NULL,
	state INTEGER DEFAULT 1 NOT NULL,
	note TEXT(1024)
);
