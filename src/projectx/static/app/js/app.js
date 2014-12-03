"use strict";

angular.module('projectx', [
	'ngResource',
	'ngRoute',
])
.config(['$routeProvider', '$locationProvider', function($routeProvider, $locationProvider) {
	$routeProvider.when('/orders', { templateUrl: '/app/js/partials/orders.html', controller: 'ordersController' });
	$routeProvider.when('/order/:id', { templateUrl: '/app/js/partials/order.html', controller: 'orderController' });
	$routeProvider.otherwise({ redirectTo: '/orders' });
	
	$locationProvider.html5Mode({
		enabled: true
	});
}])
.factory('Order', ['$resource', function($resource) {
	return $resource('/rest/orders');
}])
.controller('ordersController', ['$scope', 'Order', function($scope, Order) {
	$scope.orders = Order.query();
}])
.controller('orderController', ['$scope', 'Order', '$routeParams', function($scope, Order, routeParams) {
	$scope.order = Order.get({ id: routeParams.id });
}])
;