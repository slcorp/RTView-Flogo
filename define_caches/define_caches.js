// *********************************************************
// RTView - Simple Example

// Configure the application to make use of the rtview-utils package
// and reference it using the variable name 'rtview'
   
var rtview = require('rtview-utils')();

// Configure package to send data to an RTView DataServer running on 'localhost:3270'
// and exposing an http servlet named 'rtvpost' (the default out-of-the-box dataserver)

//var url = 'http://localhost:3275';
var url = 'http://localhost:3270/rtvpost';
rtview.set_targeturl(url);

// Send data when number of buffered rows reaches 10

var size = 10;            // default is 50
rtview.set_batchsize(size);

// Send data at an interval of five seconds

var interval = 5000;       // default is 2000 ms
rtview.set_interval(interval);

// Create a cache named SensorData with the specified properties and structure

var sensorCacheName  = "SensorData";
var sensorProperties = {
    "indexColumnNames" : "ID",
    "historyColumnNames" : "temperature;humidity",
    "condenseRowsGroupBy": "temperature:average;humidity:average"
};
var sensorMetadata = [ 
    { "ID" : "string" },
    { "temperature" : "double" },
    { "humidity" : "double" },
    { "temp_unit" : "string"}    // a column that is neither index nor history
];
rtview.create_datacache(sensorCacheName, sensorProperties, sensorMetadata);

var numberOfSensors = 5;
var temperatureSeed = 71;
var humiditySeed = 53;

processData();

async function processData() {
    
	while (true) {

		for (var i = 0; i < numberOfSensors; i++) {
			
			var data = {};
			data.ID = 'Sensor-' + (i+1).toString();
			data.temperature = parseFloat(temperatureSeed + (Math.random() * 0.3)).toFixed(2);
			data.humidity = parseFloat(humiditySeed + (Math.random() * 0.4)).toFixed(2);
			data.temp_unit = '\xB0C';    // '°C'
			
			console.log('\n... sending data: ' + JSON.stringify(data));
			rtview.send_datatable(sensorCacheName, data);
		}
		console.log();             // Separate the output per interval
			
		await sleep(10000);        // create data points every 10 seconds
    }
}

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

