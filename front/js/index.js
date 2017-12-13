const angular = require('angular')
const units = require('ethereumjs-units')

const endpoint = 'api'

const sitekey = 'YOUR SITE KEY'

window.addEventListener('load', function() {
  grecaptcha.render('captcha', {
    'sitekey' : sitekey,
    'theme' : 'dark'
  });
});

ttt = angular.module('dimpleApp', [])
ttt.controller('dimpleCtl', function DimpleCtrl($scope, $http, $timeout) {
    var line = {txid: "…", addr: "…", amount: "…"}
    $scope.transactions = [line]
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
        console.log('captcha response: ' + grecaptcha.getResponse());
    }
    updateList()
});
