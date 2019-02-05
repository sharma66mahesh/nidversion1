'use strict';

var log4js = require('log4js');
var logger = log4js.getLogger('TradeApp');
var express = require('express');
var bodyParser = require('body-parser');
var http = require('http');
var util = require('util');
var app = express();
var expressJWT = require('express-jwt');
var jwt = require('jsonwebtoken');
var bearerToken = require('express-bearer-token');
var cors = require('cors');
var path = require('path');

var invokeCC = require('./invoke.js');
var queryCC = require('./query.js');

var host = process.env.HOST || 'localhost' ;
var port = process.env.PORT || 4000;

app.options('*', cors());
app.use(cors());

app.use(bodyParser.json());

var server = http.createServer(app).listen(port, function() {});
console.log('****************** SERVER STARTED ************************');
console.log('***************  http://%s:%s  ******************',host,port);
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
    
    invokeCC.invokeChaincode(fcn,args).then(() => {
        res.json({success: true, message: 'Chaincode invoked'}); 
    }, (err) => {
        res.json({success: false, message: err});
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
    
    queryCC.queryChaincode(fcn,args).then((result) => {
        res.json({success: true, message: result}); 
    }, (err) => {
        res.json({success: false, message: err});
    });
});
