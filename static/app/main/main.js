'use strict';

angular.module('videosystem.main', ['ngRoute','ui.bootstrap'])

    .config(['$routeProvider', function ($routeProvider) {
        $routeProvider.when('/main', {
            templateUrl: '/app/main/main.html',
            controller: 'MainCtrl'
        });
    }])

    .controller('MainCtrl', function ($scope, Auth, $location ,$q) {
        $scope.tabs =[];
        $scope.activetext = '';
        var init = $q.defer();
        $scope.init = function(){
            return  init.promise;
        };
        //登录状态
        //Auth.checklogin().then(function () {
            //取出登录数据


        //}).catch(function () {
        //    $location.path('/videosystem/main');
        //});
        var lastActive ;
        $scope.loadTab = function (menu) {
            console.log(menu);
            if(menu.js){
                $script(menu.js, function () {
                    var flag = false;
                    for (var i = 0; i < $scope.tabs.length; i++) {
                        if ($scope.tabs[i].text == menu.text) {
                            flag = true;

                        }else{
                            $scope.tabs[i].active = false;
                        }
                    }
                    if (!flag) {
                        menu.close=true;
                        $scope.tabs.push(menu)
                    }
                    menu.active = true;
                    $scope.$digest();
                });
            }else{
                var flag = false;
                for (var i = 0; i < $scope.tabs.length; i++) {
                    if ($scope.tabs[i].text == menu.text) {
                        flag = true;

                    }else{
                        $scope.tabs[i].active = false;
                    }
                }
                if (!flag) {
                    menu.close=true;
                    $scope.tabs.push(menu)
                }
                menu.active = true;
            }
        };
        //for(var i =0  ;i<10;i++){
        //    var dashboard = {
        //        "text": '测试'+i,
        //        'html': 'app/test/test.html',
        //        "close": false
        //    };
        //    $scope.loadTab(dashboard);
        //}

        $("#sidebar-collapse").on('click', function () {
            if (!$('#sidebar').is(':visible'))
                $("#sidebar").toggleClass("hide");
            $("#sidebar").toggleClass("menu-compact");
            $(".sidebar-collapse").toggleClass("active");
            var b = $("#sidebar").hasClass("menu-compact");

            if ($(".sidebar-menu").closest("div").hasClass("slimScrollDiv")) {
                $(".sidebar-menu").slimScroll({ destroy: true });
                $(".sidebar-menu").attr('style', '');
            }
            if (b) {
                $(".open > .submenu")
                    .removeClass("open");
            } else {
                if ($('.page-sidebar').hasClass('sidebar-fixed')) {
                    var position = (readCookie("rtl-support") || location.pathname == "/index-rtl-fa.html" || location.pathname == "/index-rtl-ar.html") ? 'right' : 'left';
                    $('.sidebar-menu').slimscroll({
                        height: 'auto',
                        position: position,
                        size: '3px',
                        color: themeprimary
                    });
                }
            }
            //Slim Scroll Handle
        });

        $scope.user = Auth.getUser();
        //取出菜单数据
        $scope.menu = [{
            "text": '视频管理',
            "js": "/app/videomanager/videomanager.js",
            'html': '/app/videomanager/videomanager.html'
        }];

        //var dashboard = {
        //    "text": '简介',
        //    "js": "/app/dashboard/dashboard.js",
        //    'html': '/app/dashboard/dashboard.html',
        //    "close": false
        //};
        //$scope.loadTab(dashboard);
        $scope.loadTab($scope.menu[0]);
    });