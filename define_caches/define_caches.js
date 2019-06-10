// Configure the application to make use of the rtview-utils package
// and reference it using the variable name 'rtview'

var rtview = require('rtview-utils')();

// Configure package to send data to an RTView DataServer running on 'localhost:3270'
// and exposing an http servlet named 'rtvpost' (the default out-of-the-box dataserver)

var url = 'http://localhost:3270/rtvpost';
rtview.set_targeturl(url);

// Create a cache named SensorData with the specified properties and structure

var FlowPerfCacheName  = "fgFlowPerformance";
var FlowPerfProperties = {
    "indexColumnNames" : "hostName;flowName",
    "historyColumnNames" : "duration"
};
var FlowPerfMetadata = [ 
    { "hostName" : "string" },
    { "flowName" : "string" },
    { "duration" : "double" },
];

rtview.create_datacache(FlowPerfCacheName, FlowPerfProperties, FlowPerfMetadata);




