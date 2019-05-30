// *********************************************************
// RTView - Simple Example to Define Caches

// Configure the application to make use of the rtview-utils package
// and reference it using the variable name 'rtview'
   
var rtview = require('rtview-utils')();

// Configure package to send data to an RTView DataServer running on 'localhost:3270'
// and exposing an http servlet named 'rtvpost' (the default out-of-the-box dataserver)

//var url = 'http://localhost:3275';
var url = 'http://localhost:3270/rtvpost';
rtview.set_targeturl(url);

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



