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

        function ping() {
            post("version", {}).then(function (res) {
                console.log(res);
            }, function (err) {
                console.error(err);
            });
        }

        function post(url, data) {
            return $http.post(host + '/' + url, data);
        }

    }

    kdb.service('KairosdbClient', ['$http', KairosdbClient]);

})();