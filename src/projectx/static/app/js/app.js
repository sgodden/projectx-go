"use strict";

angular.module('projectx', [
		'ngResource',
		'ngRoute'
	])

	.config(['$routeProvider', '$locationProvider', function ($routeProvider, $locationProvider) {
		$routeProvider.when('/orders', {templateUrl: '/app/js/partials/orders.html', controller: 'ordersController'});
		$routeProvider.when('/order/:id', {templateUrl: '/app/js/partials/order.html', controller: 'orderController'});
		$routeProvider.when('/neworder', {templateUrl: '/app/js/partials/order.html', controller: 'orderController'});
		$routeProvider.otherwise({redirectTo: '/orders'});

		$locationProvider.html5Mode({
			enabled: true
		});
	}])

	.factory('Order', ['$resource', function ($resource) {
		return $resource('/rest/order/:id');
	}])

	.controller('ordersController', ['$scope', 'Order', '$location', function ($scope, Order, $location) {
		$scope.orders = Order.query();
		$scope.new = function () {
			$location.path('/neworder');
		};
	}])

	.controller('orderController', ['$scope', 'Order', '$routeParams', function ($scope, Order, routeParams) {
		if (routeParams.id) {
			$scope.order = Order.get({id: routeParams.id});
		} else {
			$scope.order = { OrderNumber: undefined};
		}

		$scope.submit = function () {
			if ($scope.form.$invalid) {
				console.debug("The form is invalid");
				return;
			} else {
				console.debug("Saving the order");
				Order.save($scope.order);
			}
		}
	}]);