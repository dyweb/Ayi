/**
 * Created by gpl on 15/10/4.
 */
(function () {
    'use strict';
    var app = angular.module('ayi', ['kairosdb']);
    app.run(['KairosdbClient',function(KairosdbClient){
        KairosdbClient.ping();
    }]);
})();