-- v4 (compatible with v2+): Clear invalid disappearing timers in portal table
UPDATE portal
SET disappear_type=NULL, disappear_timer=NULL
WHERE disappear_timer=0 OR disappear_timer IS NULL OR disappear_type IS NULL;
