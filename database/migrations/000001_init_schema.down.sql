DROP TABLE IF EXISTS robot_slots;
DROP INDEX IF EXISTS robots_location_idx;
DROP TABLE IF EXISTS robots;
-- PostGIS extension shouldn't be dropped as it might be used by other databases on the same cluster, or just better left enabled.
