CREATE EXTENSION IF NOT EXISTS postgis;

CREATE TABLE robots (
    id UUID PRIMARY KEY,
    status VARCHAR(50) NOT NULL, -- IDLE, EN_ROUTE, MAINTENANCE
    location GEOGRAPHY(Point, 4326) NOT NULL, -- Tọa độ PostGIS
    current_battery DECIMAL(5,2) NOT NULL
);
CREATE INDEX idx_robots_location ON robots USING GIST (location);

CREATE TABLE robot_slots (
    robot_id UUID REFERENCES robots(id),
    available_slots INT NOT NULL,
    PRIMARY KEY (robot_id)
);

CREATE TABLE outbox_events (
    id UUID PRIMARY KEY,
    aggregate_type VARCHAR(255) NOT NULL,
    aggregate_id VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    payload JSONB NOT NULL,
    status VARCHAR(50) DEFAULT 'PENDING' -- PENDING, PUBLISHED
);
