
/* TODO: Perhaps we need to get more relevant details about radio stations
* like location, type (community, commercial etc)

*/
CREATE TABLE radio_stations (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(150) NOT NULL COMMENT "The name of the radio station like 'Peace FM', 'Joy FM' etc...",
    frequency VARCHAR(8) NOT NULL COMMENT "The radio station's frequency like '104.3', '99.7' etc...",
    stream_url VARCHAR(255) NOT NULL COMMENT "The url the radio station is streaming on"
) Engine=InnoDB Charset=UTF8mb4 Comment="This table holds the details of radio stations we are monitoring";

CREATE TABLE recordings (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    date_recorded DATETIME NOT NULL,
    date_recorded_unix VARCHAR(20) NOT NULL,
    radio_station INT UNSIGNED NOT NULL COMMENT 'Which radio station is this recording from',
    FOREIGN KEY radio_station_rec (radio_station) REFERENCES radio_stations(id)
)Engine=InnoDB Charset=UTF8mb4 Comment="This table holds details of recordings we've made";