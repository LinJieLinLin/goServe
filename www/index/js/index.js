var indexApp = "";
indexApp = angular.module("linIndex", ["ngRoute","ngResource"])
    .config(['$routeProvider', function ($routeProvider) {
        $routeProvider.
            when('/', { controller:uploadCtrl, templateUrl: 'upload/a.html' }).
            when("/upload", {controller: uploadCtrl, templateUrl: 'upload/upload.html'}).
            when("/textEdit", {controller: editCtrl, templateUrl: 'edit/edit.html'})
    }]);
function indexCtrl($scope){
    $scope.aa = "aaa";
}