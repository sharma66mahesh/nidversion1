'use strict';

var express = require('express');
var bodyParser = require('body-parser');
var http = require('http');
var app = express();
var cors = require('cors');
var path = require('path');

var createFormForAdmin = require('./createFormForAdmin.js');
var listOfPdma = require('./listOfPdma.js');

var port = process.env.PORT || 4000;

app.options('*', cors());
app.use(cors());

app.use(bodyParser.json());

var server = http.createServer(app).listen(port, function() {});
console.log(`Server running on port ${port}`);
server.timeout = 24000;

function getErrorMessage(field) {
	var response = {
		success: false,
		message: field + ' field is missing or Invalid in the request'
	};
	return response;
}

app.post('/api/:fcn', async function(req,res) {
    var fcn = req.params.fcn;
    var args = req.body.args;

    if (!fcn) {
        res.json(getErrorMessage('\'fcn\''));
        return;
    }
    if (!args) {
		res.json(getErrorMessage('\'args\''));
		return;
    }
    
    createFormForAdmin.invokeChaincode(fcn,args).then(() => {
        res.json({success: true, message: 'Chaincode invoked', args: args}); 
    }, (err) => {
        res.json({success: false, message: err.message});
    });
});

app.get('/api/:fcn', async function(req,res) {
    var fcn = req.params.fcn;
    var args = req.body.args;

    if (!fcn) {
        res.json(getErrorMessage('\'fcn\''));
        return;
    }
    if (!args) {
        res.json(getErrorMessage('\'args\''));
        return;
    }
    
    listOfPdma.queryChaincode(fcn,args).then((result) => {
        res.json({success: true, message: result}); 
    }, (err) => {
        res.json({success: false, message: err.message});
    });
});
