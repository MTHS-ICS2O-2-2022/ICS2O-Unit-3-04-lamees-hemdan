// Copyright (c) 2020 Mr. Coxall All rights reserved
//
// Created by: Lamees Hemdan
// Created on: April 2023
// This file contains the JS functions for index.html

// This function converts from Fahrenheit to Celsius

  function fahrenheitToCelsius () {
// input

  const fahrenheit = parseFloat(document.getElementById('fahrenheit').value);
// process

  const celsius = (fahrenheit - 32) * 5 / 9;

// output

  document.getElementById('celsius').innerHTML = 'Temperature in Celsius: ' + celsius;
}
