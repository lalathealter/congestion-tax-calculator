# API README

## congestion-calculator endpoint

The Congestion Tax Calculator's main endpoint is at "/congestion-calculator/:location", where ":location" must be a name of the rule set we are calculating by. 

The only rule set developed yet is of Gothenburg; To access it, simply address "/congestion-calculator/gothenburg"

To use the Calculator, send an HTTP POST request, the body of which should comply to the following schema (in JSON):
```JSON
{
		Intervals: [
			"2013-02-08 15:29:00",
			"2013-02-08 15:47:00",
			"2013-02-08 16:01:00",
	  ], 
    Vehicle: { Type: "Diplomat", Data: {} }
}
```

- ```Intervals``` is an array of time strings representing the exact moment a vehicle passed a checker; All should be in a format of "YYYY-MM-DD HH:MM:SS"
- ```Vehicle``` must be an object with the two fields of ```Type``` and ```Data```
- - ```Type``` is a string representing a type of the vehicle passing; An empty string would be interpreted as a basic vehicle type (that is, a car). The only other options allowed are: 
- - - "Diplomat", "Tractor", "Military", "Foreign", "Motorcycle", "Emergency", "Bus"
- - ```Data``` is an object that is dedicated to any supplementary data about a vehicle (it may be used whenever a certain rule set has some additional calculations based on the characteristic of a vehicle). For the time being none of the types use this field, so it is recommended to leave it an empty object
