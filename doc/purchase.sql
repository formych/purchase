CREATE TABLE `purchase_info` (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	user VARCHAR(20) NOT NULL DEFAULT '',
	company VARCHAR(30) NOT NULL DEFAULT '',
	tel VARCHAR(20) NOT NULL DEFAULT '',
	purchase_num BIGINT(20) NOT NULL DEFAULT '',
	purchase_time timstamp NOT NULL,
	created_time tiemstamp NOT NULL DEFAULT CUREENT_TIMESTAMP,
	updated_time tiemstamp NOT NULL DEFAULT CUREENT_TIMESTAMP
)