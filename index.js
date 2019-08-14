'use strict';
 
const XMLHttpRequest = require("xmlhttprequest").XMLHttpRequest;
const functions = require('firebase-functions');
const request = require('request');
const {dialogflow} = require ('actions-on-google');
const {WebhookClient} = require('dialogflow-fulfillment');
const {Card, Suggestion} = require('dialogflow-fulfillment');
const WELCOME_INTENT = 'Default Welcome Intent';
const FALLBACK_INTENT = 'Default Fallback Intent';
const WEATHER_INTENT = 'weather';
const WEATHER_ENTITY = 'weather';

function getData(url, callback) {
  return new Promise((resolve, reject) => {
    request.get(url, (error, response, body) => {
      callback(body);
      resolve(body);
    });
});
}
 
const app = dialogflow();
app.intent(WELCOME_INTENT, (conv) => {
    conv.add("Welcome to my Sample Agent!");
});
app.intent(WEATHER_INTENT, (conv, WEATHER_ENTITY) => {
     const queryUrl = "https://us-central1-swift-setup-249716.cloudfunctions.net/getWeather";
     return getData(queryUrl,(data) => {
            conv.add(data);
        });
});

exports.dialogflowFirebaseFulfillment = functions.https.onRequest(app);
