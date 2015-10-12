'use strict';


app.controller('VideoManagerCtrl', function ($scope, i18nService, $modal, $log,cfpLoadingBar,$http) {
    //初始化查询参数
    //以下处理是为了防止树内容一次加载导致的性能问题,改为了点击加载 , 使用了非deepcopy
    $scope.clients = [];
    $http.get("/video/clients").then(function(ret){
        $scope.clients = ret.data;
        console.log(ret.data);
    });

});
