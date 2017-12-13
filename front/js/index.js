const angular = require('angular')
const $ = require('jquery')
const units = require('ethereumjs-units')

const endpoint = 'api'

window.addEventListener('load', function() {
  var p = document.getElementById("content-node");
  var p_prime = p.cloneNode(true);
  document.body.appendChild(p_prime);
});

ttt = angular.module('dimpleApp', [])
ttt.controller('dimpleCtl', function DimpleCtrl($scope, $http, $timeout) {
    var line = {txid: "…", addr: "…", amount: "…"}
    $scope.transactions = [line]
    $scope.canSend = true;
    var updateList = function () {
        $http.get(endpoint + '/list').
          then(function(response) {
              $scope.status = response.status;
              $scope.transactions = response.data.map(x => { return {
                  txid: x.txid,
                  addr: x.addr,
                  amount: units.convert(x.amount, 'wei', 'eth') + ' eth'
              }});
            }, function(response) {
              console.log(response.data || 'Request failed')
              $scope.status = response.status
        });
        $timeout(updateList, 1000)
    }
    $scope.send = function () {
        var addr = $('#addr').val()
        $http.get(endpoint + '/send?to=' + addr,{
            headers: {
        'Authorization': grecaptcha.getResponse()
    }}).
          then(function(response) {
              $scope.status = response.status;
              console.log(response)
            }, function(response) {
              console.log(response.data || 'Request failed')
              $scope.status = response.status
        });
        $scope.canSend = false;
    }
    updateList()
});
