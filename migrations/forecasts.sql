CREATE TABLE IF NOT EXISTS forecasts ( 
	day int(10) unsigned NOT NULL, 
	drought int(10) unsigned NOT NULL, 
	optimal_weather int(10) unsigned NOT NULL, 
	rainy int(10) unsigned NOT NULL, 
	UNIQUE KEY day_unique (day) 
) ENGINE=InnoDB