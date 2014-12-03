"use strict";

angular.module('projectx', [
	'ngResource'
])
.factory('Order', ['$resource', function($resource) {
	return $resource('/rest/orders');
}])
.controller('ordersController', ['$scope', 'Order', function($scope, Order) {
	$scope.orders = Order.query();
}]);