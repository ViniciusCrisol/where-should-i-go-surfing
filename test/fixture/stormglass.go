package fixture

const (
	StormglassTwoValidPointsJSONResponse = `
		{
		    "hours": [
		        {
		            "time": "2020-04-26T00:00:00+00:00",
		            "swellDirection": {"noaa": 64.26}, "swellHeight": {"noaa": 0.15}, "swellPeriod": {"noaa": 3.89},
		            "waveDirection": {"noaa": 23.38}, "waveHeight": {"noaa": 0.47},
		            "windDirection": {"noaa": 19.45}, "windSpeed": {"noaa": 10.1}
		        },
		        {
		            "time": "2020-04-26T01:00:00+00:00",
		            "swellDirection": {"noaa": 12.4}, "swellHeight": {"noaa": 0.21}, "swellPeriod": {"noaa": 3.67},
		            "waveDirection": {"noaa": 23.1}, "waveHeight": {"noaa": 0.46}, 
		            "windDirection": {"noaa": 131.4}, "windSpeed": {"noaa": 10.1}
		        }
		    ],
		    "meta": {
		        "lat": 100.1,
		        "lng": 100.1,
		        "start": "2020-04-26 00:00",
		        "end": "2020-04-27 00:00",
		        "source": ["noaa"],
		        "params": [
		            "swellDirection", "swellHeight", "swellPeriod",
		            "waveDirection", "waveHeight",
		            "windDirection", "windSpeed"
		        ],
		        "cost": 1,
		        "dailyQuota": 25,
		        "requestCount": 10
		    }
		}
	`

	StormglassOneValidPointJSONResponse = `
		{
		    "hours": [
		        {
		            "time": "2020-04-26T00:00:00+00:00",
		            "swellDirection": {"noaa": 64.26}, "swellHeight": {"noaa": 0.15}, "swellPeriod": {"noaa": 3.89},
		            "waveDirection": {"noaa": 23.38}, "waveHeight": {"noaa": 0.47},
		            "windDirection": {"noaa": 19.45}, "windSpeed": {"noaa": 10.1}
		        },
		        {
		            "time": "2020-04-26T01:00:00+00:00",
		            "swellDirection": {"nooa": 12.4}, "swellHeight": {"nooa": 0.21}, "swellPeriod": {"nooa": 3.67},
		            "waveDirection": {"nooa": 23.1}, "waveHeight": {"nooa": 0.46}, 
		            "windDirection": {"nooa": 31.4}, "windSpeed": {"nooa": 10.1}
		        }
		    ],
		    "meta": {
		        "lat": 100.1,
		        "lng": 100.1,
		        "start": "2020-04-26 00:00",
		        "end": "2020-04-27 00:00",
		        "source": ["noaa"],
		        "params": [
		            "swellDirection", "swellHeight", "swellPeriod",
		            "waveDirection", "waveHeight",
		            "windDirection", "windSpeed"
		        ],
		        "cost": 1,
		        "dailyQuota": 25,
		        "requestCount": 10
		    }
		}
	`

	StormglassRateLimitReachedJSONResponse = `
		{
		    "errors": ["rate limit reached"]
		}
	`
)
