/**
 * Created by gpl on 15/10/4.
 */
(function () {
    var kdb = angular.module('kairosdb', []);

    var host = "http://localhost:9008/api/v1";

    kdb.run(function () {
        console.log('kdb client init!');
    });


    // the api wrapper
    function KairosdbClient($http) {
        // test if inject is ok
        this.ping = ping;
        this.metricNames = metricNames;

        function ping() {
            get("version").then(function (res) {
                console.log(res.data.version);
            }, function (err) {
                console.error(err);
            });
        }

        function metricNames() {
            get("metricnames").then(function (res) {
                console.log(res.data.results);
            }, function (err) {
                console.error(err);
            });
        }

        function get(url) {
            return $http.get(host + '/' + url);
        }

        function post(url, data) {
            return $http.post(host + '/' + url, data);
        }

    }

    kdb.service('KairosdbClient', ['$http', KairosdbClient]);

})();