'use strict';

var Fabric_Client = require('fabric-client');
var path = require('path');
var util = require('util');
var os = require('os');
var jsonConverter = require('json-style-converter/es5')
//
var fabric_client = new Fabric_Client();

// setup the fabric network
var channel = fabric_client.newChannel('nid-channel');
var peer = fabric_client.newPeer('grpc://localhost:7051');
channel.addPeer(peer);

//
// var member_user = null;
var store_path = path.join(__dirname, 'hfc-key-store');
console.log('Store path:'+store_path);
var tx_id = null;

function marshalArgs(args) {
	if (!args) {
	  return args;
	}
  
	if (typeof args === 'string') {
	  return [args];
	}
  
	var snakeArgs = jsonConverter.camelToSnakeCase(args);
  
	if (Array.isArray(args)) {
	  return snakeArgs.map(
		arg => typeof arg === 'object' ? JSON.stringify(arg) : arg.toString());
	}
  
	if (typeof args === 'object') {
	  return [JSON.stringify(snakeArgs)];
	}
}
// create the key value store as defined in the fabric-client/config/default.json 'key-value-store' setting
function queryChaincode(funcName, args) {
	return Fabric_Client.newDefaultKeyValueStore({ path: store_path
	}).then((state_store) => {
		// assign the store to the fabric client
		fabric_client.setStateStore(state_store);
		var crypto_suite = Fabric_Client.newCryptoSuite();
		// use the same location for the state store (where the users' certificate are kept)
		// and the crypto store (where the users' keys are kept)
		var crypto_store = Fabric_Client.newCryptoKeyStore({path: store_path});
		crypto_suite.setCryptoKeyStore(crypto_store);
		fabric_client.setCryptoSuite(crypto_suite);

		// get the enrolled user from persistence, this user will sign all requests
		return fabric_client.getUserContext('user2', true);
	}).then((user_from_store) => {
		if (user_from_store && user_from_store.isEnrolled()) {
			console.log('Successfully loaded user2 from persistence');
			var member_user = user_from_store;
		} else {
			throw new Error('Failed to get user2.... run registerUser.js');
		}
		// queryCar chaincode function - requires 1 argument, ex: args: ['CAR4'],
		// queryAllCars chaincode function - requires no arguments , ex: args: [''],
		const request = {
			//targets : --- letting this default to the peers assigned to the channel
			chaincodeId: 'nidchain',
			fcn: funcName,
			args: marshalArgs(args)
		};

		// send the query proposal to the peer
		return channel.queryByChaincode(request);
	}).then((response_payloads) => {
		if (response_payloads) {
			var value = '';
			for(let i = 0; i < response_payloads.length; i++) {
				if(value === '') {
					value = response_payloads[i].toString('utf8');
				} else if(value !== response_payloads[i].toString('utf8')) {
					throw new Error('Responses from peers don\'t match');
				}
			}
			return JSON.parse(value);
		} else {
			console.log('response_payloads is null');
			throw new Error('Failed to get response on query');
		}
	},	(err) => {
		console.log('Failed to send query due to error: ' + err.stack ? err.stack : err);
		throw new Error('Failed, got error on query');
	});

};

module.exports.queryChaincode = queryChaincode
